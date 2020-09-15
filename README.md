# Etpmls-Admin-Server

[Ecology|Plug-in development|i18n globalization]

[生态|插件式开发|i18n国际化]





## Configuration 配置

此项目为Etpmls-Admin后端源码

This project is the Etpmls-Admin backend source code



1. Copy .env.example to .env

    将.env.example复制到.env

2. Copy storage/config/app.yaml.example to storage/config/app.yaml

    将storage / config / app.yaml.example复制到storage / config / app.yaml

3. Copy storage/config/app_debug.yaml.example to storage/config/app_debug.yaml

   将storage / config / app_debug.yaml.example复制到storage / config / app_debug.yaml

And configure them

并且配置它们

4
```shell script
go mod vendor
```


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