package utils

func SplitSlice(slice []string, chunkSize uint) [][]string {
	if chunkSize == 0 {
		panic("Chunk size must be more than 0")
	}
	sliceLen := uint(len(slice))
	if sliceLen == 0 {
		return [][]string{}
	}
	if sliceLen < chunkSize {
		resultSlice := make([][]string, 1)
		resultSlice = append(resultSlice, slice)
		return resultSlice
	}
	resSliceLen := sliceLen / chunkSize // The integer part of division
	modulo := sliceLen % chunkSize
	if modulo > 0 {
		resSliceLen++
	}
	resultSlice := make([][]string, 0)
	for start := uint(0); start < sliceLen; start += chunkSize {
		end := start + chunkSize
		if end > sliceLen {
			end = sliceLen
		}
		chunkSlice := append(make([]string, 0, end-start), slice[start:end]...)
		resultSlice = append(resultSlice, chunkSlice)
	}
	return resultSlice
}
