package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
		conf := askForConfirmation("Are you sure you want to kill your Tilt session? (Invoke this CLI with args to stop individual resources.)")
		if !conf {
			return nil
		}
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

func askForConfirmation(s string) bool {
	// With thanks to https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [Y/n]: ", s)

	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	response = strings.ToLower(strings.TrimSpace(response))

	if response == "n" || response == "no" {
		return false
	}

	return true
}
