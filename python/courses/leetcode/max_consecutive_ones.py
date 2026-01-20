class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        maxOnes = 0
        currentOnes = 0
        
        for i in nums:
            if i != 1:
                if maxOnes < currentOnes:
                    maxOnes = currentOnes
                currentOnes = 0
            else:
                currentOnes += 1
        
        if maxOnes < currentOnes:
            maxOnes = currentOnes
                
        return maxOnes
