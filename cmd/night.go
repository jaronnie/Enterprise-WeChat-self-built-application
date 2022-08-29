/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/jaronnie/ewsba/api"
	"github.com/jaronnie/ewsba/util"
	"github.com/jaronnie/ewsba/wesdk"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var NightTemplate = `😘😘😘😘😘😘😘😘😘😘😘😘😘 睡觉觉啦 😘😘😘😘😘😘😘😘😘😘😘😘😘
缘儿的猪发消息来啦

当前时间：{{.CurrentlyTime}}

每日一句：{{.SentenceEnglish}}
{{.SentenceChinese}}
`

type NightData struct {
	CurrentlyTime   string
	SentenceEnglish string
	SentenceChinese string
}

// nightCmd represents the night command
var nightCmd = &cobra.Command{
	Use:   "night",
	Short: "send night message to your girl friend",
	Long:  `send night message to your girl friend`,
	Run:   SendNight,
}

func SendNight(cmd *cobra.Command, args []string) {
	sentenceEnglish, sentenceChinese, err := api.GetSentence()
	if err != nil {
		fmt.Printf("get sentence meet error, Err: [%v]\n", err)
		return
	}
	db := NightData{
		CurrentlyTime:   cast.ToString(time.Now().Format("2006-01-02")),
		SentenceEnglish: sentenceEnglish,
		SentenceChinese: sentenceChinese,
	}

	template, err := util.ParseTemplate(db, []byte(NightTemplate))
	if err != nil {
		fmt.Printf("generate template night data meet error, Err: [%v]\n", err)
		return
	}

	err = wesdk.SendMessage(&wesdk.SendMessageRequest{
		ToUser:  "NieJian",
		MsgType: "text",
		AgentID: cast.ToInt(os.Getenv("agentID")),
		Text: wesdk.Content{
			Content: string(template),
		},
	})
	if err != nil {
		fmt.Printf("send message meet error, Err: [%v]\n", err)
	}
}

func init() {
	rootCmd.AddCommand(nightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nightCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
