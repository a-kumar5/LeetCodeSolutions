class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        values = {}
        for id, num in enumerate(nums):
            if target - num in values:
                return [values[target-num],id]
            values[num] = id
        