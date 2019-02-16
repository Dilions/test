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
	id		int64   //身份证号
	age 	int
	hunger 	int
	edu		int
	happy	int
	health 	int

	character	[]int 	//特性也得做成列表，通过索引去找或者直接用数组代替
	skills 		[]int 	//工作技能

	work string
	work_placex []int
	work_placey []int
	
	live_placex		int
	live_placey		int
	move_distance	int 

	life_level		int 	//生活等级，调控需求
	life_circle	   	int 	//生活周期，与职业财产等相关，调控需求

	stuff_using  []	int64  	//装配使用中
	stuff_taking []	int64	//生产富余或闲置物资

	stuff_mental []	int64   //生活精神等非消耗物资
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
	boss 		string
	leader 	[]	string 	//董事会
	stockcode 	int 	//股票代码，也可作为集团统一识别码

	date_setup 	string

	product 	int //procedure工序   formula配方
	product_resource []int64
	product_madeup	 []int64
	work_placex int
	work_placey int

	worker 	[]	string

	distance_resource 	int
	distance_sells 		int
	
	stuff_hardware	[]	int64

	stuff_resource  []	int64  	//准备用作生产的物资
	stuff_product 	[]	int64	//生产富余或闲置物资
	stuff_progress 	int

	//记录一年的条目，每半年检查一次，删除超出一年的记录
	cash 				int64
	debt 				int64
	record_output 	[]	string //支出 记录格式 品类#数量#单价 
	record_income 	[]	string //收入 记录格式 品类#数量#单价
}

/*--
考虑下来发现物物交易和金钱交易似乎并不互相排斥
而物物交换的核心是找到物品里的“硬通货”
目前看来似乎食用稻米比较合适
还需要考虑金钱如何流通到大家手里这个问题

//还需要考虑消费升级
*/
func maintance_personal() {   //个人消费消耗
	//日常损耗
	for i := 0; i < len(pop); i++ {
		for j := 0; j < len(pop[i].stuff_using); j++ {
			pop[i].stuff_using[j]-=1
			if (pop[i].stuff_using[j]%100000<=pop[i].life_circle) { 
				尝试搜索周边并购买
				if 买到了 {
					修改持有列表
						没有则添加新项目并重新排列
						目前有相应物资的情况就找到相应项目来增加其耐久度
				}else{
					发布所缺物品需求
				}

				if (pop[i].stuff_using[j]%90000==1) { //报废  pop = append(pop[:i], pop[i+1:]...)
					//删除物品
					pop[i].stuff_using[j] = append(pop[i].stuff_using[:j], pop[i].stuff_using[j+1:]...)
				}	
			}
		}
	}

	//特殊需求损耗
	for i := 0; i < len(pop); i++ {
		检查当下可交易物资
		检查现有第三需求物资存量 ， 购买周期同其他物品相似，通过耐久来控制周期，
		比如教育周期或者医疗疗程，甚至是某周口感吃腻了，用耐久值来模拟娱乐兴奋度的消减，同时通过权值来控制兴奋度调剂购买选择趋向（权值为满且周期极短则为毒品）
			优先满足 饥饿（口感）， 健康 ，  娱乐  ， 教育
	}
}

var Companys Company
//企业创立
//此函数调用于条件判断之后的企业数据初始化
func Company_setup(leadername string,setup_posex int,setup_posey int,object int) {

	var newcompany Company
	newcompany.live=true
	newcompany.name=leadername+"'s inc"
	newcompany.boss=leadername

	newcompany.date_setup=strconv.Itoa(date_year)+"-"+strconv.Itoa(date_month)+"-"+strconv.Itoa(date_day)

	newcompany.work_placex=setup_posex
	newcompany.work_placey=setup_posey

	newcompany.product=object

	newcompany.distance_resource=4
	newcompany.distance_sells	=4

	加载配方

	Companys = append(Companys, newcompany)
}
/*
考虑到企业生产用硬件设备有的是增加效率的，有的是生产必要条件。那么按照当下的命名制度要区别这些东西只能通过物品类别编号来区分了
。。。还有很多事情要做啊。。。 
*/

//从外部数据company大类中提取数据进行维护
/*一次性的原材料采购数量应该设为多少，和什么相关？库存呢
考虑到企业间可能是天差地别的，需要一个覆盖面广的标准，采购数量和库存应根据销售数据来计算，
销售数据“平均数”应该和产品价值，工期长短，上周或上半月或上个月的销量来计算其 “基准数”
而实际采购量和产品库存应在基准数的基础上加2或者5或者10？（工期越短加成数越大，价值越高加成数量越小）

*/
func maintance_company() {
	for i := 0; i < len(Companys); i++ {
		
	
	更新生产进度
		根据设备情况、人员数量、技能水准等因素计算生产速率 来推进生产（点数）
		*生产点数=基数*人数*技能水平+基数*设备数量*设备性能
		如果一轮生产结束
			用除法和余数来更新原材料每次使用的数量（份）
			更新原材料和产品库存
			概率增加技术水平熟练度

	检查销售情况
		计算库存额定值
	检查原材料库存 	于额定值对比
		不够就去购买
			买不全就发布需求
			买不到就发布需求
	检查产品库存		于额定值对比
		不足则继续生产
		已经满足则停产

	检查日期发放人员工资

	/*
		公司人员结构怎么处理？？？多少岗位，什么岗位多少人怎么办？？？
		不同 的公司人员结构天差地别，该怎么合理的分配
		**	或许不同类型的公司就该不同模型处理。。。。
			生产制造：农林牧矿，制造合成加工，电力燃气供水，建筑。。。
			服务：运输，娱乐，销售，医疗，教育，计算机软件，科学研究。。。

			生产制造这些都好说，产品是实物，需要这个东西，但是服务业的产品多种多样，这个对于其他人其他公司的影响就非常复杂了
			比如有的是提高生产效率的，有的是维护个人属性如健康，饥饿，娱乐等。健康饥饿娱乐都可以量化购买，那么运输...软件，研究，教育...
		*	教育值？来增加技能水平？那怎么让这些人意识到受教育的好处而去主动选择并接受

			软件或者研究提供的产品同医疗娱乐等都是属性类的，应该在正规生产之外再开个分立的判断来确定是否购买。
		***	生产需求和提升需求分立！

			如何把第三产业产品性能体现在标识中 类（属性） 编号 权值（购买意愿考量和产品效果1-999） 数量（教育产品中表示等级） 耐久（使用时间，教育产品中指代本次受教育时长）
												01	   12 	958  2
		用菱形搜索的时候应该有每日搜索次数上限
		人员应带有属性选项
	*/
	}
	
}






