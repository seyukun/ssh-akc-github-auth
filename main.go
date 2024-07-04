package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type User struct {
	Login string `json:"login"`
}

type PubKey struct {
	Key string `json:"key"`
}

func main() {
	var (
		token = flag.String("token", "", "token")
		org = flag.String("org", "", "org name")
		team = flag.String("team", "", "team name")
	)

	flag.Parse()

	if *token == "" {
		fmt.Fprintf(os.Stderr, "missing token\n")
		os.Exit(1)
	}

	if *org == "" {
		fmt.Fprintf(os.Stderr, "missing org name\n")
		os.Exit(1)
	}

	if *team == "" {
		fmt.Fprintf(os.Stderr, "missing team name\n")
		os.Exit(1)
	}

	req, _ := http.NewRequest("GET", "https://api.github.com/orgs/"+*org+"/teams/"+*team+"/members", nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+*token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, _ := new(http.Client).Do(req)

	org_bytes, _ := io.ReadAll(res.Body)

	var users []User
	if err := json.Unmarshal(org_bytes, &users); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	for _, user := range users {
		req, _ := http.NewRequest("GET", "https://api.github.com/users/"+user.Login+"/keys", nil)

		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", "Bearer "+*token)
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

		res, _ := new(http.Client).Do(req)

		user_bytes, _ := io.ReadAll(res.Body)

		var pubkeys []PubKey
		if err := json.Unmarshal(user_bytes, &pubkeys); err != nil {
			continue
		}
		for _, pubkey := range pubkeys {
			if strings.HasPrefix(pubkey.Key, "sk-") {
				fmt.Println(pubkey.Key)
			}
		}
	}
}
