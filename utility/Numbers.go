package utility

import (
	"log"
	"strconv"
)

func BytesToInt(byteArray []byte) int {
	/*
		lets convert to a string
	*/
	bytesAsString := string(byteArray)
	//for len(bytesAsString) < 8{
	//	bytesAsString = "0" + bytesAsString
	//}

	anInt, err := strconv.ParseInt(bytesAsString, 2, 64)

	if err != nil {
		log.Fatalln(err)
	}
	return int(anInt)
}
