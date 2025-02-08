from typing import List

class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        mostCandies = max(candies)
        hasMost = [] 
        for numCandies in candies:
            if (numCandies + extraCandies) >= mostCandies:
                hasMost.append(True)
            else:
                hasMost.append(False)
        return hasMost        

candies = [2,3,5,1,3]
extraCandies = 3

sol = Solution.kidsWithCandies(..., candies, extraCandies)
print(f"The solution is: {sol}")