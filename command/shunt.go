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
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
)

func init() {
	shuntCmd.AddCommand(shuntStartCmd)
	shuntCmd.AddCommand(shuntStopCmd)
	shuntCmd.AddCommand(shuntRestartCmd)
}

var shuntCmd = &cobra.Command{
	Use:   "shunt",
	Short: "Load-Balance developed with golang",
	Long: `Load-Balance developed with golang. \n
execute load balance with command start or stop. \n
sometimes you need to use the server command to cooperate。`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var shuntStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a load balance service",
	Long: `start a load balance service. \n
you need to prepare a complete cmd.yml configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		startShuntCMD()
	},
}

var shuntStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a load balance service",
	Long:  `stop a load balance service`,
	Run: func(cmd *cobra.Command, args []string) {
		stopShuntCMD()
	},
}

var shuntRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart a load balance service",
	Long:  `restart a load balance service`,
	Run: func(cmd *cobra.Command, args []string) {
		stopShuntCMD()
		startShuntCMD()
	},
}

func startShuntCMD() {
	shunt, err := yamlShunt(cmdYmlPath)
	if nil == err {
		if daemon {
			command := exec.Command("./rivet", "shunt", "start", "-p", cmdYmlPath)
			_ = command.Start()
			fmt.Printf("rivet shunt start, [PID] %d running...\n", command.Process.Pid)
			_ = ioutil.WriteFile("shunt.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
			daemon = false
			os.Exit(0)
		} else {
			fmt.Println("rivet shunt start")
		}
		err = startShunt(shunt)
	}
	if nil != err {
		fmt.Println("start shunt error: ", err.Error())
	}
}

func stopShuntCMD() {
	data, _ := ioutil.ReadFile("shunt.lock")
	command := exec.Command("kill", string(data))
	_ = command.Start()
	println("rivet shunt stop")
}
