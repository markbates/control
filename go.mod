module github.com/markbates/control

go 1.12

require (
	github.com/gobuffalo/buffalo v0.14.2
	github.com/gobuffalo/envy v1.6.15
	github.com/gobuffalo/mw-csrf v0.0.0-20190129204204-25460a055517 // indirect
	github.com/gobuffalo/mw-paramlogger v0.0.0-20190224201358-0d45762ab655
	github.com/gobuffalo/packr/v2 v2.0.9
	github.com/gobuffalo/suite v2.6.1+incompatible
	github.com/jackc/pgx v3.3.0+incompatible // indirect
	github.com/markbates/going v1.0.3 // indirect
	github.com/markbates/oncer v0.0.0-20181203154359-bf2de49a0be2
	github.com/markbates/portmidi v0.0.0-20190401141759-18c24bff228d
	github.com/pkg/errors v0.8.1
	github.com/spf13/afero v1.2.1 // indirect
)

replace github.com/markbates/portmidi => ../portmidi
