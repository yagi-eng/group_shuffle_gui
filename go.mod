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
	local.packages/handlers v0.0.0
	local.packages/models v0.0.0 // indirect
	local.packages/utils v0.0.0
)

replace local.packages/handlers => ./handlers
replace local.packages/models => ./models
replace local.packages/utils => ./utils
