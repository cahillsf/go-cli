/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

//notes for imports:
//go get k8s.io/client-go@latest
// go get k8s.io/client-go/plugin/pkg/client/auth/exec@v0.25.2
// go get k8s.io/client-go/discovery@v0.25.2
// go get k8s.io/client-go/tools/clientcmd@v0.25.2

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "go-cli-fun",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },

	}
	kubeconfig string
)

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-fun.yaml)")
	// var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file (default is $HOME<OS specific Separator>.kube<OS specific Separator>config)")
	} else {
		rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file (default is $HOME<OS specific Separator>.kube<OS specific Separator>config)")
		rootCmd.MarkFlagRequired("kubeconfig")
	}

	fmt.Println(kubeconfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}
