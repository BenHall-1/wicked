package cmds

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/break_role"
	"github.com/benhall-1/wicked/internal/pkg/discord/cmds/test"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Commands = map[string]interfaces.Command{
	"createbreakrolemenu": break_role.Command{},
	"test":                test.Command{},
}
