package listeners

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/shinpuru/internal/services/database"
	"github.com/zekroTJA/shinpuru/internal/util"
	"github.com/zekroTJA/shinpuru/internal/util/static"
	"github.com/zekroTJA/shinpuru/pkg/embedbuilder"
)

type ListenerMemberAdd struct {
	db database.Database
}

func NewListenerMemberAdd(container di.Container) *ListenerMemberAdd {
	return &ListenerMemberAdd{
		db: container.Get(static.DiDatabase).(database.Database),
	}
}

func (l *ListenerMemberAdd) Handler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	autoRoleID, err := l.db.GetGuildAutoRole(e.GuildID)
	if err != nil && !database.IsErrDatabaseNotFound(err) {
		util.Log.Errorf("Failed getting autorole for guild '%s' from database: %s", e.GuildID, err.Error())
	}
	if autoRoleID != "" {
		err = s.GuildMemberRoleAdd(e.GuildID, e.User.ID, autoRoleID)
		if err != nil && strings.Contains(err.Error(), `{"code": 10011, "message": "Unknown Role"}`) {
			l.db.SetGuildAutoRole(e.GuildID, "")
		} else if err != nil {
			util.Log.Errorf("Failed setting autorole for member '%s': %s", e.User.ID, err.Error())
		}
	}

	chanID, msg, err := l.db.GetGuildJoinMsg(e.GuildID)
	if err == nil && msg != "" && chanID != "" {
		txt := ""
		if strings.Contains(msg, "[ment]") {
			txt = e.User.Mention()
		}

		msg = strings.Replace(msg, "[user]", e.User.Username, -1)
		msg = strings.Replace(msg, "[ment]", e.User.Mention(), -1)

		s.ChannelMessageSendComplex(chanID, &discordgo.MessageSend{
			Content: txt,
			Embed: embedbuilder.New().
				WithColor(static.ColorEmbedDefault).
				WithDescription(msg).
				Build(),
		})
	}
}