package buttons

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/break_role"
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/info"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Buttons = map[string]interfaces.Button{
	// Break Role Buttons
	"add_break_role_button":    break_role.Add{},
	"remove_break_role_button": break_role.Remove{},
	// Info Buttons
	"info_button":  info.Info{},
	"rules_button": info.Rules{},
}
