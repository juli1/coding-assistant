package cli

import _ "embed"

//go:embed ascii/welcome.txt
var welcomeMessage string

func GetWelcomeMessage() string {
	return welcomeMessage
}
