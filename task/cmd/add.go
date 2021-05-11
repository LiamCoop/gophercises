/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Used to add tasks to the TODO list",
	Long:  "Used to add tasks to the TODO list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a TODO item to add")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			panic("DB can't be opened")
		}

		err = db.View(func(tx *bolt.Tx) error {
			return db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte("todos"))
				if err != nil {
					return fmt.Errorf("create bucket: %s", err)
				}
				for _, v := range args {
					id, _ := b.NextSequence()
					tid := make([]byte, 8)
					binary.BigEndian.PutUint64(tid, uint64(id))
					err := b.Put(tid, []byte(v))
					if err != nil {
						return err
					}
				}
				return nil
			})
		})

		db.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
