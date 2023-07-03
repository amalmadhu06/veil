package cobra

import (
	"fmt"
	"github.com/amalmadhu06/veil"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a veil in your veil storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Key not provided. eg: veil get twitter_api_key \n")
			return
		}
		v := veil.NewVile(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("no value set")
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
