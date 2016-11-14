package main

func getCommonLines(partOne, partTwo Lines) (ret Lines) {
	for _, lineWanted := range partOne {
		for _, lineReached := range partTwo {
			if lineWanted.ID == lineReached.ID {
				ret = append(ret, lineWanted)
			}
		}
	}
	return ret
}

//GetAllPath is used to find paths between idStart and isStop busstops
func GetAllPath(idStart, idStop string) (ret []byte, err error) {
	linesAtStart, err := GetLinesByStopID(idStart)
	if err != nil {
		return nil, err
	}
	linesAtEnd, err := GetLinesByStopID(idStop)
	if err != nil {
		return nil, err
	}
	commonsLines := getCommonLines(linesAtStart, linesAtEnd)
	if commonsLines == nil {

	}
	return nil, nil
}
