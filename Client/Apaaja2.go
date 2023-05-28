package client

import (
    "GDocs/Token"
    "context"
    "net/http"

    "golang.org/x/oauth2"
)

// Retrieves a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := token.TokenFromFile(tokFile)
	if err != nil {
		tok = token.GetTokenFromWeb(config)
		token.SaveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}
