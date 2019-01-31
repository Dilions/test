package main

import ("fmt"
		// "reflect" //变量类型
		// "time"
		// "math/rand"
)
type population struct {
   name string
   work string
   work_placex int
   work_placey int
   age int
   move_distance int 
   hunger int
   live_placex int
   live_placey int
   skills []string
}
func main() {
	// var daycount int = 0 
	var pop [100]population
	maintaince(& pop)
	// var distance int =4
	// var judge int = distance%2
	// fmt.Println(judge)
	// if judge ==1 {
	// 	fmt.Println("yes")
	// }else{
	// 	fmt.Println("no")
	// }

	// fmt.Println(rand.Intn(100),pop[98].hunger)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(daycount)
	// 	daycount++
	// }
	fmt.Println(len(pop[0].skills))
	pop[0].skills = append(pop[0].skills,"noen")
	fmt.Println(len(pop[0].skills),pop[0].skills)
	var x []int
	fmt.Println(x)
	x=append(x,3)
	x=append(x,[]int{4,4,5,6,5,1}...)
	x=append(x,5)
	x=append(x,6)
	fmt.Println(x)
	x=x[0:0]
	fmt.Println(x)
}


func maintaince(pop *[100]population ) {
	// fmt.Println("land[i]")
	for i := 0; i < 99; i++ {
		pop[i].hunger=1
	}
	
	// fmt.Println(pop.hunger)
	
	// return 1
}






