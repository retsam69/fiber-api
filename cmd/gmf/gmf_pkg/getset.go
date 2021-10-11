package gmf_pkg

import (
	"fmt"
	"strings"

	"github.com/phuslu/log"
)

var Lines []string

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
