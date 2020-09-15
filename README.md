# Etpmls-Admin-Server

[Ecology|Plug-in development|i18n globalization]

[生态|插件式开发|i18n国际化]





## Configuration 配置

This project is the Etpmls-Admin backend source code

此项目为Etpmls-Admin后端源码



1. Copy .env.example to .env

    将.env.example复制到.env

2. Copy storage/config/app.yaml.example to storage/config/app.yaml

    将storage / config / app.yaml.example复制到storage / config / app.yaml

3. Copy storage/config/app_debug.yaml.example to storage/config/app_debug.yaml

   将storage / config / app_debug.yaml.example复制到storage / config / app_debug.yaml

And configure them

并且配置它们

4. ```
   go mod vendor
   ```

### Configuration instructions / 配置说明

> ```
> app:
>   # [Required] Application port number (eg: "8080")
>   # [必填] 应用端口号(如："8080")
>   port: "8080"
>   # [Required] Whether to enable the verification code function(true/false), The default is true
>   # [必填] 是否开启验证码功能(true/false)，默认为true
>   captcha: true
>   # [Required] Whether to enable the registration function (true/false), The default is false
>   # [必填] 是否开启注册功能(true/false)，默认为false
>   register: false
>   # [Optional] Encrypted salt, leave it blank to generate automatically
>   # [可选]加密盐，留空则自动生成
>   key:
>   # [Required] Whether to enable the cache function, if you have installed redis, please fill in true, otherwise fill in false, The default is false
>   # [必填] 是否开启缓存功能，如果你安装了redis等，请填写true，否则填写false，默认为false
>   cache: false
>   # token expiration time(second), The default is 86400
>   # token过期时间(秒)，默认为86400
>   token-expiration-time: 86400
>   # [Required] Whether to use the http status code as the code field in the api (true/false), the default is false
>   # [必填] 是否使用http状态码作为api中的code字段(true/false)，默认为false
>   use-http-code: false
> log:
>   # [Required] The lowest level of logging (panic/fatal/error/warning/info/debug/trace), the default is info
>   # [必填] 记录日志的最低等级(panic/fatal/error/warning/info/debug/trace)，默认为info
>   level: info
>   # [Required] The following is the output format of different log levels (1/2/3), 1 is only output in the log file, 2 is only output in the console, 3 is both output in the console and output in the log file, default Is 3
>   # [必填] 下方为不同日志等级输出格式(1/2/3)，1为仅在日志文件输出，2为仅在控制台输出，3为既在控制台输出也在日志文件输出，默认为3
>   panic: 3
>   fatal: 3
>   error: 3
>   warning: 3
>   info: 3
>   debug: 3
>   trace: 3
> database:
>   # [Required] Database address (eg: "localhost")
>   # [必填] 数据库地址(如：localhost)
>   host: localhost
>   # [Required] Database port (e.g. "5432")
>   # [必填] 数据库端口(如："5432")
>   port: "5432"
>   # [Required] Database name (eg: your-database-name)
>   # [必填] 数据库名(如：your-database-name)
>   name: your-database-name
>   # [Required] Database user name (eg: your-database-user)
>   # [必填] 数据库用户名(如：your-database-user)
>   user: your-database-user
>   # [Required] Database password (e.g. your-database-password)
>   # [必填] 数据库密码(如：your-database-password)
>   password: your-database-password
>   # [Required] Database table prefix (eg: prefix_)
>   # [必填] 数据库表前缀(如：prefix_)
>   prefix: prefix_
> cache:
>   # [Required] Cache server address (eg: localhost:6379)
>   # [必填] 缓存服务器地址(如：localhost:6379)
>   address: localhost:6379
>   # [Required] Cache server password (eg: "123456")
>   # [必填] 缓存服务器密码(如："123456")
>   password: "123456"
>   # [Required] Cache server DB (such as: 1)
>   # [必填] 缓存服务器DB(如：1)
>   db: 1
> field:
>   api:
>     # [Required] The following are the custom field names returned by the API, which can be customized to adapt to different front-end frameworks (eg: msg)
>     # [必填] 以下为API返回的自定义字段名，可自定义，以适配不同的前端框架(如：msg)
>     code: code
>     message: message
>     status: status
>     data: data
>   pagination:
>     # [Required] The following is the name of the custom paging field returned by the API, which can be customized to adapt to different front-end frameworks (eg: pageNo)
>     # [必填] 以下为API返回的自定义分页字段名，可自定义，以适配不同的前端框架(如：pageNo)
>     page_no: pageNo
>     page_size: pageSize
>     count: count
> module:
>   # [Optional] The name of the registered module (such as: ["cms"]), if there is no module that needs to be registered, please keep []
>   # [可选] 注册的模块名的（如：["cms"]），如果没有需要注册的模块，请保持[]
>   name: []
> ```

## Run 运行

PostgreSQL
```shell script
go run -tags=postgresql main.go
```

MySQL/MariaDB
```shell script
go run -tags=mysql main.go
```



## Developer Manual / 开发者手册

Develop a module of your own / 开发一个属于你自己的模块

1. Create an empty folder and pull the latest EA branch / 创建一个空文件夹，并拉取最新的EA分支

   > git clone https://github.com/Etpmls/Etpmls-Admin-Server .

2. Use git to create an orphan branch (keep the original file) / 使用git创建orphan分支（保留原文件）

3. Create your own module folder under /module, and gitignore blocks all files except your development module /  在/module下创建你自己的模块文件夹，并且gitignore屏蔽除了你开发模块之外的所有文件。