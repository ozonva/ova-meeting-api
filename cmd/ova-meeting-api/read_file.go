package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ozonva/ova-meeting-api/internal/models"
)

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

func readMeetings() {
	var meetings []models.Meeting
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	storagePath := filepath.Join(pwd, "storage")
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

	meetingFiles, err := meetingListDir.ReadDir(readAllEntities)
	if err != nil {
		log.Println(err)
		return
	}
	for index := range meetingFiles {

		meetingFile := meetingFiles[index]
		meetingFileName := meetingFile.Name()
		meeting, err := readMeetingFile(filepath.Join(storagePath, meetingFileName))
		if err != nil {
			log.Fatalln(err)
		}

		meetings = append(meetings, meeting)
	}

	log.Printf("%v", meetings)
}
