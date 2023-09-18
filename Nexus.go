package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Define your Nexus Mods API key and the game ID for the game you're interested in
	apiKey := "jGFubaaZc48o83Sn5xsV1lRZ7kj+9LR7mbzGWp9n6jQ78A4vRXDb4A==--jT/4jBkt3tN40aRK--5rNZ/eeoGTM1AwaRqUAsrQ==" // Replace with your actual Nexus Mods API key
	gameID := "stardewvalley"                                                                                        // Replace with the ID of the game you want to search mods for

	// Define the search parameters
	searchQuery := "expanded"       // Replace with your search query
	searchCategory := "mod"         // You can specify the category like "mod" or "addon"
	searchSorting := "popular_desc" // You can specify the sorting criteria

	// Create the API request URL with the search parameters
	apiURL := fmt.Sprintf("https://api.nexusmods.com/v1/games/%s/mods?search=%s&category=%s&sorting=%s",
		gameID, url.QueryEscape(searchQuery), searchCategory, searchSorting)

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the User-Agent header to identify your application

	// Set the API key header
	req.Header.Set("apikey", apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response body (which contains the mod search results)
	fmt.Println(string(responseBody))
}
