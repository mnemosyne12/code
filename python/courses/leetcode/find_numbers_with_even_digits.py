class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        evenCount = 0
        for i in nums:
            if len(str(i)) % 2 == 0:
                evenCount += 1
        return evenCount

class Solution22:
    def findNumbers(self, nums: List[int]) -> int:
        evenCount = 0
        currentNumber = 0
        for i in nums:
            currentNumber = i
            digitCount = 0
            while currentNumber > 0:
                currentNumber //= 10
                digitCount += 1
            if digitCount % 2 == 0:
                evenCount += 1
        return evenCount
