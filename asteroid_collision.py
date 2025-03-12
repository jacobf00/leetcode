from typing import List

# def asteroidCollision(asteroids: List[int]) -> List[int]:
#     st = []
#     for i, a in enumerate(asteroids):
#         lastSign = -1 if asteroids[i-1] < 0 else 1
#         currentSign = -1 if a < 0 else 1
#         if len(st) == 0:
#             st.append(a)
#         elif lastSign != currentSign:
#             if abs(a) > abs(asteroids[i-1]):
#                 st.pop()
#                 st.append(a)
#             elif abs(a) == abs(asteroids[i-1]):
#                 st.pop()
#         else:
#             st.append(a)
#     return st


def asteroidCollision(asteroids: List[int]) -> List[int]:
    res = []
    for a in asteroids:
        while res and a < 0 < res[-1]:
            if -a > res[-1]:
                res.pop()
                continue
            elif -a == res[-1]:
                res.pop()
            break
        else:
            res.append(a)
    
    return res

if __name__ == "__main__":
    # asteroids = [5,10,-5]
    asteroids = [10,2,-5]
    result = asteroidCollision(asteroids)
    print(result)