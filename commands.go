package main

import (
	"os"

	"github.com/jrasell/levant/command"
	"github.com/jrasell/levant/version"
	"github.com/mitchellh/cli"
)

// Commands returns the mapping of CLI commands for Levant. The meta parameter
// lets you set meta options for all commands.
func Commands(metaPtr *command.Meta) map[string]cli.CommandFactory {
	if metaPtr == nil {
		metaPtr = new(command.Meta)
	}

	meta := *metaPtr
	if meta.UI == nil {
		meta.UI = &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		}
	}

	return map[string]cli.CommandFactory{

		"deploy": func() (cli.Command, error) {
			return &command.DeployCommand{
				Meta: meta,
			}, nil
		},
		"dispatch": func() (cli.Command, error) {
			return &command.DispatchCommand{
				Meta: meta,
			}, nil
		},
		"render": func() (cli.Command, error) {
			return &command.RenderCommand{
				Meta: meta,
			}, nil
		},
		"scale-in": func() (cli.Command, error) {
			return &command.ScaleInCommand{
				Meta: meta,
			}, nil
		},
		"scale-out": func() (cli.Command, error) {
			return &command.ScaleOutCommand{
				Meta: meta,
			}, nil
		},
		"version": func() (cli.Command, error) {
			ver := version.Version
			rel := version.VersionPrerelease

			if rel == "" && version.VersionPrerelease != "" {
				rel = "dev"
			}
			return &command.VersionCommand{
				Version:           ver,
				VersionPrerelease: rel,
				UI:                meta.UI,
			}, nil
		},
	}
}
