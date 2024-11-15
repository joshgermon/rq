package cmd

import (
	"github.com/joshgermon/rq/rq"
	"github.com/spf13/cobra"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    filepath := args[0]
    apiSpec := rq.ParseOpenApiSpec(filepath)
    rq.CallForm(apiSpec)
	},
}

func init() {
	rootCmd.AddCommand(callCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  //
}
