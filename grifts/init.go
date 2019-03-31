package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/control/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
