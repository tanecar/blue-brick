package cmd

import (
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check login status",
	Run: func(cmd *cobra.Command, args []string) {
		tokenPath := filepath.Join(os.TempDir(), "auth_token")

		if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
			figure.NewColorFigure("FAILED TO LOG IN", "", "cyan", true).Print()
			color.Red("❌ Not logged in")
		} else {
			figure.NewColorFigure("MILAN TANCIC LOGGED IN", "", "cyan", true).Print()
			color.Green("✅ Logged in")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
