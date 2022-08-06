package cmds

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/break_role"
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/event"
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/menu"
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/phrase"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Commands = map[string]interfaces.Command{
	"createbreakrolemenu": break_role.Command{},
	"menu":                menu.Command{},
	"phrase":              phrase.Command{},
	"event":               event.Command{},
}
