package modals

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/modals/break_role"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Modals = map[string]interfaces.Modal{
	"break_role_modal": break_role.Modal{},
}
