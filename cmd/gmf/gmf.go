package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/attapon-th/phuslulogger"
	"github.com/phuslu/log"
)

const (
	AppName         = "GoMakefile"
	AppnDescription = `for make file go project.`
)

var (
	makefile = "./Makefile"
	Lines    []string
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
	Readlines(makefile)
	switch cmd {
	case "h":
		help()
	case "i":
		IncrementVersion()
		WriteLines(makefile)
	case "b":
		// _, oldBuild := GetName("build")
		b := GitBuild()
		SetName("build", b)
		WriteLines(makefile)
	case "v":
		_, l := GetVersion()
		log.Info().Msgf("Version: %s", l)
	default:
		help()

	}
}

func help() {
	log.Info().Str("command", "h").Msg("Help")
	log.Info().Str("command", "v").Msg("Show version")
	log.Info().Str("command", "i").Msg("Increment minor version")
	log.Info().Str("command", "b").Msg("Build hash")
}

func GetName(name string) (int, string) {
	name = strings.ToUpper(name)
	for index, l := range Lines {
		if strings.HasPrefix(l, name+"=") {
			return index, strings.TrimPrefix(l, name+"=")
		}
	}
	return -1, ""
}

func SetName(name, value string) {
	name = strings.ToUpper(name)
	i, _ := GetName(name)
	if i < 0 {
		return
	}
	Lines[i] = fmt.Sprintf("%s=%s", name, value)
	log.Info().Str(name, value).Msgf("updated")
}

func OutputCmd(name string, cmd ...string) string {
	c := exec.Command(name, cmd...)
	s, err := c.Output()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	return string(s)
}

func GitBuild() string {
	return OutputCmd("git", "rev-parse", "HEAD")
}

func Readlines(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Lines = append(Lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func WriteLines(filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	defer file.Close()
	for i, l := range Lines {
		if _, err = file.WriteString(fmt.Sprintf("%s\n", l)); err != nil {
			log.Error().Int("index", i).Err(err).Msg("")
		}
	}

}

func GetVersion() (int, string) {
	for index, l := range Lines {
		if strings.HasPrefix(l, "VERSION=") {
			return index, strings.TrimPrefix(l, "VERSION=")
		}
	}
	return -1, ""
}

func SetVersion(version string) {
	i, _ := GetVersion()
	if i < 0 {
		return
	}
	Lines[i] = fmt.Sprintf("VERSION=%s", version)
	log.Info().Str("VERSION", version).Msg("updated")
}

func IncrementVersion() {
	i, v := GetVersion()
	if i < 0 {
		return
	}
	log.Info().Str("VERSION", v).Msg("older")
	sp := strings.Split(v, ".")
	if len(sp) <= 0 {
		return
	}
	if MinorVersion, err := strconv.ParseInt(sp[len(sp)-1], 10, 64); err != nil {
		log.Error().Err(err).Msg("")
	} else {
		MinorVersion++
		sp[len(sp)-1] = fmt.Sprint(MinorVersion)
	}
	Lines[i] = fmt.Sprintf("VERSION=%s", strings.Join(sp, "."))
	log.Info().Str("VERSION", strings.Join(sp, ".")).Msg("updated")
}
