package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

var args = struct {
	list    bool
	host    bool
	version bool
	config  string
}{
	false,
	false,
	false,
	"aws-inventory.ini",
}

func init() {
	flag.BoolVar(&args.list, "list", false, "Print Ansible formatted inventory")
	flag.BoolVar(&args.host, "host", false, "no-op since all information is given via --list")
	flag.BoolVar(&args.version, "v", false, "Print version")
	flag.StringVar(&args.config, "c", args.config, "Configuration filename")
}

func main() {
	flag.Parse()
	switch {
	case args.version:
		fmt.Printf("%s v%s\n", appName, appVersion)
		return
	case args.host:
		fmt.Fprint(os.Stdout, "{}")
		return
	case args.list:
		cfg, err := getConfig(args.config)
		if err != nil {
			die("Error reading configuration file:\n%s\n", err)
		}
		printList(cfg)
		return
	}

	flag.PrintDefaults()
}

func printList(cfg *configuration) {
	auth, err := aws.GetAuth(cfg.AccessKey, cfg.SecretKey)
	if err != nil {
		die("Error creating AWS auth:\n%s\n", err)
	}

	e := ec2.New(auth, aws.EUWest)
	instances, err := e.Instances([]string{}, nil)
	if err != nil {
		die("Error fetching EC2 instances:\n%s\n", err)
	}
	inv, err := newInventory(instances)
	if err != nil {
		die("Error creating inventory from EC2 instances:\n%s\n", err)
	}
	invJSON, err := inv.toJSON()
	if err != nil {
		die("Error generatin inventory JSON:\n%s\n", err)
	}
	os.Stdout.Write(invJSON)
}

func die(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
