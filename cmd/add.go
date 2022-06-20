/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/binary"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var taskDescription *string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds trailing string to a TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(BUCKET_NAME))

			if len(args) > 0 {
				id, _ := b.NextSequence()
				bid := make([]byte, 8)
				binary.LittleEndian.PutUint64(bid, id)
				return b.Put(bid, []byte(args[0]))
			} else {
				fmt.Println("Please provide a task you would like to add")
				return nil
			}
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
