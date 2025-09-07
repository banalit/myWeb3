module task4 // 保持你现有的模块名

go 1.19 // 保持你的 Go 版本

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/jmoiron/sqlx v1.4.0
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.30.3
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	// 强制指定兼容旧版本 Go 的依赖版本
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

// 替换依赖以避免使用需要高版本 Go 的包
replace github.com/go-playground/validator/v10 => github.com/go-playground/validator/v10 v10.11.2

replace github.com/pelletier/go-toml/v2 => github.com/pelletier/go-toml/v2 v2.0.6

replace github.com/ugorji/go/codec => github.com/ugorji/go/codec v1.2.9
