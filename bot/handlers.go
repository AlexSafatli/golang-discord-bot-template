package bot

import (
	"github.com/bwmarrin/discordgo"
)

func getMainChannelIDForGuild(b *Bot, guildID string) string {
	var id string
	if id, ok := b.mainGuildChannelIDs[guildID]; ok {
		return id
	}
	var lowestPos = -1
	gc, err := b.Session.GuildChannels(guildID)
	if err != nil {
		return id
	}
	for _, c := range gc {
		if c.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		if lowestPos == -1 || c.Position < lowestPos {
			lowestPos = c.Position
			id = c.ID
		}
	}
	if len(id) > 0 {
		b.mainGuildChannelIDs[guildID] = id // cache the channel ID
	}
	return id
}

// OnGuildChannelCreateHandler is another specific use-case handler function that invalidates the bot's channel ID cache if the channel in the cache is deleted
func OnGuildChannelCreateHandler(b *Bot) func(*discordgo.Session, *discordgo.ChannelCreate) {
	return func(s *discordgo.Session, cc *discordgo.ChannelCreate) {
		if len(cc.GuildID) == 0 || cc.Channel == nil {
			return // check if the guild and channels are defined first
		}
		if id, ok := b.mainGuildChannelIDs[cc.GuildID]; ok {
			if id == cc.Channel.ID {
				delete(b.mainGuildChannelIDs, cc.GuildID) // invalidate cache
			} else {
				newMainChannel := getMainChannelIDForGuild(b, cc.GuildID)
				if newMainChannel != id {
					b.mainGuildChannelIDs[cc.GuildID] = newMainChannel // update
				}
			}
		}
	}
}
