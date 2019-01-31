package main

import ("fmt"
		// "reflect" //变量类型
		"time"
		"math/rand" //随机数
		"strconv"   //字符串

)

const land_lenth = 20
const land_width = 20
const normal_move= 3

type pops struct {
	live int
	name string
	age int
	hunger int
	skills []string

	work string
	work_placex []int
	work_placey []int
	
	live_placex int
	live_placey int
	move_distance int 
}

var land [land_lenth][land_width]string
var pop [100]pops

var debug_born_or_age = 0
var debug_map =1

var date_year,date_month,date_day int = 0,0,0

var age_old int =50
var population int =0
var relocate_distance int=6

var overwatch_delay time.Duration= 1000000000//time.Sleep(1000000) //1000000000 = 1s

func main() {

/*	const LOOP_COUNT = 5
	const LOOP_NUM = 10
    var s [][]string
    for i := 0; i < LOOP_COUNT; i++ {
        sl := make([]string,0,LOOP_NUM)
        for j := 0; j < LOOP_NUM; j++ {
            sl = append(sl,"a")
        }
        s = append(s,sl)
    }
    fmt.Println(s[0],"1\b\n,",s[1],s)
*/ 
    var day_count int = 0
//-------------- init ------------------------------
    
    for i := 0; i < land_lenth; i++ {
    	for j := 0; j < land_width; j++ {
    		land[i][j]="`"
    	}
    	fmt.Println(land[i])
    }
    // fmt.Println(s)
    
    pop[0].name="bob"
    pop[0].live_placex=10
    pop[0].live_placey=10
	pop[0].move_distance=5
	pop[0].hunger=6
	pop[0].work="none"
	pop[0].age=16
	pop[0].live=1
	// pop[0].skills.append("")

    for i := 0; i < 1; i++ {
    	land[pop[i].live_placex][pop[i].live_placey] = "A"
    }
//----------------- loop ------------------------------------
    for true {
		// time.Sleep(1000000) //1000000000 = 1s
//输出土地使用情况
		
		calendar()

	    day_count++
	    // fmt.Println(day_count,)
	    if day_count==365 {
	    	day_count = 0
	    	age_maintaince()
	    }
	    
	    if day_count%30==0 {
	    	for i := 0; i < 100; i++ {
	    		work(&pop[i])
	    	}
	    }
		//if pop[0].hunger<=5 {
		//		work(&pop[0])
		//}
    }

}
//用于增加年龄
func age_maintaince() {
	for i := 0; i < 100; i++ {
		if pop[i].live==1 {
			pop[i].age+=1
			if debug_born_or_age ==1 {
				fmt.Println(pop[i].name,"is",pop[i].age)
			}
			
			if (pop[i].age >= 18)&&(pop[i].age <= 30) { //生育年龄区间
				if(rand.Intn(365)<36){					//概率生育法
					var x,y=lozenge_search(pop[i].live_placex,pop[i].live_placey,3,"`")
					if x==-1 {							//没地儿下崽就举家搬迁 
						var x,y int=lozenge_search(pop[i].live_placex,pop[i].live_placey,relocate_distance,"`")
						if (x!=-1) &&( y!=-1) {			//搬迁成功
							pop[i].live_placex,pop[i].live_placey=x,y
							fmt.Printf("\n住在%d %d没地儿下崽浑身难受, 搬迁在%d,%d\n",pop[i].live_placex,pop[i].live_placey,x,y)
						}else{							//搬迁失败，只能忍着暂时
							fmt.Println("住在",pop[i].live_placex,pop[i].live_placey,"没地儿下崽浑身难受")
						}

						
					}else{
						create_baby(pop[i].name,x,y)
					}
				}			
			}

			if pop[i].age > 50 {
				var randnum=rand.Intn(50)
				if (randnum<(pop[i].age-50)) {
					fmt.Println( pop[i].live_placex,pop[i].live_placey,pop[i].name,"去世，享年",pop[i].age)
					pop[i].live = 0
					land[pop[i].live_placex][pop[i].live_placey]="`"
					land[pop[i].work_placex[0]][pop[i].work_placey[0]]="`"
					land[pop[i].work_placex[1]][pop[i].work_placey[1]]="`"
					land[pop[i].work_placex[2]][pop[i].work_placey[2]]="`"
					land[pop[i].work_placex[3]][pop[i].work_placey[3]]="`"
					pop[i].work_placex=pop[i].work_placex[0:0]
					pop[i].work_placey=pop[i].work_placey[0:0]
				}
			}	
		}
		
	}
}
//单个人找工作用的
func work(pop *pops) {

	if pop.work == "none" {
		if len(pop.skills)==0 {
			var x0,y0 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
			if (x0!=-1) &&( y0!=-1) {
				land[x0][y0]="+"
				var x1,y1 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
				if (x1!=-1) &&( y1!=-1) {
					land[x1][y1]="+"
					var x2,y2 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
					if (x2!=-1) &&( y2!=-1) {
						land[x2][y2]="+"
						var x3,y3 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
						if (x3!=-1) &&( y3!=-1) {
							land[x3][y3]="+"
							pop.work_placex=append(pop.work_placex,[]int{x0,x1,x2,x3}...)
							pop.work_placey=append(pop.work_placey,[]int{y0,y1,y2,y3}...)
							// pop.work_placex[0],pop.work_placey[0]=x0,y0
							// pop.work_placex[1],pop.work_placey[1]=x1,y1
							// pop.work_placex[2],pop.work_placey[2]=x2,y2
							// pop.work_placex[3],pop.work_placey[3]=x3,y3
							pop.work="hunter" //四块地皮，否则无法生存

							if debug_map ==1 {
								fmt.Println("---------",date_year,"年",date_month,"月",date_day,"日，人口：",population,"--------")
								for i := 0; i < land_lenth; i++ {
									fmt.Println("     ",land[i],i)
								}
								time.Sleep(overwatch_delay) 
							}

						}else{ //找不到工作举家搬迁
							land[x0][y0]="`"
							land[x1][y1]="`"
							land[x2][y2]="`"
							var x,y int=lozenge_search(pop.live_placex,pop.live_placey,relocate_distance,"`")
							if (x!=-1) &&( y!=-1) {
								pop.live_placex,pop.live_placey=x,y
								fmt.Printf("\nfind no place to work, reloacte at %d,%d\n",x,y)
							}
						}
					}else{ //找不到工作举家搬迁
						land[x0][y0]="`"
						land[x1][y1]="`"
						var x,y int=lozenge_search(pop.live_placex,pop.live_placey,relocate_distance,"`")
						if (x!=-1) &&( y!=-1) {
							pop.live_placex,pop.live_placey=x,y
							fmt.Printf("\nfind no place to work, reloacte at %d,%d\n",x,y)
						}
					}
				}else{ //找不到工作举家搬迁
					land[x0][y0]="`"
					var x,y int=lozenge_search(pop.live_placex,pop.live_placey,relocate_distance,"`")
					if (x!=-1) &&( y!=-1) {
						pop.live_placex,pop.live_placey=x,y
						fmt.Printf("\nfind no place to work, reloacte at %d,%d\n",x,y)
					}
				}	
			}else{ //找不到工作举家搬迁
				var x,y int=lozenge_search(pop.live_placex,pop.live_placey,relocate_distance,"`")
				if (x!=-1) &&( y!=-1) {
					pop.live_placex,pop.live_placey=x,y
					fmt.Printf("\nfind no place to work, reloacte at %d,%d\n",x,y)
				}
			}	
			

		}else{ //技能不为0的情况！！！
			if debug_map ==1 {
				fmt.Println("---------",date_year,"年",date_month,"月",date_day,"日，人口：",population,"--------")
				for i := 0; i < land_lenth; i++ {
					fmt.Println("     ",land[i],i)
				}
				time.Sleep(overwatch_delay) 
			}
		}
	}
}

