package cobra

import (
	"fmt"
	"github.com/amalmadhu06/veil"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf(
				`
Please provide a key and value
eg:
	veil set api_key aflskdfweroi12398sfdkjh28fksjdh
`)
			return
		}

		v := veil.NewVile(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
