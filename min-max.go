func miniMaxSum(arr []int32)  {
    arrCopy := make([]int, len(arr))
    
    for i, num := range arr{
        arrCopy[i] = int(num)
    }
    
    sort.Ints(arrCopy)
    // Calculate the minimum and maximum sums
    minSum := arrCopy[0] + arrCopy[1] + arrCopy[2] + arrCopy[3]
    maxSum := arrCopy[len(arr)-1] + arrCopy[len(arr)-2] + arrCopy[len(arr)-3] + arrCopy[len(arr)-4]

    // Print the minimum and maximum sums
    fmt.Printf("%d %d\n", minSum, maxSum)  
}
