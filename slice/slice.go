package slice

import "sort"

//添加权限去重复
func GetUniqPolices(polices [][]string) (uniqPolices [][]string) {
	mp := make(map[string]interface{}, 0)
	//polices := make([][]string, 0)
	for _, api := range polices {
		rKey := api[0]
		path := api[1]
		action := api[2]
		mapKey := rKey + "-" + path + "-" + action
		if mp[mapKey] != "" {
			mp[mapKey] = ""
			uniqPolices = append(uniqPolices, []string{rKey, path, action})
		}
	}
	return
}

// arrayInGroups 分组
func arrayInGroups(arr []int, num int64) [][]int {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

//分组
func arrayTwoGroups(arr [][]int, num int64) [][][]int {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][][]int{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][][]int, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

type student struct {
	name string
	age  int
}

func ArrayStuGroups(stu []student, num int64) [][]student {
	max := int64(len(stu))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]student{stu}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]student, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, stu[start:end])
		} else {
			segments = append(segments, stu[start:])
		}
		start = i * num
	}
	return segments
}

//按某个字段排序
type Person struct {
	Name string
	Age  int
}
type sortByAge []Person

func (s sortByAge) Len() int           { return len(s) }
func (s sortByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByAge) Less(i, j int) bool { return s[i].Age < s[j].Age }

//SliceSortAsc 切片排序
func SliceSortAsc(list []Person) []Person {
	sort.Sort(sortByAge(list))
	return list
}
