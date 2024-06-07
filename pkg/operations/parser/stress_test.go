package parser

import (
	"fmt"
	"goast/pkg/utils"
	"strconv"
	"testing"
	"time"
)

func TestStress1(t *testing.T) {
	numIterations := 20
	times := make([]int64, numIterations)
	stringSize := 500
	inputString := ""
	for i := 0; i < stringSize; i++ {
		inputString += strconv.Itoa(i)
		if i < stringSize-1 {
			inputString += " + "
		}
	}
	fmt.Println(inputString)
	for i := 0; i < numIterations; i++ {
		start := time.Now().UnixMilli()
		ParseOperation(inputString)
		end := time.Now().UnixMilli()
		times[i] = end - start
	}
	stringRep := utils.SliceToString(times)
	fmt.Println(stringRep)
}
