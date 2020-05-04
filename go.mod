module group_shuffle_gui

go 1.14

require (
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/valyala/fasttemplate v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/text v0.3.2 // indirect
	local.packages/api v0.0.0
	local.packages/model v0.0.0 // indirect
	local.packages/pkg v0.0.0 // indirect
)

replace local.packages/api => ./api

replace local.packages/model => ./model

replace local.packages/pkg => ./pkg
