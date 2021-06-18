package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Devopsengineer75/copyfastgo/internal/fssync"
	"github.com/jlaffaye/ftp"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "copyfast",
		Short: "Sync folder to target",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			origin := args[0]
			target := args[1]
			fssync.CopyRecursive(origin, target)
			//fssync.ScanRecursive(origin, target)
			c, err := ftp.Dial("localhost:21", ftp.DialWithTimeout(5*time.Second))
			if err != nil {
				log.Fatal(err)
			}

			err = c.Login("toto", "toto")
			if err != nil {
				log.Fatal(err)
			}

			// Do something with the FTP conn
			if err := c.Quit(); err != nil {
				log.Fatal(err)
			}

			items, _ := ioutil.ReadDir(".")
			for _, item := range items {
				if item.IsDir() {
					subitems, _ := ioutil.ReadDir(item.Name())
					for _, subitem := range subitems {
						if !subitem.IsDir() {
							// handle file there
							fmt.Println(item.Name() + "/" + subitem.Name())
							data := bytes.NewBufferString("Hello World")
							err = c.Stor(subitem.Name(), data)
						}
					}
				} else {
					// handle file there
					fmt.Println(item.Name())
					data := bytes.NewBufferString("Hello World")
					err = c.Stor(item.Name(), data)
				}
			}

			if err != nil {
				panic(err)
			}
			r, err := c.Retr("test-file.txt")
			if err != nil {
				panic(err)
			}
			defer r.Close()

			buf, err := ioutil.ReadAll(r)
			println(string(buf))
		},
	}
	rootCmd.Execute()
}
