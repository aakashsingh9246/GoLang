func freqQuery(queries [][]int32) []int32 {
    var output []int32
    temp := make(map[int32]int32)
    temp2:=make(map[int32]int32)
    for _,val:=range queries{
        if val[0] == 1{
            _,ok := temp2[temp[val[1]]]
            if ok{
                delete(temp2, temp[val[1]])
            }
            temp[val[1]]++
            temp2[temp[val[1]]] = val[1]
        }else if val[0]==2{
            _,ok:=temp[val[1]]
            if ok{
                _,ok := temp2[temp[val[1]]]
                if ok{
                delete(temp2, temp[val[1]])
            }
                temp[val[1]]--
                temp2[temp[val[1]]]=val[1]
            }
        }else if val[0]==3{
           _,ok := temp2[val[1]]
           if ok{
               output=append(output,1)
           }else{
               output=append(output,0)
           }
        }
    }
    return output

}