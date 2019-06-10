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
	"github.com/ennoo/rivet/utils/string"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	cmdYmlPath string
	daemon     bool
)

func init() {
	bowCmd.AddCommand(bowStartCmd)
	bowCmd.AddCommand(bowStopCmd)
	bowCmd.AddCommand(bowRestartCmd)
	bowStartCmd.Flags().StringVarP(&cmdYmlPath, "path", "p", "", "the cmd.yml profile path must be specified.")
	bowStartCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
}

var bowCmd = &cobra.Command{
	Use:   "bow",
	Short: "gateway developed with golang",
	Long: `gateway developed with golang. \n
execute gateway with command start or stop. \n
sometimes you need to use the server command to cooperate。`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var bowStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a gateway service",
	Long: `start a gateway service. \n
you need to prepare a complete cmd.yml configuration file`,
	Args: func(cmd *cobra.Command, args []string) error {
		if str.IsEmpty(cmdYmlPath) {
			return errors.New("the cmd.yml profile path must be specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		startBowCMD()
	},
}

var bowStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a gateway service",
	Long:  `stop a gateway service`,
	Run: func(cmd *cobra.Command, args []string) {
		stopBowCMD()
	},
}

var bowRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart a gateway service",
	Long:  `restart a gateway service`,
	Run: func(cmd *cobra.Command, args []string) {
		stopBowCMD()
		startBowCMD()
	},
}

func startBowCMD() {
	bow, err := yamlBow(cmdYmlPath)
	if nil == err {
		if daemon {
			command := exec.Command("./rivet", "bow", "start", "-p", cmdYmlPath)
			_ = command.Start()
			fmt.Printf("rivet bow start, [PID] %d running...\n", command.Process.Pid)
			_ = ioutil.WriteFile("bow.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
			daemon = false
			os.Exit(0)
		} else {
			fmt.Println("rivet bow start")
		}
		err = startBow(bow)
	}
	if nil != err {
		fmt.Println("start bow error: ", err.Error())
	}
}

func stopBowCMD() {
	data, _ := ioutil.ReadFile("bow.lock")
	command := exec.Command("kill", string(data))
	_ = command.Start()
	println("rivet bow stop")
}
