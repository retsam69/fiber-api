package main

import (
	"os"

	"github.com/attapon-th/phuslulogger"
	"github.com/phuslu/log"
	"gitlab.com/indev-moph/fiber-api/cmd/gmf/gmf_pkg"
)

const (
	AppName         = "GoMakefile"
	AppnDescription = `for make file go project.`
)

var (
	makefile = "./Makefile"
)

func init() {
	log.DefaultLogger = phuslulogger.GetLoggerConsole(0, 1)
}

func main() {
	log.Info().Str("Name", AppName).Str("Description", AppnDescription).Msg("")
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	gmf_pkg.Readlines(makefile)
	switch cmd {
	case "h":
		help()
	case "i":
		gmf_pkg.IncrementVersion()
		gmf_pkg.WriteLines(makefile)
	case "b":
		// _, oldBuild := gmf_pkg.GetName("build")
		b := gmf_pkg.GitBuild()
		gmf_pkg.SetName("build", b)
		gmf_pkg.WriteLines(makefile)
	case "v":
		_, l := gmf_pkg.GetVersion()
		log.Info().Msgf("Version: %s", l)
	default:
		help()

	}
}

func help() {
	log.Info().Str("command", "h").Msg("Help")
	log.Info().Str("command", "v").Msg("Show version")
	log.Info().Str("command", "vi").Msg("Increment minor version")
	log.Info().Str("command", "b").Msg("Build hash")
}
