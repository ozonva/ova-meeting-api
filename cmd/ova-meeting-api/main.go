package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ozonva/ova-meeting-api/internal/models"
)

func main() {
	// fmt.Println("Hello from ova-meeting-api. 👋")
	pwd, _ := os.Getwd()
	storagePath := pwd + "/storage"
	meetingListDir, err := os.Open(storagePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(d *os.File) {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}(meetingListDir)

	meetingFiles, _ := meetingListDir.ReadDir(0)
	var meetings []models.Meeting
	for index := range meetingFiles {

		meetingFile := meetingFiles[index]
		meetingFileName := meetingFile.Name()
		meeting, err := readMeetingFile(storagePath + "/" + meetingFileName)
		if err != nil {
			log.Fatalln(err)
		}

		meetings = append(meetings, meeting)
	}

	log.Printf("%v", meetings)
}

func readMeetingFile(filename string) (models.Meeting, error) {
	var m models.Meeting
	f, err := os.Open(filename)
	if err != nil {
		return m, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(byteValue, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}
