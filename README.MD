# API Service

# By `gofiber/fiber` Framework

## Step 1: Get Project

```shell
cd __YOUR_PROJECT__

git clone https://gitlab.com/indev-moph/fiber-api.git . \
&& ./init.sh

# OR clone with ssh
git clone git@gitlab.com:indev-moph/fiber-api.git . \
&& ./init.sh

```

## Set `go/bin` path in `$PATH` Envolopment if your not set

Check go path

```shell
go env GOPATH
echo $PATH
```

Set `$PATH`

```shell
export PATH=$PATH:$GOPATH/bin
```

## Step 2: Uncomment

1. file: `main.go`

```go
func main() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	logger.SetDefaultlogger()
	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// ---- Plaase Uncommant ----
	// load config or etc.
	loader.Init()

	// start http server
	// Serv(controller.Init, route.Init) // <---- Uncommend Line
}
```


### Config With `ENV`

```dotenv
APP.DEV=false
APP.BASEURL=https://${DOMAIN}:${PORT}/${PREFIX}
APP.PORT=${PORT}

```

### Add git confg in `~/.gitconfig`

```
[alias]
  cpush = "!f() { git commit -am \"$(echo $@)\" && git push; }; f" "$([ -z \"$(git status --porcelain | grep '^\\s\\?A\\s')\" ] || git status --porcelain | grep '^\\s\\?A\\s' | awk 'BEGIN{print \"Added:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?C\\s')\" ] || git status --porcelain | grep '^\\s\\?C\\s' | awk 'BEGIN{print \"Copied:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?D\\s')\" ] || git status --porcelain | grep '^\\s\\?D\\s' | awk 'BEGIN{print \"Deleted:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?M\\s')\" ] || git status --porcelain | grep '^\\s\\?M\\s' | awk 'BEGIN{print \"Modified:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?R\\s')\" ] || git status --porcelain | grep '^\\s\\?R\\s' | awk 'BEGIN{print \"Renamed:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')";
```
