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
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a TODO list item as completed",
	Long:  "Marks a TODO list item as completed",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a valid todo item")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			panic("DB can't be opened")
		}

		err = db.Update(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte("todos"))
			for _, v := range args {
				tid := make([]byte, 8)
				num, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				binary.BigEndian.PutUint64(tid, uint64(num))
				err = b.Delete(tid)
				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			panic("error on dbView")
		}
		db.Close()
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
