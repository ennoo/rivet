/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package command

import (
	"errors"
	"fmt"
	"github.com/ennoo/rivet"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(bowCmd)
	rootCmd.AddCommand(shuntCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of rivet",
	Long:  `print the version number of rivet`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rivet Version:v" + rivet.Version)
	},
}

var rootCmd = &cobra.Command{
	Use:   "rivet",
	Short: "rivet is a cli library for bow and shunt",
	Long: `rivet is a cli library for bow and shunt. \n
use rivet can operation bow and shunt, like start or stop. \n
also can operate server if server used`,
	Args: func(cmd *cobra.Command, args []string) error {

		// Do Stuff Here
		if len(args) < 1 {
			return errors.New("command is required , Use rivet -h to get more information ")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute cmd start
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		//zap.S().Debug(err)
		os.Exit(1)
	}
}
