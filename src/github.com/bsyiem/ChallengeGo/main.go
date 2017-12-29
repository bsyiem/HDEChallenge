/*
author: Brandon V. Syiem
description: used to calculate the sum of squares of positive numbers from a list of numbers.
 Used for an arbitrary number of test cases 'N'.
 */

package main

import (
	"fmt"
	"math"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	numberOfTestCases := getNumberOfTestCases(reader)
	var tempAnswer = make([]int,numberOfTestCases)
	answers := getAnswerForAllTestCases(0,numberOfTestCases,tempAnswer,reader)
	printResult(0,numberOfTestCases,answers)
}

/*
prints the result for 'numberOfTestCases' cases with the corresponding 'answers'
 */
func printResult(current int,numberOfTestCases int, answers []int){
	if current == numberOfTestCases {
		return
	}
	fmt.Println(answers[current])
	current++
	printResult(current,numberOfTestCases,answers)
}

/*
returns a list of answers for each test case
 */
func getAnswerForAllTestCases(current int, numberOfTestCases int, answer []int,reader *bufio.Reader) []int{
	if current == numberOfTestCases {
		return answer
	}
	answer[current] = getAnswerForNextTestCase(reader)
	current++
	return getAnswerForAllTestCases(current,numberOfTestCases,answer,reader)
}

/*
gets the answer for the Next unread test case
 */
func getAnswerForNextTestCase(reader *bufio.Reader) int {
	lengthOfTestCase:=getLengthOfTestCase(reader)
	temp := make([]int,lengthOfTestCase)
	nums := getNumbers(0,lengthOfTestCase,temp,reader)
	return getSumOfSquares(lengthOfTestCase,nums,0)
}

/*
gets the number of test cases
 */
func getNumberOfTestCases(reader *bufio.Reader) int {
	return getOneInt(read(reader))
}

/*
gets the number of 'Numbers' for the next test case
 */
func getLengthOfTestCase(reader *bufio.Reader) int {
	return  getOneInt(read(reader))
}

/*
get the actual numbers whose sums of squares is to be computed for the next test case
 */
func getNumbers(current int,lengthOfTestCase int, nums []int, reader *bufio.Reader) []int {
	return getNInts(read(reader),lengthOfTestCase,nums,current)
}


/*
gets the sum of square for positive numbers in the current test case
 */
func getSumOfSquares(lengthOfTestCase int, num []int, sum int) int {
	squarePos := int(0)
	if num[lengthOfTestCase-1] >= 0 {
		squarePos = int(math.Pow(float64(num[lengthOfTestCase-1]),2))
	}

	sum = sum+squarePos

	if lengthOfTestCase - 1 > 0 {
		return getSumOfSquares(lengthOfTestCase-1,num,sum)
	} else {
		return sum
	}
}

/*
read a line from stdin and returns an array of string using delim = \n
 */
func read(reader *bufio.Reader) []string {
	numString,err := reader.ReadString('\n')
	if err != nil{
		panic(err)
	}
	numString = strings.TrimSpace(numString)
	arr := strings.Split(numString," ")
	return arr

}

/*
takes in a string array and returns a single int
throws an exception if the string array has greater than one value
 */
func getOneInt(sarr []string) int {
	if len(sarr) != 1 {
		panic("Invalid Input")
	}
	return getInt(sarr,0)
}

/*
takes in a string array and an index
converts the string at the index to an int and returns that int
 */
func getInt(sarr []string,index int) int{
	num,err := strconv.Atoi(sarr[index])
	if err != nil{
		panic(err)
	}
	return num
}

/*
takes in a string array and return an array of ints
 */
func getNInts(sarr []string,lengthOfTestCase int, nums []int, current int) []int{
	if len(sarr) != lengthOfTestCase {
		panic("invalid input")
	}

	if current == lengthOfTestCase{
		return nums
	}
	nums[current] = getInt(sarr,current)
	current++
	return getNInts(sarr,lengthOfTestCase,nums,current)
}
