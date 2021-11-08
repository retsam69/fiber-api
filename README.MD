# API Service

# By `gofiber/fiber` Framework

## Step 1: Get Project

```shell
cd __YOUR_PROJECT__

git clone https://gitlab.com/indev-moph/fiber-api.git .
./init.sh
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

```
// ---- Plaase Uncommant ----
// // load config or etc.
// cmd_api.Init()
// // start http server
// StartFiberServer(route.Init)
```

2. file: `./internal/route/route.go`

```
// // ---- Plaase Uncommant ----
// rg := app.Group(UrlPrefix)
// for i, v := range controller.RegisRoutes {
// 	log.Info().Msgf("Registor Endpoint: %d", i+1)
// 	v(rg)
// }
```

### Config Default file: `default.yml`

```yaml
app:
  dev: false
  baseurl: https://example.com/api/v2
auth:
  basic: user:pass
db:
  mysql:
    -
```

### Config With `ENV`

```dotenv
APP__DEV=true
APP__BASEURL=https://example.com/api/v2

```

### Add git confg in `~/.gitconfig`

```
[alias]
  cpush = "!f() { git commit -am \"$(echo $@)\" && git push; }; f" "$([ -z \"$(git status --porcelain | grep '^\\s\\?A\\s')\" ] || git status --porcelain | grep '^\\s\\?A\\s' | awk 'BEGIN{print \"Added:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?C\\s')\" ] || git status --porcelain | grep '^\\s\\?C\\s' | awk 'BEGIN{print \"Copied:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?D\\s')\" ] || git status --porcelain | grep '^\\s\\?D\\s' | awk 'BEGIN{print \"Deleted:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?M\\s')\" ] || git status --porcelain | grep '^\\s\\?M\\s' | awk 'BEGIN{print \"Modified:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')" "$([ -z \"$(git status --porcelain | grep '^\\s\\?R\\s')\" ] || git status --porcelain | grep '^\\s\\?R\\s' | awk 'BEGIN{print \"Renamed:\"} {printf \"%s%s\",(NR>1?\", \":\"\"), $2} END{print \"\"}')";
```
