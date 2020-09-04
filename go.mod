module github.com/isfk/aio

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/billcobbler/casbin-redis-watcher/v2 v2.0.0-20200828043318-3301e9a5ebc2
	github.com/casbin/casbin/v2 v2.11.2
	github.com/casbin/gorm-adapter/v3 v3.0.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/douyu/jupiter v0.2.4
	github.com/go-cache/cache v1.1.0
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/log/v7 v7.0.2
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.3.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.0
)
