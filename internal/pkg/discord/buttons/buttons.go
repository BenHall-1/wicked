package buttons

import (
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/break_role"
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/info"
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/info/art_rules"
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/info/community_rules"
	"github.com/benhall-1/wicked/internal/pkg/discord/buttons/info/tour"
	"github.com/benhall-1/wicked/internal/pkg/interfaces"
)

var Buttons = map[string]interfaces.Button{
	// Break Role Buttons
	"add_break_role_button":    break_role.Add{},
	"remove_break_role_button": break_role.Remove{},
	// Info Buttons
	"info_button": info.Info{},
	// Main Menu Buttons
	"community_rules_button": community_rules.CommunityRules1{FromMainMenu: true},
	"art_rules_button":       art_rules.ArtRules1{FromMainMenu: true},
	// Community Rules Buttons
	"community_rules_1": community_rules.CommunityRules1{},
	"community_rules_2": community_rules.CommunityRules2{},
	// Art Rules Buttons
	"art_rules_1": art_rules.ArtRules1{},
	"art_rules_2": art_rules.ArtRules2{},
	// Tour Buttons
	"get_started": tour.GetStarted{},
	"roles":       tour.Roles{},
}
