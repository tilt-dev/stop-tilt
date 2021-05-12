package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tilt-dev/tilt-ci-status/pkg/stop"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := stop.NewStopper()
	if err != nil {
		return err
	}

	resourceNames := os.Args[1:]
	if len(resourceNames) == 0 {
		return s.StopTiltSession(ctx)
	}

	return s.StopResources(ctx, resourceNames)
}

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
