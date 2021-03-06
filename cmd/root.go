package cmd

import (
	"fmt"
	"os"

	"github.com/anderslundholm/advent_of_code_2021/day01"
	"github.com/anderslundholm/advent_of_code_2021/day02"
	"github.com/anderslundholm/advent_of_code_2021/day03"
	"github.com/anderslundholm/advent_of_code_2021/day04"
	"github.com/anderslundholm/advent_of_code_2021/day05"
	"github.com/anderslundholm/advent_of_code_2021/day06"
	"github.com/anderslundholm/advent_of_code_2021/day07"
	"github.com/anderslundholm/advent_of_code_2021/day08"
	"github.com/anderslundholm/advent_of_code_2021/day09"
	"github.com/anderslundholm/advent_of_code_2021/day10"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	dayFlag  string
	partFlag string
	rootCmd  = &cobra.Command{
		Use:   "advent_of_code_2021",
		Short: "Advent of Code 2021",
		Long:  `My solutions to the Avent of Code 2021 code challange.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	day01.AppendCommand(rootCmd)
	day02.AppendCommand(rootCmd)
	day03.AppendCommand(rootCmd)
	day04.AppendCommand(rootCmd)
	day05.AppendCommand(rootCmd)
	day06.AppendCommand(rootCmd)
	day07.AppendCommand(rootCmd)
	day08.AppendCommand(rootCmd)
	day09.AppendCommand(rootCmd)
	day10.AppendCommand(rootCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent_of_code_2021.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".advent_of_code_2021" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".advent_of_code_2021")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
