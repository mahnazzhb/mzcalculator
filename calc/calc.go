package calc

import (
	"errors"
	"fmt"
	"strconv"
)

//AgrsToInts function get the arguman the user inputs and convert them to int.

// func ArgsToInts(args []string)([]int,error){
// 	numsInt:=make([]int,0,len(args))
// 	for _,str:=range args{
// 		n,err:=strconv.Atoi(str)

// 		if err!=nil{
// 			return nil,
// 			fmt.Errorf("%q wrong format", str)
// 		}
// 		numsInt=append(numsInt,n )
// 	}
// 	return numsInt,nil

// }

// ArgsTofloat64 function get the arguman the user inputs and convert them to float64.
func ArgsTofloat64(args []string) ([]float64, error) {

	nums := make([]float64, 0, len(args))
	for _, str := range args {
		n, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil,
				fmt.Errorf("%q wrong format", str)
		}
		nums = append(nums, n)
	}
	return nums, nil

}

// Do function gets the numbers the user input and also get the type of function then send the result to the root.
// it implement to make the program simple and merge all 4function together.
func Do(nums []float64, fn func(a, b float64) (float64, error)) (float64, error) {
	ret := nums[0]
	var err error
	for _, n := range nums[1:] {
		ret, err = fn(ret, n)
		if err != nil {
			return 0, err
		}
	}
	return ret, nil
}

// these 4function do the main operand on 2by2 numbers.
func Add(a, b float64) (float64, error) {
	return a + b, nil
}
func Sub(a, b float64) (float64, error) {
	return a - b, nil
}
func Mul(a, b float64) (float64, error) {
	return a * b, nil
}
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("zero is not accepted")
	}
	return a / b, nil
}

//functions to do the main operand before implementing Do function

// func Add(nums []float64)float64{
// 	sum:=nums[0]
// 	for _,n:=range nums[1:]{
// 		sum+=n
// 	}
// 	return sum
// }
// func Sub(nums []float64)float64 {
// 	res:=nums[0]

// 	for _,n:=range nums[1:] {
// 		res-=n

// 	}
// 	return res
// }

// func Mul(nums []float64)float64{
// 	mulRes:=nums[0]
// 	for _,n:=range nums[1:]{
// 		mulRes*=n
// 	}
// 	return mulRes
// }

// func Div(nums []float64)float64{

// 	divRes:=nums[0]

// 	for _,n:=range nums[1:]{
// 		if (n==0){
// 			fmt.Println("zero is not accepted for division")
// 			return 0
// 		}

//          divRes /=n
// 	}
// 	return divRes
//
// }
