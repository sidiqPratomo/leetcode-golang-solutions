package twosums

func twoSum(nums []int, target int) []int {
    for i:=0; i<= len(nums) -1 ; i++{
        for j:= 0; j<i; j++{
            if nums[i]+nums[j] == target{
                return []int{j,i}
            }
        }
    }
    return nil
}