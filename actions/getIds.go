package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

type music_profile struct {
	danceability     float64
	energy           float64
	key              float64
	loudness         float64
	speechiness      float64
	acousticness     float64
	instrumentalness float64
	liveness         float64
	valence          float64
	tempo            float64
}

var key string = "Bearer BQALfx4PDAYXXVkadSaFlwqX9SNpMUFgFQL9gT-5Jksmgl7ymX9CQDxYofkwh-HSOdlwfQevsm-8hVMsYrfcxR97J7sJU6RLI6g_K7LfjqCJtP7jW96Ywgt6metN82saCdzOGh4GGBJz3hb4ibMyc84cSyEce09BgMA"
var userId string = "1268570385"

func init() {
	fmt.Println("initialized")
}

// Gets song IDs from users saved library
func getIds() []string {
	client := &http.Client{}
	var ids []string
	songRequest, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/tracks?limit=50&offset=0", nil)
	songRequest.Header.Add("Authorization", key)
	response, err := client.Do(songRequest)
	if err != nil {
		fmt.Println("Request failed with error:", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		items := gjson.Get(string(data), "items")
		for i := 0; i < len(items.Array()); i++ {
			track := gjson.Get(items.Array()[i].String(), "track")
			id := gjson.Get(track.String(), "id")
			ids = append(ids, id.String())
		}
	}
	ids = append(ids, getPlaylistIds()...) // Calls to get song IDs from user playlists
	return fixIds(ids)
}

// Removes empty entries in ids slice
func fixIds(ids []string) []string {
	var finalIds []string
	for _, id := range ids {
		if id != "" {
			finalIds = append(finalIds, id)
		}
	}
	return finalIds
}

// Gets IDs for users playlists and calls getPlaylistSongIds to get song IDs in playlists
func getPlaylistIds() []string {
	client := &http.Client{}

	var ids []string
	url := "https://api.spotify.com/v1/users/" + userId + "/playlists?limit=50"
	playlistRequest, err := http.NewRequest("GET", url, nil)

	playlistRequest.Header.Add("Authorization", key)
	response, err := client.Do(playlistRequest)
	if err != nil {
		fmt.Println("Request failed with error:", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		items := gjson.Get(string(data), "items")
		for i := 0; i < len(items.Array()); i++ {
			id := gjson.Get(items.Array()[i].String(), "id")
			ids = append(ids, id.String())
		}
	}
	return getPlaylistSongIds(ids)
}

// Gets song IDs from user playlists
func getPlaylistSongIds(playlistIds []string) []string {
	client := &http.Client{}

	var ids []string

	for _, plId := range playlistIds {
		url := "https://api.spotify.com/v1/playlists/" + plId + "/tracks?market=US&fields=items(track(id))&limit=50&offset=0"
		playlistRequest, err := http.NewRequest("GET", url, nil)
		playlistRequest.Header.Add("Authorization", key)
		response, err := client.Do(playlistRequest)
		if err != nil {
			fmt.Println("Request failed with error:", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			items := gjson.Get(string(data), "items")
			for i := 0; i < len(items.Array()); i++ {
				id := gjson.Get(items.Array()[i].String(), "track.id")
				ids = append(ids, id.String())
			}
		}
	}
	return ids
}
