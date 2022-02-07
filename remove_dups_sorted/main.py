# %%
import cmd

# remove duplicates from sorted list
# accepted


class Solution:
    def removeDuplicates(self, nums: list) -> int:
      if len(nums) <= 1:
        return len(nums)
      
      for i in range(len(nums)-2, -1, -1):
        cmpInd = i+1
        if nums[i] == nums[cmpInd]:
          nums.pop(cmpInd)

      return len(nums)
  
s = Solution()

#input = [1,1,2]
input = [0,0,1,1,1,2,2,3,3,4]
#input = [1,1,2,5,6,7,8]
print(input)
print(s.removeDuplicates(input))
print(input)
#input.pop(0)
#print(input)


# %%
