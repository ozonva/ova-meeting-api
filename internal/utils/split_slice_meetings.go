package utils

import "github.com/ozonva/ova-meeting-api/internal/models"

// SplitSliceMeetings split slice of models.Meetings to slice of slices
func SplitSliceMeetings(slice []models.Meeting, chunkSize uint) [][]models.Meeting {
	if chunkSize == 0 {
		panic("Chunk size must be more than 0")
	}
	sliceLen := uint(len(slice))
	if sliceLen == 0 {
		return [][]models.Meeting{}
	}
	if sliceLen < chunkSize {
		resultSlice := make([][]models.Meeting, 1)
		resultSlice[0] = slice
		return resultSlice
	}
	resSliceLen := sliceLen / chunkSize // The integer part of division
	modulo := sliceLen % chunkSize
	if modulo > 0 {
		resSliceLen++
	}
	resultSlice := make([][]models.Meeting, 0, resSliceLen)
	for start := uint(0); start < sliceLen; start += chunkSize {
		end := start + chunkSize
		if end > sliceLen {
			end = sliceLen
		}
		chunkSlice := append(make([]models.Meeting, 0, end-start), slice[start:end]...)
		resultSlice = append(resultSlice, chunkSlice)
	}
	return resultSlice
}
