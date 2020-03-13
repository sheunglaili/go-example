package sum

//Sum sum number in array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
// SumAll sum all number in all array provided and put them into one array
func SumAll(numbersToSum ...[]int) ( []int) {
   var sums []int 
   for _ , numbers := range numbersToSum {
	   sums = append(sums,Sum(numbers))
   }
   return sums
}

//SumAllTails sum all number except the first number in all array provided and put them into one array
func SumAllTails(numbersToSum ...[]int) ([]int) {

	var sums []int 
	for _,numbers := range numbersToSum {
		if(len(numbers) == 0 ){
			sums = append(sums,0)
		}else {
			tail := numbers[1:]
			sums = append(sums,Sum(tail))
		}
	}

	return sums
}