package ArrayAndSlice

import "fmt"

func TestOne()  {
	arr := [9]int {1,2,3,4,5,6,7,8,9}
	s := arr[1:4]
	fmt.Printf("Array length %d,cap:%d \n",len(arr),cap(arr)) //9-9
	fmt.Printf("slice length %d,cap:%d \n",len(s),cap(s)) // 3-8
}

func TestTwo()  {
	arr := [9]int {1,2,3,4,5,6,7,8,9}
	s := arr[:4]
	fmt.Printf("Array length %d,cap:%d \n",len(arr),cap(arr)) //9-9
	fmt.Printf("slice length %d,cap:%d \n",len(s),cap(s)) // 4-9
	// cap(s) = cap(arr) - leftV
	// len(s) = rightV - leftV
}
func TestThree()  {
	arr := [9]int {1,2,3,4,5,6,7,8,9}
	s := arr[:4]

	s  = append(s, 10)
	fmt.Println(s) //只要末尾元素不是原数组的尾端，修改会直接影响原数组
	fmt.Println(arr)
	}