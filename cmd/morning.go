/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/jaronnie/ewsba/util"
	"github.com/jaronnie/ewsba/wesdk"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"os"
)

// morningCmd represents the morning command
var morningCmd = &cobra.Command{
	Use:   "morning",
	Short: "send morning message to your girl friend",
	Long:  `send morning message to your girl friend`,
	Run:   SendMorning,
}

var MorningTemplate = `
😘😘😘😘😘😘😘😘😘😘😘😘😘
`

type MorningData struct{}

func SendMorning(cmd *cobra.Command, args []string) {

	db := MorningData{}

	template, err := util.ParseTemplate(db, []byte(MorningTemplate))
	if err != nil {
		fmt.Printf("generate template morning data meet error, Err: [%v]\n", err)
		return
	}

	err = wesdk.SendMessage(&wesdk.SendMessageRequest{
		ToUser:  "@all",
		MsgType: "text",
		AgentID: cast.ToInt(os.Getenv("agentID")),
		Text: struct {
			Content string `json:"content"`
		}{
			Content: string(template),
		},
	})
	if err != nil {
		fmt.Printf("send message meet error, Err: [%v]\n", err)
	}
}

func init() {
	rootCmd.AddCommand(morningCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// morningCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// morningCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
