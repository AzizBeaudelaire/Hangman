package hangman

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// StopAndSaveGame : Stop the game and sore the game data in a TXT file
func StopAndSaveGame(data GameData) {
	DataGameJson, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile("Saves/save.txt", DataGameJson, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// StartWithFlag : Restart the game with the save saved in a txt file
func StartWithFlag(Start string) GameData {
	var data GameData
	JsonData, _ := ioutil.ReadFile(Start)
	err := json.Unmarshal(JsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
