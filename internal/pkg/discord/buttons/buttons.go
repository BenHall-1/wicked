package buttons

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/break_role"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Buttons = map[string]interfaces.Button{
	"add_break_role_button":    break_role.Add{},
	"remove_break_role_button": break_role.Remove{},
}
