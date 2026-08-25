class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        # sort in ascending - check pairs of numbers 
        # or
        # create a hash map of numbers, if the number exists in the map already then return 
        result = {}
        for num in nums:
                #python tip - dict.get() does not return key error as result[num] ( which can be used with try catch)
            if result.get(num) is not None :
                return True
            else:
                result[num] = True
        return False

        