module examples

go 1.15

replace (
	github.com/nobonobo/spago v1.0.7 => ../../spago
	github.com/nobonobo/spago-spectre v0.0.0 => ../
)

require (
	github.com/mattn/cho v0.0.6 // indirect
	github.com/nobonobo/spago v1.0.7
	github.com/nobonobo/spago-spectre v0.0.0
)
