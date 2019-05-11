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
 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func launch() {
	fmt.Println("nuclear launch detected")
}

func commencingCountDown(canLunch chan int) {
	c := time.Tick(1 * time.Second)
	for countDown := 10; countDown > 0; countDown-- {
		//if countDown == 15 {
		//	canLunch <- 15
		//	return
		//}
		//if nil == canLunch {
		//	return
		//}
		canLunch <- countDown
		fmt.Println("0 = ", countDown)
		<-c
	}
	canLunch <- -1
}

func commencingCountDown1(canLunch1 chan int) {
	c := time.Tick(1 * time.Second)
	for countDown := 10; countDown > 0; countDown-- {
		//if countDown == 10 {
		//	canLunch1 <- 10
		//	return
		//}
		//if nil == canLunch1 {
		//	return
		//}
		canLunch1 <- countDown
		fmt.Println("1 = ", countDown)
		<-c
	}
	canLunch1 <- -1
}

func isAbort(abort chan int) {
	//_, _ = os.Stdin.Read(make([]byte, 1))
	abortOther = abort
	//abort <- -1
}

func isAbortOther() {
	_, _ = os.Stdin.Read(make([]byte, 1))
	abortOther <- -1
}

var abortOther chan int

func main() {
	fmt.Println("Commencing coutdown")

	abort := make(chan int)
	canLunch := make(chan int)
	canLunch1 := make(chan int)
	go isAbort(abort)
	go isAbortOther()
	go commencingCountDown(canLunch)
	go commencingCountDown1(canLunch1)
	for {
		select {
		case x := <-canLunch:
			if x == 5 {
				fmt.Println("canLunch is nil")
				canLunch = nil
			}
		case y := <-canLunch1:
			//fmt.Println("canLunch1 = ", y)
			if y == 8 {
				canLunch = nil
				//canLunch = make(chan int)
				//go commencingCountDown(canLunch)
			} else if y == 2 {
				fmt.Println("canLunch1 is nil")
				canLunch1 = nil
			}
		case r := <-abort:
			fmt.Println("Launch aborted! and r = " + strconv.Itoa(r))
			return
		}
		if canLunch == nil && canLunch1 == nil {
			return
		}
		//launch()
	}
}
