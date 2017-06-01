package main

import (
	"fmt"
	"github.com/appbaseio/abc/appbase/app"
)

// runApps runs `apps` command
func runApps(args []string) error {
	flagset := baseFlagSet("apps")
	flagset.Usage = usageFor(flagset, "abc apps")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	args = flagset.Args()

	switch len(args) {
	case 0:
		return app.ShowUserApps()
	default:
		fmt.Println("No such option. See --help")
	}
	return nil
}
