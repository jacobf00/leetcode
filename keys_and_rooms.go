package main

// canVisitAllRooms uses Depth-First Search (DFS) to determine if all rooms can be visited
// DFS explores by going as deep as possible along each path before backtracking:
// 1. Start at room 0
// 2. When we find a key, immediately explore that room before trying other keys
// 3. If we reach a room with no new keys, backtrack and try other paths
// 4. Continue until we've tried all possible paths
func canVisitAllRooms(rooms [][]int) bool {
    n := len(rooms)
    visited := make([]bool, n) // Track visited rooms
    
    // DFS recursively explores rooms:
    // - Marks current room as visited
    // - For each key, immediately explores that room if not visited
    // - Naturally backtracks when a path is exhausted
    var dfs func(room int)
    dfs = func(room int) {
        visited[room] = true // Mark current room as visited
        // Try all keys in current room
        for _, key := range rooms[room] {
            if !visited[key] {
                dfs(key)
            }
        }
    }
    
    dfs(0) // Start from room 0
    
    // Check if all rooms were visited
    for _, v := range visited {
        if !v {
            return false
        }
    }
    return true
}
