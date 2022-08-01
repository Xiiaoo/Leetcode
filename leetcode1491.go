package main

import "fmt"

func average(salary []int) float64 {
	min,max,count:=salary[0],salary[0],salary[0]
	for i := 1; i < len(salary); i++ {
		count=count+salary[i]
		if salary[i]<salary[i-1] && salary[i]<min{
			min=salary[i]
		}else if salary[i]>salary[i-1] && salary[i]>max {
			max=salary[i]
		}
	}
	return float64(float64((count-max-min))/float64((len(salary)-2)))
}

func main() {
	fmt.Println(average([]int{48000,59000,99000,13000,78000,45000,31000,17000,39000,37000,93000,77000,33000,28000,4000,54000,67000,6000,1000,11000}))
}