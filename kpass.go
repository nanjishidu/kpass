package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/nanjishidu/gomini"
	"github.com/urfave/cli"
)

//go:generate go-bindata -o=asset/asset.go -pkg=asset static/... template/...
var (
	loger        *log.Logger
	kpassDir     string
	kpassSign    = "^%!_$#@*^%!_$#@*^%!_$#@*"
	emap         = make(map[string]string)
	kpassDirFlag = cli.StringFlag{
		Name:  "kpassdir,kd",
		Value: "./data",
		Usage: "kpass dir",
	}
	kpassFileFlag = cli.StringFlag{
		Name:  "kpassfile,kf",
		Value: "./kpass-encrypt",
		Usage: "kpass encrypt file",
	}
	kpassWordFlag = cli.StringFlag{
		Name:  "kpassword,kp",
		Value: "",
		Usage: "kpass password",
	}
	kpassCryptoFlag = cli.StringFlag{
		Name:  "kcrypto,kc",
		Value: "aes-cfb",
		Usage: "kpass crypto",
	}
	kpassHostFlag = cli.StringFlag{
		Name:  "host",
		Value: "127.0.0.1",
		Usage: "kpass web host",
	}
	kpassPortFlag = cli.IntFlag{
		Name:  "port",
		Value: 9988,
		Usage: "kpass web port",
	}
)

func main() {
	loger = log.New(os.Stdout, "", log.LstdFlags)
	app := cli.NewApp()
	app.Name = "kpass"
	app.Usage = "password management tool for golang"
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create kpassfile by kpassword",
			Flags: []cli.Flag{
				kpassDirFlag,
				kpassFileFlag,
				kpassWordFlag,
				kpassCryptoFlag,
			},
			Action: func(c *cli.Context) error {
				if c.String("kpassword") != "" {
					if !gomini.IsExist(c.String("kpassdir")) {
						gomini.Mkdir(c.String("kpassdir"))
					}
					kpassdfile := filepath.Join(c.String("kpassdir"), c.String("kpassfile"))
					if gomini.IsExist(kpassdfile) {
						loger.Printf("%s is exist", c.String("kpassfile"))
					} else {
						kcrypto := c.String("kcrypto")
						hook, err := NewInstance(kcrypto)
						if err != nil {
							loger.Printf("%s is not exist", kcrypto)
						} else {
							bse, err := hook().Encrypt(c.String("kpassword"), "+ welcome to use kpass to protect your password")
							if err != nil {
								loger.Println(err.Error())
								loger.Printf("create %s failed", kpassdfile)
							} else {
								_, err = gomini.FilePutContent(kpassdfile, kpassSign+bse)
								if err != nil {
									loger.Printf("create %s failed", kpassdfile)
								} else {
									loger.Printf("create %s succeed", kpassdfile)
								}
							}

						}
					}
				} else {
					loger.Println("init need kpassword")
				}
				return nil
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "kpass run",
			Flags: []cli.Flag{
				kpassDirFlag,
				kpassHostFlag,
				kpassPortFlag,
			},
			Action: func(c *cli.Context) error {
				kpassDir = c.String("kpassdir")
				if !gomini.IsExist(kpassDir) {
					gomini.Mkdir(kpassDir)
				}
				runWebApp(fmt.Sprintf("%s:%d", c.String("host"), c.Int("port")))
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	app.Version = "1.0"
	app.Compiled = time.Now()
	app.Run(os.Args)
}
