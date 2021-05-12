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

func (s *Stopper) Kill(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return errors.Wrapf(err, "finding process %d", pid)
	}

	// TODO: should probably send a SIGTERM first, and only kill if nothing happens
	err = proc.Kill()
	if err != nil {
		return errors.Wrapf(err, "killing process %d", pid)
	}
	return nil
}

func (s *Stopper) StopResources(ctx context.Context, resources []string) error {
	cmds, err := s.CmdsForResources(ctx, resources)
	if err != nil {
		return errors.Wrap(err, "getting cmd for resource")
	}

	for _, cmd := range cmds {
		err = s.StopCmd(cmd)
		if err != nil {
			fmt.Printf("🚨 error stopping Cmd %q: %v\n", cmd.Name, err)
		} else {
			fmt.Printf("✅ successfully stopped Cmd %q\n", cmd.Name)
		}
	}
	return nil
}

func (s *Stopper) CmdsForResources(ctx context.Context, resources []string) ([]*v1alpha1.Cmd, error) {
	allCmds, err := s.AllCmdsByResource(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting all Cmds")
	}

	var allResourcesWithCmds []string
	for resource := range allCmds {
		allResourcesWithCmds = append(allResourcesWithCmds, resource)
	}

	var ret []*v1alpha1.Cmd
	var notFound []string
	for _, resource := range resources {
		cmd, ok := allCmds[resource]
		if !ok {
			notFound = append(notFound, resource)
			continue
		}
		ret = append(ret, &cmd)
	}

	if len(notFound) > 0 {
		fmt.Printf("🚨 no Cmd(s) found for resource(s): %s\n\t(found Cmds for the following resources: %s)\n",
			strings.Join(notFound, ", "), strings.Join(allResourcesWithCmds, ", "))
	}

	if len(ret) == 0 {
		return nil, fmt.Errorf("no Cmds found to stop")
	}
	return ret, nil
}

func (s *Stopper) AllCmdsByResource(ctx context.Context) (map[string]v1alpha1.Cmd, error) {
	cmds, err := s.Cli.TiltV1alpha1().Cmds().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error listing Cmds")
	}

	ret := make(map[string]v1alpha1.Cmd, len(cmds.Items))
	for _, cmd := range cmds.Items {
		ret[cmd.ObjectMeta.Annotations[v1alpha1.AnnotationManifest]] = cmd
	}
	return ret, nil
}

func (s *Stopper) StopCmd(cmd *v1alpha1.Cmd) error {
	if cmd.Status.Running == nil {
		return fmt.Errorf("cannot stop Cmd because it's not currently running (Status: %+v)", cmd.Status)
	}

	pid := int(cmd.Status.Running.PID)
	return s.Kill(pid)
}

func (s *Stopper) StopTiltSession(ctx context.Context) error {
	// TODO: multiple/differently-named sessions?
	session, err := s.Cli.TiltV1alpha1().Sessions().Get(ctx, "Tiltfile", v1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "getting session `Tiltfile`")
	}

	pid := int(session.Status.PID)
	return s.Kill(pid)
}
