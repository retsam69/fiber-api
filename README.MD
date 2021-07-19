# API Service
# By `gofiber/fiber` Framework

## Get Project
### Step 1
```shell
export MYPROJECT=<your-project-name>
git clone https://gitlab.com/indev-moph/fiber-api.git $MYPROJECT
cd $MYPROJECT
rm -f go.mod
rm -f go.sum
rm -rf .git
```

### Step 2
```shell
go mod init <go project name>
make mod
```

## OR
    curl -Ssl 