/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// runCmd represents the run command
var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run called")
			fmt.Println(testString)
			confFile, err := cmd.Flags().GetString("kubeconfig")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(confFile)

			config, err := clientcmd.BuildConfigFromFlags("", confFile)
			if err != nil {
				panic(err.Error())
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}
			pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		},
	}
	testString string
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&testString, "ptest", "p", "", "hi this is a test string")
	runCmd.MarkFlagRequired("ptest")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
