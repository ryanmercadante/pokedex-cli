package cli

import "os"

func exit(_ *CliConfig, _ ...string) error {
	os.Exit(0)
	return nil
}
