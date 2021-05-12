package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tilt-dev/tilt-ci-status/pkg/stop"
)

func parseArgs() (resource string, err error) {
	if len(os.Args) != 2 {
		return "",
			fmt.Errorf("need exactly one arg (resource name to stop). Got args: %v", os.Args[1:])
	}
	return os.Args[1], nil
}
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resource, err := parseArgs()
	if err != nil {
		return err
	}

	s, err := stop.NewStopper()
	if err != nil {
		return err
	}

	return s.StopResource(ctx, resource)
}

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
