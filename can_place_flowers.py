from typing import List
import math

def canPlaceFlowers(flowerbed: List[int], n: int) -> bool:
    numFlowersCanPlace = 0
    lastValue = 1
    indexToNonFlowers = {}
    flowerbed = [0] + flowerbed + [0]
    for i, isFlower in enumerate(flowerbed):
        if isFlower == 0 and lastValue != 0:
            try:
                nextFlower = flowerbed.index(1, i)
                indexToNonFlowers[i] = len(flowerbed[i:nextFlower])
            except ValueError:
                indexToNonFlowers[i] = len(flowerbed[i:])
        lastValue = isFlower

    for index, numNonFlowers in indexToNonFlowers.items():
        if numNonFlowers >= 3:
            numFlowersCanPlace += 1 + math.trunc((numNonFlowers - 3)/2)

    if n <= numFlowersCanPlace:
        return True
    else:
        return False   
            
 
# find sequences of 0s
flowerbed = [1,0,0,0,1,0,0]
n = 2
# 1:0
# 2:0
# 3:1
# 4:1
# 5:2
# 6:2
# 7:3
# 8:3
# 9:4
print(canPlaceFlowers(flowerbed, n))