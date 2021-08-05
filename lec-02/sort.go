package main

import (
	"fmt"
)

//implement buble sort
func bubleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		for j := i + 1; j < len(arr); j++ {

			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i] //swap(arr[j], arr[i])
			}

		}
	}

	return arr
}

//implement quick sort
func partition(arr []int, l int, r int) int {

	pivot := arr[r]         //pivot là phần tử cuối
	var i, j int = l - 1, l //i:0->i sẽ là những phần tử nhỏ hơn pivot
	for ; j <= r-1; j++ {
		if arr[j] <= pivot { //nhỏ hơn pivot thì swap với arr[i]
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	//swap pivot vs a[i+1] sẽ tách đc thành 2 mảng,một nhỏ hơn pivot ,1 lớn hơn pivot
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1

}

func quickSort(arr []int, l int, r int) []int {

	if l < r {
		p := partition(arr, l, r) //chia mảng theo pivot

		quickSort(arr, r, p-1) //sort mảng trái
		quickSort(arr, p+1, r) //sort mảng phải
	}

	return arr
}

//implement merge sort
func merge(arr []int, l int, r int) {

	mid := l + (r-l)/2
	leftLen := mid - l + 1
	rightLen := r - mid

	left := arr[l:mid]      //mảng trái từ l đến mid
	right := arr[mid+1 : r] //mảng phải từ mid+1 đến r

	i, j, k := 0, 0, l

	//merge cho đến khi 1 trong 2 mảng kết thúc
	for i < leftLen && j < rightLen {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	//Merge nốt phần còn lại nêu có
	for i < leftLen {
		arr[k] = left[i]
		i++
		k++
	}

	for j < rightLen {
		arr[k] = right[j]
		j++
		k++
	}

}

func mergeSort(arr []int, l int, r int) []int {

	if l < r {

		mid := l + (r-l)/2

		mergeSort(arr, l, mid)

		//đệ quy các mảng con
		merge(arr, mid+1, r)
		merge(arr, l, r)

	}
	return arr
}

func main() {

	arr := []int{12, 23, 2, 4, 2, 13, 43, 12, 32, 32}

	fmt.Println("before sort:", arr)

	fmt.Println("buble sort:", bubleSort(arr))

	fmt.Println("quick sort:", quickSort(arr, 0, len(arr)-1))

	fmt.Println("quick sort:", mergeSort(arr, 0, len(arr)-1))

}
