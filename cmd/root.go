/*
Copyright © 2022 Mohammed Tanveer tanveer941@gmail.com

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
	"ecs-pause-revive/actions"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ecs-pause-revive",
	Short: "Pause or revive an ECS service",
	Long:  `Pause or revive an ECS service by choosing a cluster and from a list of services`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		actionChoice, err := actions.ChoosePauseOrRevive()
		if err != nil {
			return err
		}
		log.Printf("Action choice: ", actionChoice)
		ecsActionInst := &actions.EcsAction{}
		ecsClient, _ := ecsActionInst.NewECSClient()
		clusterChoice, err := ecsClient.ChooseCluster()
		if err != nil {
			return err
		}
		log.Printf("Cluster choice: ", clusterChoice)
		serviceChoice, err := ecsClient.ChooseService(clusterChoice)
		if err != nil {
			return err
		}
		log.Printf("Service choice: ", serviceChoice)
		ecsClient.PerformAction(actionChoice, serviceChoice, clusterChoice)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ecs-pause-revive.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
