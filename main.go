package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var FailingInput = readInput(os.Args[2])

func main() {

	/*	1. Start with n = 2 and 𝝙 (CF ) as test
		2. Test each 𝝙i and each ∇i
		3. Three possible outcomes:
		a. Some 𝝙i causes failure: Goto 1 with 𝝙 = 𝝙i and n = 2
		b. Some ∇i causes failure: Goto 1 with 𝝙 = ∇i and n = n - 1
		c. No test causes failure:
		If granularity can be redefined (n*2 <= | 𝝙 |) go to step 1 with 𝝙 = 𝝙 and n = n * 2
		Otherwise: Done, return the 1-minimum test
	*/
	input := readInput(os.Args[1])
	result := executeDeltaDebugging(input, 2)
	fmt.Println("1-minimum test success. The failing input is:")
	fmt.Println(result)
}

func isTestFailing(input string, failingInput string) bool {
	match, _ := regexp.MatchString(failingInput, input)

	return match
}

func divideInput(input string, n int) []string {
	inputLength := len(input)
	pieceLength := inputLength / n
	pieces := make([]string, n)

	for i := 0; i < n-1; i++ {
		pieces[i] = input[i*pieceLength : i*pieceLength+pieceLength]
	}

	pieces[n-1] = input[(n-1)*pieceLength : inputLength]

	return pieces
}

func mergeExceptElementAt(elements []string, skippingElementIndex int) string {
	var sb strings.Builder
	for i, v := range elements {
		// Jump over element to skip
		if i == skippingElementIndex {
			continue
		}
		sb.WriteString(v)
	}
	return sb.String()
}

func executeDeltaDebugging(delta string, n int) string {
	pieces := divideInput(delta, n)

	deltas, nablas := divideIntoDeltasNablas(pieces)
	for i, v := range deltas {
		if isTestFailing(v, FailingInput) {
			fmt.Printf("test fails on delta[%d] %s\n", i, v)
			fmt.Println("---------------------------------------------")

			return executeDeltaDebugging(v, 2)
		} else if isTestFailing(nablas[i], FailingInput) {
			fmt.Printf("test fails on nabla[%d] %s\n", i, nablas[i])
			fmt.Println("---------------------------------------------")
			return executeDeltaDebugging(nablas[i], n-1)
		}
	}
	if canRedefineGranularity(delta, n) {
		fmt.Println("test does not fail and granularity can be redefined")
		fmt.Println("---------------------------------------------")

		return executeDeltaDebugging(delta, n*2)
	} else {
		fmt.Println("test does not fail and granularity cannot be redefined")
		fmt.Println("---------------------------------------------")

		return delta
	}
	return ""
}

func canRedefineGranularity(delta string, n int) bool {
	return len(delta) >= n*2
}

func divideIntoDeltasNablas(pieces []string) (deltas []string, nablas []string) {
	// We need to assign to deltas and nablas
	n := len(pieces)
	deltas = make([]string, n)
	nablas = make([]string, n)
	for i, v := range pieces {
		deltas[i] = v
		nablas[i] = mergeExceptElementAt(pieces, i)
	}

	// debug
	for i := 0; i < n; i++ {
		fmt.Printf("delta %d \t%s\n", i, deltas[i])
		fmt.Printf("nabla %d \t%s\n", i, nablas[i])
	}

	return deltas, nablas
}

func readInput(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("error opening file: " + err.Error())
	}
	return string(content)
}
