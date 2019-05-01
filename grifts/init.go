package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gtrox115/audio_profile/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
