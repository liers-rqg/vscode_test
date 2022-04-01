package demo3
import (
	"fmt"
)
//map操作

func InitMyMap(my_map map[int]string){
	
	my_map[0]="demo1"
	my_map[1]="demo2"
	my_map[2]="demo3"
	fmt.Println(my_map)
}
func DeleteMap(my_map map[int]string,key int){
	delete(my_map,key)
	fmt.Println(my_map)
}
func FindKeyInMap(my_map map[int]string,key int){
	val,isInMap := my_map[key]
	if isInMap {
		fmt.Printf("get key %v\n",val)
	}else {
		fmt.Println("no key")
	}
}
func ReadMap(my_map map[int]string){
	for k,v := range my_map {
		fmt.Printf("key = %v, val = %v\n",k,v)
	}
}
func SlipOfMap(my_map map[int]string){
	var my_maps []map[int]string
	my_maps = make([]map[int]string,2)
	if my_maps[0] == nil {
		my_maps[0] = make(map[int]string)
		my_maps[0] = my_map
	}
	if my_maps[1] == nil{
		my_maps[1] = make(map[int]string)
		my_maps[1][0] = "learn01"
		my_maps[1][1] = "learn02"
	}
	newMap := map[int]string{
		0 : "base01",
		1 : "base02",
	}
	my_maps = append(my_maps,newMap)
	fmt.Println(my_maps)
}