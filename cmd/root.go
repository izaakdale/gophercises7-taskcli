/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a simple TODO list manager",
	Long: `Task is CLI tool for storing, listing and completing TODO tasks.

The three commands available are 
-- add
-- list
-- do
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer db.Close()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var db *bolt.DB

const (
	FILE_NAME   = "task.db"
	BUCKET_NAME = "UserTasks"
)

func init() {
	var err error
	db, err = bolt.Open(FILE_NAME, 0600, nil)
	if err != nil {
		log.Fatal("Failed to open bolt db")
		os.Exit(2)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return err
	})

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
