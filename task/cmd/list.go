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
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all current items",
	Long:  "lists all current items in the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			panic("DB can't be opened")
		}

		err = db.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte("todos"))

			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("%d: %s\n", binary.BigEndian.Uint64(k), v)
			}

			return nil
		})
		if err != nil {
			panic("error on dbVIew")
		}
		db.Close()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