func abs(n int) int{
	if n <0 {
		return -n
		}	
	return n
}

/*
	[` ` ` ` ` 5 ` ` ` `]
	[` ` ` ` 5 ` 5 ` ` `]
	[` ` ` ` 4 3 ` ` ` `]
	[` ` ` 4 3 2 3 ` ` `]
	[` 5 4 3 2 1 2 3 ` `]
	[5 4 3 2 1 A 1 2 3 `]
	[` ` ` 3 2 1 2 3 ` `]
	[` ` ` ` 3 2 3 ` ` `]
	[` ` ` ` ` 3 ` ` ` `]
	[` ` ` ` ` ` ` ` ` `]
*///																	xy坐标					行动距离 	标记符号
func lozenge_search(locationx int,locationy int,distance_max int,mark string) (int,int){
	// var judge int = distance%2
	// if judge == 1 {  //单数菱形

	for distance := 1; distance <= distance_max; distance++ {
		// var luck int = 
	switch rand.Intn(3) {
	case 0:
	//method 1/4    
		for i := 0; i < distance; i++ {
			var x int=locationx-distance+i
			var y int=locationy-i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
			}
		}
	case 1:
	//method 2/4
		for i := 0; i < distance; i++ {
			var x int=locationx+i
			var y int=locationy-distance+i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
	
			}
		}
	case 2:
	//method 3/4
		for i := 0; i < distance; i++ {
			var x int=locationx+distance-i
			var y int=locationy+i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
	
			}
		}
	case 3:
	//method 4/4
		for i := 0; i < distance; i++ {
			var x int=locationx-i
			var y int=locationy+distance-i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
	
			}
		}
	}
	/////////////////////////////////////////////////////////////////////////////////
		// start:
	//loop 1/4    
		for i := 0; i < distance; i++ {
			var x int=locationx-distance+i
			var y int=locationy-i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
			}
		}

	//loop 2/4
		for i := 0; i < distance; i++ {
			var x int=locationx+i
			var y int=locationy-distance+i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
	
			}
		}
	//loop 3/4
		for i := 0; i < distance; i++ {
			var x int=locationx+distance-i
			var y int=locationy+i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
	
			}
		}
	//loop 4/4
		for i := 0; i < distance; i++ {
			var x int=locationx-i
			var y int=locationy+distance-i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if land[x][y]==mark {
					return x,y
				}
	
			}
		}
	}
	return -1,-1	
}

