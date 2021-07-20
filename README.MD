# API Service
# By `gofiber/fiber` Framework

## Get Project
```shell
cd __YOUR_PROJECT__

git clone https://gitlab.com/indev-moph/fiber-api.git .
./init.sh
```

### Config BY ENV

> ENV Prefix: `APP_`

| ENV Name                   | Type   | Descripttion                              |
| -------------------------- | ------ | ----------------------------------------- |
| APP_DEV                    | Bool   | สำหรับการ dev mode                         |
| APP_LISTEN                 | String | Server Listener `request`                 |
| APP_PREFIX                 | String | URL Prefix or BaseURL                     |
| USER_{username}={password} | String | Set username and password for `BasicAuth` |

