/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "task do IN, removes a task from the list of IN",
	Run: func(cmd *cobra.Command, args []string) {

		db.Update(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte(BUCKET_NAME))
			if b == nil {
				return errors.New("No bucket by that name")
			}

			if len(args) > 0 {
				in, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}
				c := b.Cursor()
				k, _ := c.First()
				for i := 1; i < in; i++ {
					k, _ = c.Next()
				}
				err = b.Delete(k)
			} else {
				fmt.Println("Please provide an item number")
			}

			var err error
			return err
		})

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