func create_baby(fathername string,fatherchosenplacex,fatherchosenplacey int) {
	var namehead string= "#0"
	for i := 0; i < 100; i++ {
		if pop[i].live != 1 {
			
			if debug_born_or_age ==1 {
			fmt.Println("baby born")}
			pop[i].live_placex=fatherchosenplacex
			pop[i].live_placey=fatherchosenplacey
			pop[i].live =1
			pop[i].age =0 
			pop[i].name = fathername+namehead+strconv.Itoa(i) 
			pop[i].hunger=10
			pop[i].work = "none"
			pop[i].move_distance=3
			land[fatherchosenplacex][fatherchosenplacey]="A"
			
			if debug_map ==1 {
				fmt.Println("---------",date_year,"年",date_month,"月",date_day,"日，人口：",population,"--------")
				for i := 0; i < land_lenth; i++ {
					fmt.Println("     ",land[i],i)
				}
				time.Sleep(overwatch_delay) 
			}
			break
		}
	}
}

func calendar() {
	date_day+=1
	if date_day == 31 {
		date_day=1
		date_month+=1
		if date_month==13 {
			date_month=1
			date_year+=1

			// population_cal() //人口统计
		}
	}
	
}

// func population_cal() {
// 	var cnt int =0
// 	for i := 0; i < 100; i++ {
// 		if pop[i].live==1 {
// 			cnt+=1
// 		}
// 	}
// 	population = cnt
// }





