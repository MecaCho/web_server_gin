package config

import (
	"flag"
	"github.com/golang/glog"
)

// Config server handle configs
type Config struct {
	EmailConn            string
	EmailSuffix          string
	DBConn               string
	Debug                bool
	ServerIP             string
	ServerPort           int
	CertData             string
	KeyData              string
	APIDocsPath          string
	CasbinPolicyFilePath string
}

// InitConfig init configs
func InitConfig() (APIConfig Config) {
	flag.StringVar(&APIConfig.DBConn, "dbConnect", "root:xxxxxxx@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai", "db connect, use mysql")
	flag.StringVar(&APIConfig.ServerIP, "serverIp", "127.0.0.1", "server api ip")
	flag.IntVar(&APIConfig.ServerPort, "serverPort", 80, "server api port")
	flag.StringVar(&APIConfig.CertData, "certData", "server.cert", "server cert path")
	flag.StringVar(&APIConfig.KeyData, "keyData", "server.key", "server key path")
	flag.BoolVar(&APIConfig.Debug, "debug", false, "run server in debug mode")
	flag.StringVar(&APIConfig.APIDocsPath, "apiDocsPath", "/var/web_backend/docs/api_reference", "api docs html files path")
	flag.StringVar(&APIConfig.CasbinPolicyFilePath, "rolePolicyFile", "./auth/keymatch_policy.csv", "casbin policy file path")
	flag.StringVar(&APIConfig.EmailConn, "emailConnect", "xxxxx:xxxxxx@smtp.xxxx.com", "email sender configs")
	flag.StringVar(&APIConfig.EmailSuffix, "emailSuffix", "xxxxx.com", "email address suffix")

	flag.Parse()
	glog.Infoln("Init config ...")
	return
}
