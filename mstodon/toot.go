package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	m "github.com/mattn/go-mastodon"
	"github.com/e10ulen/qqw/lib"
)

func init(){
	RootCmd.AddCommand(tootCmd)
}


//tootCmd = &cobra.Command{
var tootCmd = &cobra.Command{
	Use:	"toot",
	Short:	"Mastodon toot",
	Run: func(cmd *cobra.Command, args []string){
		viper.SetConfigName(".qqw")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		lib.Check(err)
		config := &m.Config{
			Server:			viper.GetString("server"),
			ClientID:		viper.GetString("clientid"),
			ClientSecret:	viper.GetString("clientsecret"),
		}
		//var email,pass string
		email := viper.GetString("email")
		pass := viper.GetString("pass")
		c := m.NewClient(config)
		c.Authenticate(context.Background(), email, pass)
		//	ここから
		var toot string
		toot = args[0]
		c.PostStatus(context.Background(), &m.Toot{
			Status: toot,
		})
	},
}
