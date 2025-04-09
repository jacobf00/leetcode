func nearestExit(maze [][]byte, entrance []int) int {
    rows, cols := len(maze), len(maze[0])
    dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    
    // Mark entrance as visited
    queue := [][3]int{{entrance[0], entrance[1], 0}}
    maze[entrance[0]][entrance[1]] = '+'
    
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        row, col, steps := curr[0], curr[1], curr[2]
        
        // Check all four directions
        for _, dir := range dirs {
            newRow, newCol := row+dir[0], col+dir[1]
            
            // Check if position is valid and unvisited
            if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && maze[newRow][newCol] == '.' {
                // If we reached a border and it's not the entrance, we found an exit
                if newRow == 0 || newRow == rows-1 || newCol == 0 || newCol == cols-1 {
                    return steps + 1
                }
                
                // Mark as visited and add to queue
                maze[newRow][newCol] = '+'
                queue = append(queue, [3]int{newRow, newCol, steps + 1})
            }
        }
    }
    
    return -1
}
