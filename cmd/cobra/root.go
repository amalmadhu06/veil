package cobra

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mitchellh/go-homedir"
)

const logo = `
                                               iiii    lllllll 
                                              i::::i   l:::::l 
                                               iiii    l:::::l 
                                                       l:::::l 
vvvvvvv           vvvvvvv   eeeeeeeeeeee                l::::l 
 v:::::v         v:::::v  ee::::::::::::ee   i:::::i    l::::l 
  v:::::v       v:::::v  e::::::eeeee:::::ee  i::::i    l::::l 
   v:::::v     v:::::v  e::::::e     e:::::e  i::::i    l::::l 
    v:::::v   v:::::v   e:::::::eeeee::::::e  i::::i    l::::l 
     v:::::v v:::::v    e:::::::::::::::::e   i::::i    l::::l 
      v:::::v:::::v     e::::::eeeeeeeeeee    i::::i    l::::l 
       v:::::::::v      e:::::::e             i::::i    l::::l 
        v:::::::v       e::::::::e           i::::::i  l::::::l
         v:::::v         e::::::::eeeeeeee   i::::::i  l::::::l
          v:::v           ee:::::::::::::e   i::::::i  l::::::l
           vvv              eeeeeeeeeeeeee   iiiiiiii  llllllll
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
|                                                               |
|          Veil is an API key and other secrets manager         |
|                                                               |
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
`

var RootCmd = &cobra.Command{
	Use:   "veil",
	Short: "Veil is an API key and other secrets manager",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Printf(logo) },
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the key to use when encoding and decoding secrets")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
