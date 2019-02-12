package main

import ("fmt"
		// "reflect" //变量类型
		"time"
		"math/rand" //随机数
		"strconv"   //字符串
		"strings"
		"./pack"
)

const land_lenth = 20
const land_width = 20
const normal_move= 3
type landinfo struct {
	username string
	usedfor string
} 

type landstruct struct {
	icon 	string
	status 	string 	//business  house   factory   farm    none
	level 	int    	//各类别建筑的细分使用level还是用string呢再考虑
	record []landinfo
}

type pops struct {
	live 	int
	name 	string
	age 	int
	hunger 	int
	skills 	[]string

	work string
	work_placex []int
	work_placey []int
	
	live_placex		int
	live_placey		int
	move_distance	int 

	lifecircle	   int
	stuff_using  []int64  	//装配使用中
	stuff_taking []int64	//生产富余或闲置物资
}
var landinterface	[land_lenth][land_width]landstruct
var landsurface 	[land_lenth][land_width]string
var pop []pops

var debug_born_or_die = 1
var debug_age_report = 0
var debug_map =1

var date_year,date_month,date_day int = 0,0,0

var age_old int =50
var population int =1
var relocate_distance int=6

var overwatch_delay time.Duration= 1000000000//time.Sleep(1000000) //1000000000 = 1s

