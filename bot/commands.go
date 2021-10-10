package bot

import (
	"github.com/AlexSafatli/golang-discord-bot-template/chat"

	"github.com/bwmarrin/discordgo"
)

// AboutMessageCommand takes a created message and returns an About embed message
func AboutMessageCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	chat.SendRawEmbedMessage(s, m.ChannelID, chat.GetRawAboutEmbedMessage(s))
}

// AboutSlashCommand returns an About embed message for a slash command
func AboutSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	chat.SendInteractionRawEmbedForAction(s, i, chat.GetRawAboutEmbedMessage(s), nil)
}
