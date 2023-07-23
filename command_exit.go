package main

import "os"

func commandExit(cfg *config, commandArgs []string) error {
	os.Exit(0)
	return nil
}
