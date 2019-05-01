package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gtrox115/audioprofile/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
