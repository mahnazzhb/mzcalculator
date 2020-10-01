package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mahnazzhb/mzcalculator/calc"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"
)

var cfgFile string
var (
	rootCmd = &cobra.Command{
		Use:   "calcit",
		Short: "This is your simple calculator",
	}

	// rootCmd represents the base command when called without any subcommands
	AddCmd = &cobra.Command{
		Use:                "add",
		Short:              "add nums",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			nums, err := calc.ArgsTofloat64(args)
			if err != nil {
				return err
			}

			x, y := calc.Do(nums, calc.Add)
			fmt.Println(x)
			return y
		},
	}

	subCmd = &cobra.Command{
		Use:   "sub",
		Short: "Subtracts numbers",

		RunE: func(cmd *cobra.Command, args []string) error {
			nums, err := calc.ArgsTofloat64(args)
			if err != nil {
				return err
			}

			x, y := calc.Do(nums, calc.Sub)
			fmt.Println(x)
			return y
		},
	}

	mulCmd = &cobra.Command{
		Use:   "mul",
		Short: "multiply numbers",

		RunE: func(cmd *cobra.Command, args []string) error {
			nums, err := calc.ArgsTofloat64(args)
			if err != nil {
				return err
			}

			x, y := calc.Do(nums, calc.Mul)
			fmt.Println(x)
			return y
		},
	}

	divCmd = &cobra.Command{
		Use:   "div",
		Short: "divide numbers",

		RunE: func(cmd *cobra.Command, args []string) error {
			nums, err := calc.ArgsTofloat64(args)

			if err != nil {
				fmt.Println(err)
				return err

			}

			x, err := calc.Do(nums, calc.Div)
			if err != nil {
				return err
			}

			fmt.Println(x)
			return nil
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(
		AddCmd,
		subCmd,
		mulCmd,
		divCmd,
	)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobiTest2.yaml)")

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

		// Search config in home directory with name ".mzcalculator" (without extension).
		viper.AddConfigPath(home)

		viper.SetConfigName(".mzcalculator")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
