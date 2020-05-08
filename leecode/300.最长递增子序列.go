package main

func main() {

}

/*
 *
    300. 最长上升子序列
    给定一个无序的整数数组，找到其中最长上升子序列的长度。

    示例:

    输入: [10,9,2,5,3,7,101,18]
    输出: 4 
    解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
    说明:

    可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
    你算法的时间复杂度应该为 O(n2) 。
    进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/
func lengthOfLIS(nums []int) int {
    //初始化dp数组的值均为1
    //dp[i]表示下标i对应的值num[i]的最长递增子序列的个数
    max := 0
    dp := make([]int, len(nums))
    for i := 0; i < len(nums);i++ {
        dp[i] = 1
        //对于小于num[i]的值
        for j := 0; j < i; j++ {
            //只要num[j]的值小于num[i]
            if nums[i] > nums[j] {
                //那么在j位置的最长递增子序列个数+1，即可能为i位置的最长递增子序列的值，我们选取最大的值作为dp[i]
                if dp[i] < dp[j] + 1 {
                    dp[i] = dp[j] + 1
                }
            }
        }
        //dp数组中，最大的值，即为最长递增子序列
        if dp[i] > max {
            max = dp[i]
        }
    }
    return max
}

func lengthOfLIS2(nums []int) int {
    //初始化dp数组的值均为1
    dp := make([]int, len(nums))
    for i:= 0; i < len(nums);i++ {
        dp[i] = 1

        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    //
    maxVal := max(dp...)
    return maxVal
}

func max(nums ...int) int {
    maxVal := 0
    for _,v := range nums {
        if maxVal < v {
            maxVal = v
        }
    }
    return maxVal
}