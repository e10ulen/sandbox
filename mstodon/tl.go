package cmd

import (
	"fmt"
	"context"
	"log"
	"strings"

	m "github.com/mattn/go-mastodon"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/e10ulen/qqw/lib"
)

func init(){
	RootCmd.AddCommand(timelineCmd)
}

var timelineCmd = &cobra.Command{
	Use:	"tl",
	Short:	"Mastodon GetTimeline.",
	Run: func(cmd *cobra.Command, args []string){
		viper.SetConfigName(".qqw")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		lib.Check(err)
		config := &m.Config{
			Server: viper.GetString("server"),
			ClientID: viper.GetString("clientid"),
			ClientSecret: viper.GetString("clientsecret"),
		}
		//var email string
		//var pass string
		email := viper.GetString("email")
		pass := viper.GetString("pass")
		c := m.NewClient(config)
		c.Authenticate(context.Background(), email, pass)
		if err != nil {
			log.Println("Client:", err)
		}
		//	GetTimeline
		wsc := c.NewWSClient()
		q, err := wsc.StreamingWSPublic(context.Background(), true)
		lib.Check(err)
		for e := range q {
			if t, ok := e.(*m.UpdateEvent); ok {
				s := t.Status.Content
				s = strings.Replace(s, "<p>", "", -1)
				s = strings.Replace(s, "</p>", "", -1)
				blue := color.New(color.FgBlue).SprintFunc()
				magenta := color.New(color.FgMagenta).SprintFunc()
				red := color.New(color.FgRed).SprintFunc()
				//
				fmt.Printf("%s %s(%s) %s\n", blue(t.Status.CreatedAt.Local().Format("15:04:05")), magenta(t.Status.Account.DisplayName), magenta(t.Status.Account.Acct), red(s))
				//

			}
		}
	},

}
