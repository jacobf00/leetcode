package main

func orangesRotting(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    
    rows, cols := len(grid), len(grid[0])
    queue := make([][2]int, 0)
    fresh := 0
    
    // Find all rotten oranges and count fresh ones
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                fresh++
            }
        }
    }
    
    // If no fresh oranges, return 0
    if fresh == 0 {
        return 0
    }
    
    // BFS directions
    dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    minutes := 0
    
    // BFS process
    for len(queue) > 0 && fresh > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            curr := queue[0]
            queue = queue[1:]
            
            // Try all 4 directions
            for _, dir := range dirs {
                newRow := curr[0] + dir[0]
                newCol := curr[1] + dir[1]
                
                // Check bounds and if orange is fresh
                if newRow >= 0 && newRow < rows && 
                   newCol >= 0 && newCol < cols && 
                   grid[newRow][newCol] == 1 {
                    grid[newRow][newCol] = 2
                    fresh--
                    queue = append(queue, [2]int{newRow, newCol})
                }
            }
        }
        minutes++
    }
    
    // Check if any fresh oranges remain
    if fresh > 0 {
        return -1
    }
    return minutes
}