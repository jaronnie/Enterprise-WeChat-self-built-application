/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	"github.com/jaronnie/ewsba/util"
	"github.com/jaronnie/ewsba/wesdk"
	"github.com/spf13/cast"
	"os"

	"github.com/spf13/cobra"
)

var timing2dot30pmTemplate = `下午两点半啦
记得签到喔
`

// timing2dot30pmCmd represents the timing2dot30pm command
var timing2dot30pmCmd = &cobra.Command{
	Use:   "timing2dot30pm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: SendTiming2dot30pm,
}

func SendTiming2dot30pm(cmd *cobra.Command, args []string) {
	template, err := util.ParseTemplate(struct {
	}{}, []byte(timing2dot30pmTemplate))
	if err != nil {
		fmt.Printf("generate template timing2dot30pm data meet error, Err: [%v]\n", err)
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
	rootCmd.AddCommand(timing2dot30pmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timing2dot30pmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timing2dot30pmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
