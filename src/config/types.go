package config

const (
	DbLocal SqlEnvEnum = iota
	DbTest
)

const (
	DocNone DocEnum = iota
	DocAll
	DocYApi
	DocReadMe
)

const (
	GoCtlSwaggerGit  = "https://github.com/libvirgo/goctl-swagger.git"
	GoCtlTemplateGit = "https://git.unemeta.com/Backstage/backend_template.git"
	GoCtlInstall     = "github.com/zeromicro/go-zero/tools/goctl@v1.4.0"
)

type SqlEnvEnum int64

type DocEnum int64
