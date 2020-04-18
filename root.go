package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/bl3-save/cmd"
	"github.com/spf13/cobra"
)

func rootFactory(s *discordgo.Session, e *discordgo.MessageCreate) *cobra.Command {
	c := &cobra.Command{
		Short:   "bl3-save bot",
		Example: "!convert bl3(code)",
	}
	c.AddCommand(cmd.ConvertCmd)
	return c
}
