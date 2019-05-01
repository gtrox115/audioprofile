package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// Helper function to create profile
func createProfile(ids []string) music_profile {
	features := getFeatures(ids)
	return getAverages(features)
}

// Gets song features for all IDs in slice 'initialIds'
func getFeatures(initialIds []string) []gjson.Result {
	var results []gjson.Result
	client := &http.Client{}

	for i := 0; i < (len(initialIds)/100)+1; i++ {
		var ids []string
		if i == len(initialIds)/100 {
			ids = initialIds[(i * 100):]
		} else {
			ids = initialIds[(i * 100):((i + 1) * 100)]
		}

		formattedIds := ids[0]
		ids = ids[1:]
		for _, id := range ids {
			if !strings.Contains(id, " ") && id != "" {
				formattedIds += "%2C"
				formattedIds += id
			}
		}
		apiReq := "https://api.spotify.com/v1/audio-features?ids=" + formattedIds
		featureRequest, err := http.NewRequest("GET", apiReq, nil)
		featureRequest.Header.Add("Authorization", key)

		response, err := client.Do(featureRequest)
		if err != nil {
			fmt.Println("Request failed with error:", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			items := gjson.Get(string(data), "audio_features")
			results = items.Array()
		}
	}
	return results
}

// Averages features from all songs into a single profile
func getAverages(features []gjson.Result) music_profile {
	var profile music_profile
	var energy []float64
	var danceability []float64
	var key []float64
	var loudness []float64
	var speechiness []float64
	var acousticness []float64
	var instrumentalness []float64
	var liveness []float64
	var valence []float64
	var tempo []float64

	for _, value := range features {
		energy = append(energy, gjson.Get(value.String(), "energy").Float())
		danceability = append(danceability, gjson.Get(value.String(), "danceability").Float())
		key = append(key, gjson.Get(value.String(), "key").Float())
		loudness = append(loudness, gjson.Get(value.String(), "loudness").Float())
		speechiness = append(speechiness, gjson.Get(value.String(), "speechiness").Float())
		acousticness = append(acousticness, gjson.Get(value.String(), "acousticness").Float())
		instrumentalness = append(instrumentalness, gjson.Get(value.String(), "instrumentalness").Float())
		liveness = append(liveness, gjson.Get(value.String(), "liveness").Float())
		valence = append(valence, gjson.Get(value.String(), "valence").Float())
		tempo = append(tempo, gjson.Get(value.String(), "tempo").Float())
	}

	profile.energy = avg(energy)
	profile.danceability = avg(danceability)
	profile.key = avg(key)
	profile.loudness = avg(loudness)
	profile.speechiness = avg(speechiness)
	profile.acousticness = avg(acousticness)
	profile.instrumentalness = avg(instrumentalness)
	profile.liveness = avg(liveness)
	profile.valence = avg(valence)
	profile.tempo = avg(tempo)
	return profile
}

// Average function
func avg(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	result := sum / float64(len(values))
	return result
}

func mapProfile(profile music_profile) map[string]float64 {
	value := map[string]float64{"Energy:": profile.energy, "danceability:": profile.danceability, "key:": profile.key, "loudness:": profile.loudness, "speechiness:": profile.speechiness, "acousticness:": profile.acousticness, "instrumentalness:": profile.instrumentalness, "liveness:": profile.liveness}
	return value
}
