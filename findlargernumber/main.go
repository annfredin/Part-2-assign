package main

import (
	"fmt"
)

func main() {	
	
	fmt.Println(findLargerNumberBySwap("3579"))
}


func findLargerNumberBySwap(s string) interface{} {
	
	res := []byte(s)

	for i := 0; i < len(res); i++ {
		ind := -1
		for j := i+1; j < len(res); j++ {
			left := int(res[i])
			right := int(res[j])

			if (left % 2 == 0 && right % 2 == 0 ) || (left % 2 != 0 && right % 2 != 0 ){
				if left < right{
				if ind == -1{
					ind = j
				}else{
					if int(res[ind]) < right{
						ind = j
					}
				}
				}
			}else{
				break
			}
		}

		if ind >=0{
		res[ind],res[i] = res[i], res[ind]
		}
		
	}
	
	return string(res)
}
