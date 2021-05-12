package stop

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	clientset "github.com/tilt-dev/tilt-ci-status/pkg/clientset/versioned"
	"github.com/tilt-dev/tilt-ci-status/pkg/config"
	"github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Look I don't even care about naming anymore okay?!
type Stopper struct {
	Cli *clientset.Clientset
}

func NewStopper() (Stopper, error) {
	// TODO - how do we handle multiple tilt instances?
	cfg, err := config.NewConfig()
	if err != nil {
		return Stopper{}, errors.Wrap(err, "getting tilt api config")
	}

	return Stopper{
		Cli: clientset.NewForConfigOrDie(cfg),
	}, nil
}

func (s *Stopper) StopResource(ctx context.Context, resource string) error {
	cmd, err := s.CmdForResource(ctx, resource)
	if err != nil {
		return errors.Wrap(err, "getting cmd for resource")
	}

	return s.StopCmd(ctx, cmd)

}
func (s *Stopper) CmdForResource(ctx context.Context, resource string) (*v1alpha1.Cmd, error) {
	cmds, err := s.AllCmdsByResource(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting all Cmds")
	}
	cmd, ok := cmds[resource]
	if !ok {
		var allResources []string
		for resource := range cmds {
			allResources = append(allResources, resource)
		}
		return nil, fmt.Errorf("no Cmd found for resource %q (found Cmds for the following resources: %s)",
			resource, strings.Join(allResources, ", "))
	}
	return &cmd, nil
}
func (s *Stopper) AllCmdsByResource(ctx context.Context) (map[string]v1alpha1.Cmd, error) {
	cmds, err := s.Cli.TiltV1alpha1().Cmds().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error watching tilt sessions")
	}

	ret := make(map[string]v1alpha1.Cmd, len(cmds.Items))
	for _, cmd := range cmds.Items {
		ret[cmd.ObjectMeta.Annotations[v1alpha1.AnnotationManifest]] = cmd
	}
	return ret, nil
}

func (s *Stopper) StopCmd(ctx context.Context, cmd *v1alpha1.Cmd) error {
	// fmt.Printf("🤖 deleting previous version of command %q\n", cmd.Name)
	// err := s.Cli.TiltV1alpha1().Cmds().Delete(ctx, cmd.Name, v1.DeleteOptions{})
	// if err != nil {
	// 	return errors.Wrapf(err, "deleting existing cmd %q", cmd.Name)
	// }
	if cmd.Status.Running == nil {
		return fmt.Errorf("cannot stop Cmd %q because it's not currently running (Status: %+v)", cmd.Name, cmd.Status)
	}

	pid := int(cmd.Status.Running.PID)
	proc, err := os.FindProcess(pid)
	if err != nil {
		return errors.Wrapf(err, "finding process %d (for cmd %q)", pid, cmd.Name)
	}

	// TODO: should probably send a SIGTERM first, and only kill if nothing happens
	err = proc.Kill()
	if err != nil {
		return errors.Wrapf(err, "killing process %d (for cmd %q)", pid, cmd.Name)
	}

	return nil
}
