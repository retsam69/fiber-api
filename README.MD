# API Service
# By `gofiber/fiber` Framework

## Get Project
### Step 1
```shell
git clone https://gitlab.com/indev-moph/fiber-api.git YOUR_PROJECT
```

### Step 2
```shell
cd YOUR_PROJECT
./init
```

### Config BY ENV

> ENV Prefix: `APP_`

| ENV Name                   | Type   | Descripttion                              |
| -------------------------- | ------ | ----------------------------------------- |
| APP_DEV                    | Bool   | สำหรับการ dev mode                         |
| APP_LISTEN                 | String | Server Listener `request`                 |
| APP_PREFIX                 | String | URL Prefix or BaseURL                     |
| USER_{username}={password} | String | Set username and password for `BasicAuth` |

