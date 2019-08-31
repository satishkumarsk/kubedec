package cmd

import (
	"crypto/rand"
	"fmt"
	"kubedec/cryptoImpl"
	"kubedec/utilities"
	"os"

	"github.com/spf13/cobra"
)

var secretkeypath string
var cipheryamlpath string
var rootCmd = &cobra.Command{
	Use:   "kubedec",
	Short: "Sealed secret decrypter",
	Long:  `Sealed secret decrypter`,
	Run: func(cmd *cobra.Command, args []string) {
		cinput := utilities.ParseCipherYaml(cipheryamlpath)

		kinput := utilities.ParsemasterkeyYaml(secretkeypath)

		result, _ := utilities.SecretDecoder(*kinput, *cinput)

		for k, v := range result.Data {
			dec, _ := cryptoImpl.Decrypter(rand.Reader, result.Privatekey, v, result.Label)
			output := fmt.Sprintf("%s=%s \n", k, dec)
			fmt.Print(output)

		}
	},
}

func init() {

	rootCmd.Flags().StringVarP(&secretkeypath, "masterkey", "k", "", "masterkey absolute path required")
	rootCmd.MarkFlagRequired("masterkey")

	rootCmd.Flags().StringVarP(&cipheryamlpath, "sealedyaml", "f", "", "sealedyaml absolute path required")
	rootCmd.MarkFlagRequired("sealedyaml")
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
