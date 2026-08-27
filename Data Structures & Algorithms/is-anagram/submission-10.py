class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        freqS = {}
        freqT = {}

        for char in s:
            val = freqS.get(char)
            if val:
                val = val+1
                freqS[char] = val
            else:
                freqS[char] = 1
        for char in t:
            val = freqT.get(char)
            if val:
                val+=1
                freqT[char] = val
            else:
                freqT[char] = 1
        
        for key, val in freqS.items():
            if freqT.get(key) != val:
                return False
        for key, val in freqT.items():
            if freqS.get(key)!= val:
                return False
        return True
        
        