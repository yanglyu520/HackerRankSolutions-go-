func lonelyinteger(a []int32) int32 {
    var res int32
    res = 0
    
    for _,num := range a {
        res ^= num
    }
    return int32(res) 
}
