class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        # sort in ascending - check pairs of numbers 
        # or
        # create a hash map of numbers, if the number exists in the map already then return 
        result = {}
        for num in nums:
            if result.get(num) is not None :
                return True
            else:
                result[num] = True
        return False

        