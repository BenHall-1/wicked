package cmds

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/break_role"
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/create_info_menu"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Commands = map[string]interfaces.Command{
	"createbreakrolemenu": break_role.Command{},
	"createinfomenu":      create_info_menu.Command{},
}