func main() {
	test.Name()
	var cop []test.Company
	var cop1 test.Company
	cop = append(cop,cop1)
	fmt.Println(len(cop))
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
    		landsurface[i][j]="`"
    		landinterface[i][j].icon="`"
    	}
    	fmt.Println(landsurface[i])
    	fmt.Println(landinterface[i])
    }
    // fmt.Println(s)
    var newpop pops
    newpop.name="bob#0"
    newpop.live_placex=10
    newpop.live_placey=10
	newpop.move_distance=5
	newpop.hunger=6
	newpop.work="none"
	newpop.age=16
	newpop.live=1
	pop = append (pop,newpop)


	// pop[0].skills.append("")
	// fmt.Println(len(pop))
    for i := 0; i < len(pop); i++ {
    	landinterface[pop[i].live_placex][pop[i].live_placey].status="house"
    	baseinfo := landinfo{pop[i].name,"live"}
    	landinterface[pop[i].live_placex][pop[i].live_placey].record=append(landinterface[pop[i].live_placex][pop[i].live_placey].record,baseinfo)
    	landsurface[pop[i].live_placex][pop[i].live_placey] = "A"
    }
    // fmt.Println(len(pop),"in loop")
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

	    	for i := len(pop)-1; i>=0; i-- {
	    		work(&pop[i])
	    	}
	    }
    }

}
//用于增加年龄
func age_maintaince() {
	for i := len(pop)-1; i >=0 ; i-- {
		if pop[i].live==1 {
			pop[i].age+=1
			if debug_age_report ==1 {
				fmt.Println(pop[i].name,"is",pop[i].age)
			}
			
			if (pop[i].age >= 18)&&(pop[i].age <= 30) { //生育年龄区间
				if(rand.Intn(365)<36){					//概率生育法
					var x,y=lozenge_search(pop[i].live_placex,pop[i].live_placey,3,"`")
					if x==-1 {							//没地儿下崽就举家搬迁 
						var x,y int=lozenge_search(pop[i].live_placex,pop[i].live_placey,relocate_distance,"`")
						if (x!=-1) &&( y!=-1) {			//搬迁成功
							
							fmt.Printf("\n住在%d %d没地儿下崽浑身难受, 搬迁在%d,%d\n",pop[i].live_placex,pop[i].live_placey,x,y)
							pop[i].live_placex,pop[i].live_placey=x,y
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
					population-=1
					if debug_born_or_die ==1 {
					fmt.Println( pop[i].live_placex,pop[i].live_placey,pop[i].name,"去世，享年",pop[i].age,randnum)
					}
					pop[i].live = 0
					landsurface[pop[i].live_placex][pop[i].live_placey]="`"
					landsurface[pop[i].work_placex[0]][pop[i].work_placey[0]]="`"
					landsurface[pop[i].work_placex[1]][pop[i].work_placey[1]]="`"
					landsurface[pop[i].work_placex[2]][pop[i].work_placey[2]]="`"
					landsurface[pop[i].work_placex[3]][pop[i].work_placey[3]]="`"
					pop[i].work_placex=pop[i].work_placex[0:0]
					pop[i].work_placey=pop[i].work_placey[0:0]

					pop = append(pop[:i], pop[i+1:]...) //删除人口信息

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
				landsurface[x0][y0]="+"
				var x1,y1 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
				if (x1!=-1) &&( y1!=-1) {
					landsurface[x1][y1]="+"
					var x2,y2 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
					if (x2!=-1) &&( y2!=-1) {
						landsurface[x2][y2]="+"
						var x3,y3 int=lozenge_search(pop.live_placex,pop.live_placey,pop.move_distance,"`")
						if (x3!=-1) &&( y3!=-1) {
							landsurface[x3][y3]="+"
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
									fmt.Println("     ",landsurface[i],i)
								}
								time.Sleep(overwatch_delay) 
							}

						}else{ //找不到工作举家搬迁
							landsurface[x0][y0]="`"
							landsurface[x1][y1]="`"
							landsurface[x2][y2]="`"
							var x,y int=lozenge_search(pop.live_placex,pop.live_placey,relocate_distance,"`")
							if (x!=-1) &&( y!=-1) {
								pop.live_placex,pop.live_placey=x,y
								fmt.Printf("\nfind no place to work, reloacte at %d,%d\n",x,y)
							}
						}
					}else{ //找不到工作举家搬迁
						landsurface[x0][y0]="`"
						landsurface[x1][y1]="`"
						var x,y int=lozenge_search(pop.live_placex,pop.live_placey,relocate_distance,"`")
						if (x!=-1) &&( y!=-1) {
							pop.live_placex,pop.live_placey=x,y
							fmt.Printf("\nfind no place to work, reloacte at %d,%d\n",x,y)
						}
					}
				}else{ //找不到工作举家搬迁
					landsurface[x0][y0]="`"
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
					fmt.Println("     ",landsurface[i],i)
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

//								xy坐标					行动距离 	标记符号
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
				if landsurface[x][y]==mark {
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
				if landsurface[x][y]==mark {
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
				if landsurface[x][y]==mark {
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
				if landsurface[x][y]==mark {
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
				if landsurface[x][y]==mark {
					return x,y
				}
			}
		}

	//loop 2/4
		for i := 0; i < distance; i++ {
			var x int=locationx+i
			var y int=locationy-distance+i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if landsurface[x][y]==mark {
					return x,y
				}
	
			}
		}
	//loop 3/4
		for i := 0; i < distance; i++ {
			var x int=locationx+distance-i
			var y int=locationy+i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if landsurface[x][y]==mark {
					return x,y
				}
	
			}
		}
	//loop 4/4
		for i := 0; i < distance; i++ {
			var x int=locationx-i
			var y int=locationy+distance-i
			if ((x >= 0) && (x < land_width)) && ((y >= 0)&&(y < land_lenth)) {
				if landsurface[x][y]==mark {
					return x,y
				}
	
			}
		}
	}
	return -1,-1	
}

func create_baby(fathername string,fatherchosenplacex,fatherchosenplacey int) {
	// var namehead string= "#"
	var a []string =strings.SplitN(fathername, "#",2)
	var b,error=strconv.Atoi(a[1])
	fmt.Println(error)
	var newpop pops
    newpop.name=a[0]+"#"+strconv.Itoa(b+1)
    newpop.live_placex=fatherchosenplacex
    newpop.live_placey=fatherchosenplacey
	newpop.move_distance=5
	newpop.hunger=6
	newpop.work="none"
	newpop.age=0
	newpop.live=1
	pop = append (pop,newpop)

	// for i := 0; i < 100; i++ {
	// 	if pop[i].live != 1 {
			
	// 		if debug_born_or_die ==1 {
	// 		fmt.Println("baby born")}
	// 		pop[i].live_placex=fatherchosenplacex
	// 		pop[i].live_placey=fatherchosenplacey
	// 		pop[i].live =1
	// 		pop[i].age =0 
	// 		pop[i].name = fathername+namehead+strconv.Itoa(i) 
	// 		pop[i].hunger=10
	// 		pop[i].work = "none"
	// 		pop[i].move_distance=3
			landsurface[fatherchosenplacex][fatherchosenplacey]="A"
			population+=1
	// 		if debug_map ==1 {
	// 			fmt.Println("---------",date_year,"年",date_month,"月",date_day,"日，人口：",population,"--------")
	// 			for i := 0; i < land_lenth; i++ {
	// 				fmt.Println("     ",landsurface[i],i)
	// 			}
	// 			time.Sleep(overwatch_delay) 
	// 		}
	// 		break
	// 	}
	// }
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

type Company struct {
	live 		bool
	name 		string
	age 		int
	leader 	[]	string
	stockcode 	int 	//股票代码，也可作为集团统一识别码

	date_setup 	string

	product 	string
	work_placex int
	work_placey int

	worker 	[]	string

	distance_resource 	int
	distance_sells 		int
	
	stuff_hardware	[]	int64

	stuff_resource  []	int64  	//装配使用中
	stuff_product 	[]	int64		//生产富余或闲置物资

	cash 				int64
	debt 				int64
	record_output 	[]	string
	record_income 	[]	string
}

func dailymaintance() {
	for i := 0; i < len(pop); i++ {
		for j := 0; j < len(pop[i].stuff_using); j++ {
			pop[i].stuff_using[j]-=1
			if (pop[i].stuff_using[j]%100000<=pop[i].lifecircle) { 
				尝试搜索购买
				if 买到了 {
					修改持有列表
				}else{
					发布需求
				}

				if (pop[i].stuff_using[j]%90000==1) { //报废  pop = append(pop[:i], pop[i+1:]...)
					//删除物品
					pop[i].stuff_using[j] = append(pop[i].stuff_using[:j], pop[i].stuff_using[j+1:]...)
				}	
			}
			
		}
	}
	
}








