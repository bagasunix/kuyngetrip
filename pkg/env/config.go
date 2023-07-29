package env

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Configs struct {
	Env             string
	DbHost          string
	DbPort          string
	DbName          string
	DbUsername      string
	DbPassword      string
	SmtpHost        string
	SmtpPort        string
	SmtpFrom        string
	SmtpUser        string
	SmtpPass        string
	ApiVersion      string
	Port            string
	Timezone        string
	MaxIdleConns    string
	MaxOpenConns    string
	ConnMaxLifetime string
	JwtKey          string
}

var env = map[string]map[string]string{
	"DB_HOST": {
		"Key":          "DbHost",
		"DefaultValue": "localhost",
	},
	"DB_PORT": {
		"Key":          "DbPort",
		"DefaultValue": "3306",
	},
	"DB_NAME": {
		"Key":          "DbName",
		"DefaultValue": "kuyngetrip",
	},
	"DB_USER": {
		"Key":          "DbUsername",
		"DefaultValue": "root",
	},
	"DB_PASSWORD": {
		"Key":          "DbPassword",
		"DefaultValue": "root",
	},
	"ENV": {
		"Key":          "Env",
		"DefaultValue": "development",
	},
	"SMTP_HOST": {
		"Key":          "SmtpHost",
		"DefaultValue": "mailjet",
	},
	"SMTP_PORT": {
		"Key":          "SmtpPort",
		"DefaultValue": "587",
	},
	"SMTP_FROM": {
		"Key":          "SmtpFrom",
		"DefaultValue": "noreply@notification.kuyngetrip.id",
	},
	"SMTP_USER": {
		"Key":          "SmtpUser",
		"DefaultValue": "noreply@notification.kuyngetrip.id",
	},
	"SMTP_PASS": {
		"Key":          "SmtpPass",
		"DefaultValue": "-",
	},
	"API_VERSION": {
		"Key":          "ApiVersion",
		"DefaultValue": "/api/v1",
	},
	"RUN_PORT": {
		"Key":          "Port",
		"DefaultValue": "8088",
	},
	"TIME_ZONE": {
		"Key":          "Timezone",
		"DefaultValue": "Asia/Jakarta",
	},
	"MAX_IDLE_CONNS": {
		"Key":          "MaxIdleConns",
		"DefaultValue": "10",
	},
	"MAX_OPEN_CONNS": {
		"Key":          "MaxOpenConns",
		"DefaultValue": "100",
	},
	"CONN_MAX_LIFE_TIME": {
		"Key":          "ConnMaxLifetime",
		"DefaultValue": "5m",
	},
	"JWTKEY_SECRET": {
		"Key":          "JwtKey",
		"DefaultValue": "Str0b3r1_T@g!h4n",
	},
}

func LoadEnv() (*Configs, error) {
	logs, _ := zap.NewProduction()
	defer logs.Sync()

	var conf Configs

	err := godotenv.Load()

	if err != nil {
		logs.Fatal(err.Error())
	}

	for key, element := range env {
		if os.Getenv(key) != "" {
			reflect.ValueOf(&conf).Elem().FieldByName(element["Key"]).SetString(os.Getenv(key))
		} else {
			reflect.ValueOf(&conf).Elem().FieldByName(element["Key"]).SetString(element["DefaultValue"])
		}

	}
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
