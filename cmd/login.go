package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate using a web browser",
	Run: func(cmd *cobra.Command, args []string) {
		// Show ASCII Art
		figure.NewColorFigure("MILAN TANCIC AUTH CLI", "", "cyan", true).Print()

		// Open browser for authentication
		authURL := "https://your-auth0-domain/authorize?client_id=your-client-id&response_type=code&redirect_uri=http://localhost:3000/callback"
		color.Blue("\n🔑 Opening browser for authentication...")

		if err := openBrowser(authURL); err != nil {
			color.Red("❌ Failed to open browser: %v", err)
			return
		}

		// Simulate fetching a JWT Token
		token, err := generateJWT("milan-tancic") // Replace with real token fetching
		if err != nil {
			color.Red("❌ Failed to generate token: %v", err)
			return
		}

		// Store the token
		tokenPath := filepath.Join(os.TempDir(), "auth_token")
		os.WriteFile(tokenPath, []byte(token), 0644)
		color.Green("\n✅ Login successful! Token saved.")
	},
}

// Open the browser for authentication
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}

// Generate a fake JWT (Replace with actual JWT retrieval)
func generateJWT(userID string) (string, error) {
	secret := []byte("supersecret") // Replace with actual secret or private key

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	return token.SignedString(secret)
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
