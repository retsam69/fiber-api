package gmf_pkg

import (
	"os/exec"

	"github.com/phuslu/log"
)

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
