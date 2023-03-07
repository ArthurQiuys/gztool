# Une

une 项目内部封装关于go-zero相关的工具

仅仅适用于项目内部

封装了 `goctl` 相关的命令用于简单使用

第一次打开的时候会根据以当前目录为项目目录生成配置.

自动下载依赖: `goctl`, `goctl-swagger`以及项目模版到`$cacheDir`.

配置文件为`$HOME/.unemeta_cli.toml`

## 目录

- [安装](#安装)
- [补全](#补全)
- [更新配置](#更新配置和缓存)
- [面向对象](#object-orient-cmd)
- [面向操作](#handle-orient-cmd)
- [维护与开发](#维护与开发)
- [配置模版](#配置文件模版)


## 安装

```shell
GOPRIVATE=git.unemeta.com go install github.com/ArthurQiuys/une@latest
```

## 更新配置和缓存

该操作会尝试从旧配置文件迁移配置到新的配置内容, 以及更新补全

同时也会更新 `goctl-swagger` 和 `goctl-template`

```shell
une update
```

## 补全

根据命令行提示完成操作

```shell
une update
```

补全效果参考doc文件中的截图

## Object-Orient Cmd

面向对象的使用方法

```shell
une con run api # 项目模块名称取文件名前三个字符.
une con rpc
une con api

une con up
une con up -e=test -c=false

une con mo -e=test -c=false # 手动更新model
une con mo

une con down
une con down -e=test

une con new add_table
une con new edit_table -o # sql-migrate new and open file by datagrip

une con doc
une con yapi
une con readme
```

## Handle-Orient Cmd

面向操作的使用方法

### API

apps使用前缀匹配, 可以一次生成多个模块的`api`相关文件

`une api -- [apps]`

```shell
une api con mark user trans
# or
une api -- con mark user trans
```

### RPC

使用`goctl rpc protoc`批量生成

`une rpc -- [apps]`

```shell
une rpc con mark use # goctl rpc protoc pb/{}.proto --go_out=. --go-grpc_out=. --zrpc_out=. --home={} --style=goZero
```

### Run

```shell
une run api -- con # go run $projectDir/app/console/api/console.go
une run rpc -- mark # go run $projectDir/app/market/rpc/market.go
```

### Sql

```shell
une sql up con
une sql up con -e=test -c=false

une sql mo con
une sql mo con -e=test -c=false

une sql down con
une sql down con -e=test

une sql new con add_table -o
```

### Doc

```shell
une doc yapi
une doc readme
```

## 配置文件模版

```toml
project_dir = "/Users/arthur/Documents/code/go/backend"
cache_dir = "/Users/arthur/.cache/unemeta_cli/"

[db.local]
common = "{{key:key}}{{localhost:3069}}"
modules = [
    "{{user}}{{key}}{{key:key}}{{localhost:3069}}",
]

[db.test]
common = "{{key:key.}}{{key:3306}}"
modules = [
    "{{console}}{{unemeta_console}}{{key:key.}}{{key:3306}}",
	"{{user}}{{unemeta_users}}{{key:key.}}{{key:3306}}",
	"{{market}}{{unemeta_market}}{{key:key.}}{{key:3306}}",
	"{{market}}{{unemeta_market}}{{key:key.}}{{key:3306}}"
]

[y_api]
server = "http://35.209.215.76:40001"
modules = [
    "{{console}}{{key}}",
	"{{user}}{{key}}",
	"{{integral}}{{key}}",
	"{{systeminfo}}{{key}}",
	"{{market}}{{key}}",
	"{{transaction}}{{key}}",
]

[read_me]
key = "key"
modules = [
    "{{console}}{{key}}",
	"{{integral}}{{key}}",
	"{{market}}{{key}}",
	"{{systeminfo}}{{key}}",
	"{{transaction}}{{key}}",
	"{{user}}{{key}}",
]
```

# 维护与开发

项目目录

- cmd :: 跟 cmd 命令行相关, 使用 cobra 框架, 添加自定义命令的代码...etc
  - handle_orient :: 面向操作的相关命令
  - object_orient :: 面向对象的相关命令
  - valid :: 命令行补全的有效参数相关公共代码
  - update.go :: 更新配置, 依赖等...
  - doc.go :: 更新 yapi, readme
- src
  - app :: 对应命令的实际实现
    - app.go :: 实际实现
	- app_helper :: 方便外部调用的工具函数
  - config
    - compatibility :: 配置文件的结构, 更新配置文件需要考虑前后升级的兼容性
	- config.go :: 默认会生成的配置文件的内容
	- config_helper.go :: 方便外部调用的工具函数
	- vars.go :: 初始化的一些变量