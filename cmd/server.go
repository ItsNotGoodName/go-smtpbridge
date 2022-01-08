/*
Copyright © 2021 ItsNotGoodName

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/ItsNotGoodName/smtpbridge/app"
	"github.com/ItsNotGoodName/smtpbridge/config"
	"github.com/ItsNotGoodName/smtpbridge/core"
	"github.com/ItsNotGoodName/smtpbridge/left/router"
	"github.com/ItsNotGoodName/smtpbridge/left/smtp"
	"github.com/ItsNotGoodName/smtpbridge/left/web"
	"github.com/ItsNotGoodName/smtpbridge/right/endpoint"
	"github.com/ItsNotGoodName/smtpbridge/right/repository"
	"github.com/ItsNotGoodName/smtpbridge/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverConfig *config.Config

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start SMTP server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Read config
		serverConfig.Load()

		var (
			background     []core.BackgroundPort
			db             core.DatabasePort
			attachmentREPO core.AttachmentRepositoryPort
			messageREPO    core.MessageRepositoryPort
		)

		// Init database and repositories
		if serverConfig.DB.IsBolt() {
			d := repository.NewDatabase(serverConfig)
			att := repository.NewAttachment(serverConfig, &d)
			msg := repository.NewMessage(&d, att)

			db, attachmentREPO, messageREPO = d, att, msg
		} else {
			db = repository.NewMock()
			attachmentREPO = repository.NewAttachmentMock()
			messageREPO = repository.NewMessageMock()
		}

		// Init endpoint repository
		endpointREPO := endpoint.NewRepository(serverConfig)

		// Init services
		authSVC := service.NewAuth(serverConfig)
		bridgeSVC := service.NewBridge(serverConfig, endpointREPO)
		messageSVC := service.NewMessage(serverConfig, attachmentREPO, messageREPO)
		endpointSVC := service.NewEndpoint(endpointREPO, messageSVC)
		janitorSVC := service.NewJanitor(serverConfig, attachmentREPO, messageREPO)

		// Add janitor to background
		background = append(background, janitorSVC)

		// Init app
		app := app.New(
			background,
			db,
			attachmentREPO,
			messageREPO,
			endpointREPO,
			authSVC,
			bridgeSVC,
			endpointSVC,
			messageSVC,
		)

		// Init and start http server
		if serverConfig.HTTP.Enable {
			web := web.New()
			httpServer := router.New(app, web)
			go router.Start(serverConfig, httpServer)
		}

		// Init and start smtp server
		smtpBackend := smtp.NewBackend(app)
		smtpServer := smtp.New(serverConfig, smtpBackend)
		go smtpServer.Start()

		// Start app
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverConfig = config.New()

	serverCmd.Flags().String("smtp-host", serverConfig.SMTP.Host, "smtp host address to listen on")
	viper.BindPFlag("smtp.host", serverCmd.Flags().Lookup("smtp-host"))

	serverCmd.Flags().Uint16("smtp-port", serverConfig.SMTP.Port, "smtp port to listen on")
	viper.BindPFlag("smtp.port", serverCmd.Flags().Lookup("smtp-port"))

	serverCmd.Flags().Int("smtp-size", serverConfig.SMTP.Size, "max size of email in bytes")
	viper.BindPFlag("smtp.size", serverCmd.Flags().Lookup("smtp-size"))

	viper.SetDefault("database.db", serverConfig.DB.DB)
	viper.SetDefault("database.attachments", serverConfig.DB.Attachments)

	serverCmd.Flags().Bool("http", serverConfig.HTTP.Enable, "enable http server")
	viper.BindPFlag("http.enable", serverCmd.Flags().Lookup("http"))

	serverCmd.Flags().String("http-host", serverConfig.HTTP.Host, "http host address to listen on")
	viper.BindPFlag("http.host", serverCmd.Flags().Lookup("http-host"))

	serverCmd.Flags().Uint16("http-port", serverConfig.HTTP.Port, "http port to listen on")
	viper.BindPFlag("http.port", serverCmd.Flags().Lookup("http-port"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
