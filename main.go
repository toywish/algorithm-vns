package main

import (
	"fmt"
	"math/rand"
)

var (
	//物品总数
	num = 100
	//背包容量
	maxWeight = num * 5
	//物品的重量和价值
	values, weights []int
	//最大邻域数量
	maxFlip = 3
	//迭代次数
	iterator = 1000
	//随机数种子
	seed int64 = 5113
	//邻域解决方案
	neighborhoodSolution [][]int
	//方案检查总数
	solutionCount = 0
	//方案替换总数
	solutionUpdate = 0
	//记录所有方案的物品价值和重量
	solutionValue, solutionWeight = make([]int, 0), make([]int, 0)
)

func main() {
	//设置随机数种子
	rand.Seed(seed)
	//初始化每个物品的价值
	values = randSlice(num, 10, 100)
	weights = randSlice(num, 5, 10)

	initCheck()
	//初始化解决方案
	solution := randSlice(num, 0, 1)
	solutionValue = append(solutionValue, getTotal(solution, values))
	solutionWeight = append(solutionWeight, getTotal(solution, weights))
	//给邻域空间分配内存
	neighborhoodSolution = make([][]int, num)
	fmt.Println("init solution")
	total(solution)

	best := vns(solution, iterator)
	fmt.Println("solution checked total:", solutionCount)
	fmt.Println("solution updated total:", solutionUpdate)

	fmt.Println("result solution")
	total(best)
	fmt.Printf("Press Enter To End...")
	fmt.Scanln()
}

//打印初始物品的总价值和总重量
func initCheck() {
	fmt.Println("***************************")
	v, w := 0, 0
	for i := 0; i < num; i++ {
		v += values[i]
		w += weights[i]
	}
	fmt.Println("values", values, "\n", "weights", weights)
	fmt.Printf("value total: %d, weight total: %d\n", v, w)
}

//打印解决方案的物品价值和重量
func total(solution []int) {
	v, w := 0, 0
	for i := range solution {
		v += values[i] * solution[i]
		w += weights[i] * solution[i]
	}
	fmt.Printf("value total: %d, weight total: %d\n", v, w)
}

//生成长度为len的随机数数组
func randSlice(len int, start int, end int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = rand.Intn(end-start+1) + start
	}
	return s
}

//统计总重量或总价值
func getTotal(solution []int, val []int) int {
	sum := 0
	for i := range solution {
		sum += solution[i] * val[i]
	}
	return sum
}

//评价解决方案，总重量超过上限则返回false
func evaluateSolution(solution []int) bool {
	return getTotal(solution, weights) <= maxWeight
}

//随机扰动过程
func shakingProcess(solution []int) {
	for i := range solution {
		solution[i] = rand.Intn(2)
	}
}

//vns算法
func vns(solution []int, iterator int) []int {
	best := make([]int, num)
	copy(best, solution)
	//fmt.Println(best)
	vnd(best)
	//fmt.Println(best)
	for ; iterator > 0; iterator-- {
		shakingProcess(solution)
		vnd(best)
	}
	return best
}
func vnd(solution []int) {
	flip := 0

	for flip < maxFlip {
		neighborhood(solution, flip)
		curr := make([]int, num)
		//fmt.Println(flip, len(neighborhoodSolution))
		//fmt.Println(neighborhoodSolution)
		copy(curr, solution)
		flag := false
		for _, v := range neighborhoodSolution {
			//fmt.Println(v)
			solutionCount++
			solutionValue = append(solutionValue, getTotal(v, values))
			solutionWeight = append(solutionWeight, getTotal(v, weights))
			if evaluateSolution(v) && getTotal(v, values) > getTotal(solution, values) {

				copy(curr, v)
				flag = true
			}
		}
		if flag {
			copy(solution, curr)
			solutionUpdate++
			flip = 0
		} else {
			flip++
		}
	}
}

//生成邻域
func neighborhood(solution []int, flip int) {
	//清空邻域
	neighborhoodSolution = neighborhoodSolution[0:0]
	switch flip {
	case 0:
		for i := range solution {
			s := make([]int, num)
			copy(s, solution)
			//fmt.Println(s, solution)
			if solution[i] == 1 {
				s[i] = 0
			} else {
				s[i] = 1
			}
			neighborhoodSolution = append(neighborhoodSolution, s)
			//fmt.Println("s", s)
			//fmt.Println("neighborhoodSolution", neighborhoodSolution)
		}
	case 1:
		for i := range solution {
			for j := i + 1; j < len(solution); j++ {
				s := make([]int, num)
				copy(s, solution)
				if solution[i] == 1 {
					s[i] = 0
				} else {
					s[i] = 1
				}
				if solution[j] == 1 {
					s[j] = 0
				} else {
					s[j] = 1
				}
				neighborhoodSolution = append(neighborhoodSolution, s)
				//fmt.Println("s", s)
				//fmt.Println("neighborhoodSolution", neighborhoodSolution)
			}
		}
	case 2:
		for i := range solution {
			s := make([]int, num)
			copy(s, solution)
			l := len(s)
			if i < 3 {
				s[i], s[l-1-i] = s[l-1-i], s[i]
			} else {
				s[i], s[i-3] = s[i-3], s[i]
			}
			neighborhoodSolution = append(neighborhoodSolution, s)
		}

	}
}
