package main

import "fmt"

func main() {
	var contryCapitalMap map[string]string
	// 创建集合
	contryCapitalMap = make(map[string]string)
	// map插入 key-value 对各个国家对应的首都
	contryCapitalMap["france"] = "paris"
	contryCapitalMap["italy"] = "rome"
	contryCapitalMap["japan"] = "tokyo"
	contryCapitalMap["india"] = "new delhi"

	// 使用key 输出map的值
	for country := range contryCapitalMap {
		fmt.Println(" capital of", country, "is", contryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	captial, ok := contryCapitalMap["United States"]
	// 存在 ok 是true 否则不存在
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	// delete 函数
	countryCapitalMap1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	fmt.Println("delete 原始 map")
	// 打印map
	for country := range countryCapitalMap1 {
		fmt.Println("delete Capital of", country, "is", countryCapitalMap1[country])
	}

	// 删除元素
	delete(countryCapitalMap1, "france")
	fmt.Println("delete Entry for France is deleted")
	fmt.Println("delete 删除元素后 map")

	// 打印map
	for country := range countryCapitalMap1 {
		fmt.Println("delete Capital of", country, "is", countryCapitalMap1[country])
	}

}
