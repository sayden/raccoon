package raccoon_cli

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/thehivecorporation/raccoon/server"
	"github.com/thehivecorporation/raccoon/constants"
)

func Init() {
	app := cli.NewApp()
	app.Name = constants.APP_NAME
	app.Usage = constants.APP_DESCRIPTION
	app.Version = constants.VERSION

	app.Commands = []cli.Command{
		{
			Name:   constants.INSTRUCTIONS_NAME,
			Usage:  constants.INSTRUCTIONS_USAGE,
			Action: executeZombieBook,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  constants.INSTRUCTIONS_FLAG_ALIAS,
					Usage: constants.INSTRUCTIONS_USAGE,
				},
				cli.StringFlag{
					Name: constants.HOSTS_FLAG_ALIAS,
					Usage: constants.HOSTS_FLAG_USAGE,
				},
			},
		},
		{
			Name:   constants.SERVER_NAME,
			Usage:  constants.SERVER_USAGE,
			Action: server.REST,
		},
	}

	app.Run(os.Args)
}