from typing import List
import numpy as np
import math

def productExceptSelf(nums: List[int]) -> List[int]:
    # prods = [0]*len(nums)
    # for i in range(len(nums)):
    #     prods[i] = (int(np.concatenate((np.array(nums[0:i]), np.array(nums[i+1:]))).prod()))
    # return prods 
    # prefixes = [0]*(math.ceil(len(nums)/2))
    # suffixes = [0]*(math.trunc(len(nums)/2))
    # for i in range(len(prefixes)):
    #     prefixes[i] = np.array(nums[:i+1]).prod()
    # for i in reversed(range(len(suffixes))):
    #     suffixes[i] = np.array(nums[i:]).prod()
    # pre_suf = prefixes + suffixes    
    # prods = [0]*len(nums)
    # for i in range(len(nums))[1:-2]:
    #     prods[i] = int(pre_suf[i-1] * pre_suf[i+1]) 
    # prods[0] = nums[0]
    # prods[-1] = nums[-1] 
    prefixes = [0]*(len(nums))
    suffixes = [0]*(len(nums))
    prods = [0]*len(nums)
    
    for i in range(len(nums)):
        prefixes[i] = int(np.array(nums[:i]).prod())
        suffixes[i] = int(np.array(nums[i+1:]).prod())
        prods[i] = int(prefixes[i] * suffixes[i]) 

    return prods 

# math.ceil((len(nums)/2)-1)
# nums = [-1,1,0,-3,3]
nums = [1,2,3,4]
ans = productExceptSelf(nums)
print(ans)
print(len(ans))
# print(math.ceil((4/2) - 1))