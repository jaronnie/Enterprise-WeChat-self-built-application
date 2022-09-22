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

var timing4pmTemplate = `下午四点啦
记得疫情打卡喔
也别忘啦喝当归茶喔
`

// timing4pmCmd represents the timing4pm command
var timing4pmCmd = &cobra.Command{
	Use:   "timing4pm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: SendTiming4pm,
}

func SendTiming4pm(cmd *cobra.Command, args []string) {
	template, err := util.ParseTemplate(struct {
	}{}, []byte(timing4pmTemplate))
	if err != nil {
		fmt.Printf("generate template timing4pm data meet error, Err: [%v]\n", err)
		return
	}

	err = wesdk.SendMessage(&wesdk.SendMessageRequest{
		ToUser:  "@all",
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
	rootCmd.AddCommand(timing4pmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timing4pmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timing4pmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
