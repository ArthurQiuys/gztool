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
common = "{{unemeta_backend_dev:uneune202}}{{localhost:3069}}"
modules = [
    "{{user}}{{unemeta_users}}{{unemeta_backend_dev:uneune202}}{{localhost:3069}}",
]

[db.test]
common = "{{unemeta_backend_dev:UneUne202.}}{{34.85.98.88:3306}}"
modules = [
    "{{console}}{{unemeta_console}}{{unemeta_admin:UneUne202.}}{{34.85.98.88:3306}}",
	"{{user}}{{unemeta_users}}{{unemeta_backend_dev:UneUne202.}}{{34.85.98.88:3306}}",
	"{{market}}{{unemeta_market}}{{unemeta_backend_dev:UneUne202.}}{{34.85.98.88:3306}}",
	"{{market}}{{unemeta_market}}{{unemeta_backend_dev:UneUne202.}}{{34.84.252.50:3306}}"
]

[y_api]
server = "http://35.209.215.76:40001"
modules = [
    "{{console}}{{d37b8b56ed439017efd4bcbe08f1df60d84f2ff127dac751d4f6d88fd27952ad}}",
	"{{user}}{{0eda057be16656f12676d1f00f34ded8eaedfcfe0e947ca45cd5d97a69b86a27}}",
	"{{integral}}{{ff2d5df9deac8567259aa978e9c534a6e71038e12a606d7635b1d294eda26b1b}}",
	"{{systeminfo}}{{74146e49063c48fe6cb2a5208aa1292f18d4651a21d55b944fbfba1eb6d5323c}}",
	"{{market}}{{968467dca328d146a1bdf06ad8daf098d9a15f5cc48d38fbad0e90aad6f7be58}}",
	"{{transaction}}{{c9e1d472f549da98d2368332d1b6fe650e2225f17b4095912b612856d60d94cc}}",
]

[read_me]
key = "rdme_xn8s9h17429f437a5a8a97371d49cab3b9a458698f43f045abf7b46601fe0195ee3c30"
modules = [
    "{{console}}{{639c8750455b32000f346c98}}",
	"{{integral}}{{639c895f9fb014026fe471c2}}",
	"{{market}}{{639c8981178d5606e01c47cc}}",
	"{{systeminfo}}{{639c89a9ba635200467f0e5d}}",
	"{{transaction}}{{639c89c70879c50093fb1b8d}}",
	"{{user}}{{639c89e01f22560012be3fce}}",
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