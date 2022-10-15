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

var timing10amTemplate = `上午10点啦
记得疫情打卡喔
`

// timing10amCmd represents the timing10am command
var timing10amCmd = &cobra.Command{
	Use:   "timing10am",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: SendTiming10am,
}

func SendTiming10am(cmd *cobra.Command, args []string) {
	template, err := util.ParseTemplate(struct {
	}{}, []byte(timing10amTemplate))
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
	rootCmd.AddCommand(timing10amCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timing10amCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timing10amCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
