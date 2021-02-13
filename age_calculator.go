package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	//Input name
	fmt.Println("Enter your name : ")
	inputname, _ := input.ReadString('\n')
	//Input date of birth day
	fmt.Println("Enter your Date of Birth  'Day' (1-31) :")
	inputbirthday, _ := input.ReadString('\n')
	inputbirthdayconv, _ := strconv.ParseInt(strings.TrimSpace(inputbirthday), 10, 64)
	//input date of birth mounth
	fmt.Println("Enter your Date of Birth  'Month' (1-12) :")
	inputbirthmonth, _ := input.ReadString('\n')
	inputbirthmonthconv, _ := strconv.ParseInt(strings.TrimSpace(inputbirthmonth), 10, 64)
	//input date of birth year
	fmt.Println("Enter your Date of Birth  'Year' (1999):")
	inputbirthyear, _ := input.ReadString('\n')
	inputbirthyearconv, _ := strconv.ParseInt(strings.TrimSpace(inputbirthyear), 10, 64)

	//input part
	var inputday int64 = inputbirthdayconv
	var inputmonth int64 = inputbirthmonthconv
	var inputyear int64 = inputbirthyearconv
	//Present part
	nowYearMonth := time.Now().Month().String()

	nowYearMonthint := 0
	switch nowYearMonth {
	case "January":
		nowYearMonthint = 1

		//fmt.Printf("%d \n", 1)
	case "February":
		nowYearMonthint = 2
		//fmt.Printf("%d \n", 2)
	case "March":
		nowYearMonthint = 3
		//fmt.Println("3")
	case "April":
		nowYearMonthint = 4
		//fmt.Println("4")
	case "May":
		nowYearMonthint = 5
		//fmt.Println("5")
	case "June":
		nowYearMonthint = 6
		//fmt.Println("6")
	case "July":
		nowYearMonthint = 7

		//fmt.Println("7")
	case "August":
		nowYearMonthint = 8
		//fmt.Println("8")
	case "September":
		nowYearMonthint = 9
		//fmt.Println("9")
	case "October":
		nowYearMonthint = 10
		//fmt.Println("10")
	case "November":
		nowYearMonthint = 11
		//fmt.Println("11")
	case "December":
		nowYearMonthint = 12
		//fmt.Println("12")
	default:
		fmt.Println("nothing")
	}

	//Now time part
	nowday := int64(time.Now().Day())
	nowmonthfinal := int64(nowYearMonthint)
	nowyear := int64(time.Now().Year())
	//fmt.Println(nowday)
	//fmt.Println(nowmonthfinal)
	//fmt.Println(nowyear)

	//logical part
	var myday, mymonth, myyear int64 = 0, 0, 0
	if nowday >= inputday && nowmonthfinal >= inputmonth && nowyear >= inputyear {
		myyear = nowyear - inputyear
		mymonth = nowmonthfinal - inputmonth
		myday = nowday - inputday

	} else if nowday < inputday && nowmonthfinal >= inputmonth && nowyear >= inputyear {
		myyear = nowyear - inputyear
		mymonth = nowmonthfinal - inputmonth
		mymonth = mymonth - 1
		myday = (nowday + 30) - inputday

	} else if nowday >= inputday && nowmonthfinal < inputmonth && nowyear >= inputyear {
		myyear = nowyear - inputyear
		myyear = myyear - 1
		mymonth = (nowmonthfinal + 12) - inputmonth

		myday = nowday - inputday

	} else if nowday < inputday && nowmonthfinal < inputmonth && nowyear >= inputyear {
		myyear = nowyear - inputyear
		myyear = myyear - 1
		mymonth = (nowmonthfinal + 12) - inputmonth
		mymonth = mymonth - 1
		myday = (nowday + 30) - inputday

	}

	//output part
	fmt.Printf("Your name is %v Now you are %v years old \n", inputname, myyear)
	fmt.Printf("%v years %v month %v days \n", myyear, mymonth, myday)
	//fmt.Println(inputname)
	//fmt.Println(myday)
	//fmt.Println(mymonth)
	//fmt.Println(myyear)

}