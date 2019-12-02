func countTriplets(arr []int64, r int64) int64 {
    mymap := getMap(arr)
    var result []int64
    if r>1{
        for k,_:= range mymap{
            _,ok:=mymap[k*r]
            _,dk :=mymap[k*r*r]
            if ok && dk{
                entry := getEntry(mymap, k, r)
                result = append(result, entry)
            }
            
        }
    } else if r==1 {
        for _,v:=range mymap{
            if v[0]>=3{
                var val int64 = ((v)[0]*(v[0]-1)*(v[0]-2))/6
                //fmt.Println(val)
                result = append(result, val)  
            }
        }
    }
    var outPut int64 = 0
    for _,ans := range result{
        outPut += int64(ans) 
    }
    //fmt.Println(result)
    return outPut
}

func getMap(arr []int64) map[int64][]int64{
     m := make(map[int64][]int64)
    for k,v := range arr{
        if _,ok := m[int64(v)]; ok{
            m[int64(v)][0] = m[int64(v)][0]+int64(1)
            m[int64(v)] = append(m[int64(v)], int64(k))
        }else{
            m[int64(v)] = append(m[int64(v)], int64(1))
            m[int64(v)] = append(m[int64(v)], int64(k))
        }
    }
    //fmt.Println(m)
 return m   
}

func getEntry(m map[int64][]int64, key int64, r int64) int64{
    arr1 := m[key][1:len(m[key])]
    arr2 := m[key*r][1:len(m[key*r])]
    arr3 := m[key*r*r][1:len(m[key*r*r])]
    var entry int64 = 0
    for _,v1 := range arr1{
        for _,v2 := range arr2{
            if v2>v1{
                count := len(arr3)
                for _,v3 := range arr3{
                    if v3>v2 {
                        entry += int64(count)
                        break
                    } else {
                        count--
                        continue
                    }
                }
            }
        }
    }
    return entry
}
