package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gtrox115/profileapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
