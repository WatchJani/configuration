package cli

import (
	"flag"
)

func SetupFlags() string {
	config_path := flag.String("path", "../configuration.json", "use to set up app configuration")

	flag.Parse()

	return *config_path
}
