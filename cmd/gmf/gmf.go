package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	AppName         = "GoMakefile"
	AppnDescription = `for make file go project.`
)

var (
	makefile = "./Makefile"
	Lines    []string
)

func main() {
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	Readlines(makefile)
	for _, c := range cmd {
		switch c {
		case 'h':
			help()
			return
		case 'i':
			IncrementVersion()
			WriteLines(makefile)
		case 'b':
			// _, oldBuild := GetName("build")
			b := GitBuild()
			SetName("build", b)
			WriteLines(makefile)
		case 'v':
			_, l := GetVersion()
			b := GitBuild()
			fmt.Printf("Version: %s\nBuild:%s\n", l, b)
		default:
			help()

		}
	}
}

func help() {
	h := `Helper:
	h	-	Show Help
	v	-	Show Info
	i	-	Increment minor Version
	b	-	Set Build Hash
`
	fmt.Println(h)
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
	value = strings.TrimRight(value, "\n")
	Lines[i] = fmt.Sprintf("%s=%s", name, value)
	fmt.Printf("Set %s: %s\n", name, value)
}

func OutputCmd(name string, cmd ...string) string {
	c := exec.Command(name, cmd...)
	s, err := c.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	return string(s)
}

func GitBuild() string {
	return OutputCmd("git", "rev-parse", "HEAD")
}

func Readlines(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Lines = append(Lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func WriteLines(filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	for i, l := range Lines {
		if _, err = file.WriteString(fmt.Sprintf("%s\n", l)); err != nil {
			fmt.Errorf("Error Line: %d, String: %s, Error: %s", i, l, err.Error())
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
	fmt.Println("VERSION: ", version)
}

func IncrementVersion() {
	i, v := GetVersion()
	if i < 0 {
		return
	}
	fmt.Println("Old VERSION: ", v)
	sp := strings.Split(v, ".")
	if len(sp) <= 0 {
		return
	}
	if MinorVersion, err := strconv.ParseInt(sp[len(sp)-1], 10, 64); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	} else {
		MinorVersion++
		sp[len(sp)-1] = fmt.Sprint(MinorVersion)
	}
	Lines[i] = fmt.Sprintf("VERSION=%s", strings.Join(sp, "."))
	fmt.Println("New VERSION: ", strings.Join(sp, "."))
}
