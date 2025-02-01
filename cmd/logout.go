package cmd

import (
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		tokenPath := filepath.Join(os.TempDir(), "auth_token")

		if err := os.Remove(tokenPath); err != nil {
			color.Red("❌ Failed to log out: %v", err)
			return
		}

		figure.NewColorFigure("LOGGED OUT - BYE BYE", "", "cyan", true).Print()
		color.Green("✅ Successfully logged out!")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
