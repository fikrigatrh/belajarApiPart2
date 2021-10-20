package models

import "gorm.io/gorm"

// ServerConfig untuk env
type ServerConfig struct {
	AppName        string `env:"APP_NAME"`
	AppPort        string `env:"APP_PORT"`
	LogLevel       string `env:"LOG_LEVEL"`
	Environment    string `env:"ENVIRONMENT"`
	JWTSecret      string `env:"JWT_SECRET"`
	RedisAddress   string `env:"REDIS_ADDRESS"`
	DBUsername     string `env:"DB_USERNAME"`
	DBPassword     string `env:"DB_PASSWORD"`
	DBHost         string `env:"DB_HOST"`
	DBPort         string `env:"DB_PORT"`
	DBName         string `env:"DB_NAME"`
	MinioEndpoint  string `env:"MINIO_ENDPOINT"`
	MinioAccessKey string `env:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `env:"MINIO_SECRET_KEY"`
	MinioRegion    string `env:"MINIO_REGION"`
	MinioBucket    string `env:"MINIO_BUCKET"`
	UrlJsonUser    string `env:"URL_JSON_USER"`
	UsernameSu     string `env:"USERNAME_SU"`
	PasswordSu     string `env:"PASSWORD_SU"`
}

// Auth fungsi : untuk menyimpan tiap User yang login
type Auth struct {
	gorm.Model
	// Authentication Unique ID
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
	Username string `gorm:";not null;" json:"username"`
	RoleId   string `json:"role_id"`
	IdCust   int    `json:"id_cust" gorm:"-"`
	IdOffice int    `json:"id_office" gorm:"-"`
	RoleName string `json:"role_name"`
}

// AuthDetail Isi atau Payload dari JWT
type AuthDetail struct {
	AuthUUID string
	Username string
}

type JwtModel struct {
	Token    string `json:"token"`
	IdCust   int    `json:"id_cust,omitempty"`
	IdOffice int    `json:"id_office,omitempty"`
}

const (
	Admin          = 1
	GeneralSupport = 2
	Accounting     = 3
	Customer       = 4
)

const (
	DIREJECT_GS = 2
	DITERUSKAN  = 3
	DIREJECT_AC = 4
	DISETUJUI   = 5
)
