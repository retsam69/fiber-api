package gmf_pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/phuslu/log"
)

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
