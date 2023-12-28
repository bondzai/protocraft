.PHONY: install

install:
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware
	go get github.com/spf13/viper
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
