package utility

import (
	"log"
	"strconv"
)

func StringToInt(aString string) int {
	aNumber, err := strconv.Atoi(aString)
	if err != nil {
		panic("Expected a number and got: " + aString)
	}
	return aNumber
}

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
