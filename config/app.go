package config

type Currency string

const (
	RMB Currency = "元"
	USD Currency = "$"
)

type Local string

const (
	CN Local = "zh-CN"
	EN Local = "en"
)

const AppName = "Frank's Shop"

var CurrentCurrency = RMB
var CurrentLocal = CN
