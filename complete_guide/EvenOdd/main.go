package main


func main()  {
	
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 11; i++ {
		if nums[i]%2 == 0 {
			println(nums[i], "is even")
		} else {
			println(nums[i], "is odd")
		}
	}
}