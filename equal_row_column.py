from typing import List

def equalPairs(grid: List[List[int]]) -> int:
    n = len(grid)
    num_pairs = 0
    for row in grid:
        for i in range(n):
            col = [roww[i] for roww in grid]
            if hash(tuple(row)) == hash(tuple(col)):
                num_pairs += 1
    return num_pairs
            



if __name__ == "__main__":
    grid = [[3,2,1],[1,7,6],[2,7,7]]
    pairs = equalPairs(grid)
    print(f"There are {pairs} pairs")