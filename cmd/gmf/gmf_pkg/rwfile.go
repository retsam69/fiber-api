package gmf_pkg

import (
	"bufio"
	"fmt"
	"os"

	"github.com/phuslu/log"
)

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
