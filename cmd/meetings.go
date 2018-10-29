// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// create meeting command
var cmCmd = &cobra.Command{
	Use:   "cm -t [title] -p [p1 p2 ....] -s [sTime] -e [eTime]",
	Short: "Create Meeting command",
	Long:  "Create Meeting with title, participators, start time, end time",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participators")
		sTime, _ := cmd.Flags().GetString("sTime")
		eTime, _ := cmd.Flags().GetString("eTime")
		fmt.Println("cm called with: " + title + " " + participators + " " + sTime + " " + eTime)
	},
}

// add participators command
var apCmd = &cobra.Command{
	Use:   "ap -p [p1 p2 ...]",
	Short: "Add Meeting Participators Command",
	Long:  "Add meeting participators, must be Agenda users",
	Run: func(cmd *cobra.Command, args []string) {
		participators, _ := cmd.Flags().GetString("participators")
		fmt.Println("ap called with: " + participators)
	},
}

// delete participators command
var dpCmd = &cobra.Command{
	Use:   "dp -p [p1 p2 ...]",
	Short: "Delete Meeting Participators Command",
	Long:  "Delete meeting participators, must be Agenda users",
	Run: func(cmd *cobra.Command, args []string) {
		participators, _ := cmd.Flags().GetString("participators")
		fmt.Println("dp called with: " + participators)
	},
}

// query meetings command
var qmCmd = &cobra.Command{
	Use:   "qm -s [sTime] -e [eTime]",
	Short: "Query Meetings Command",
	Long:  "Query meetings between start time and end time",
	Run: func(cmd *cobra.Command, args []string) {
		sTime, _ := cmd.Flags().GetString("sTime")
		eTime, _ := cmd.Flags().GetString("eTime")
		fmt.Println("qm called with: " + sTime + ", " + eTime)
	},
}

// delete meeting command
var dmCmd = &cobra.Command{
	Use:   "dm -t [title]",
	Short: "Delete Meeting Command",
	Long:  "Delete meeting with title, and you must be the meeting's sponsor",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		fmt.Println("dm called with: " + title)
	},
}

// quit meeting command
var quitmCmd = &cobra.Command{
	Use:   "quitm -t [title]",
	Short: "Quit Meeting Command",
	Long:  "Quit meeting with title, and you must be the meeting's participator",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		fmt.Println("quitm called with: " + title)
	},
}

// clear meetings command
var clallCmd = &cobra.Command{
	Use:   "clall",
	Short: "Clear All Meetings Command",
	Long:  "Clear all meetings you sponsor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clall called")
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "NULL", "New meeting's title")
	cmCmd.Flags().StringP("participators", "p", "NULL", "New meeting's participators")
	cmCmd.Flags().StringP("sTime", "s", "NULL", "New meeting's start time")
	cmCmd.Flags().StringP("eTime", "e", "NULL", "New meeting's end time")

	rootCmd.AddCommand(apCmd)
	apCmd.Flags().StringP("participators", "p", "NULL", "Add meeting's participators")

	rootCmd.AddCommand(dpCmd)
	dpCmd.Flags().StringP("participators", "p", "NULL", "Delete meeting's participators")

	rootCmd.AddCommand(qmCmd)
	qmCmd.Flags().StringP("sTime", "s", "NULL", "Start time of meeting")
	qmCmd.Flags().StringP("eTime", "e", "NULL", "End time of meeting")

	rootCmd.AddCommand(dmCmd)
	dmCmd.Flags().StringP("title", "t", "NULL", "Title of meeting you want to delete")

	rootCmd.AddCommand(quitmCmd)
	quitmCmd.Flags().StringP("title", "t", "NULL", "Title of meeting you want to quit")

	rootCmd.AddCommand(clallCmd)
}
