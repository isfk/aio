/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/isfk/aio/srv/api"
	"github.com/spf13/cobra"
)

var apiConfig string

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api srv",
	Long: "api srv",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api srv called")
		api.Start(apiConfig)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	apiCmd.PersistentFlags().StringVar(&apiConfig, "config", "./config.yaml", "config file, default: ./config.yaml")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
