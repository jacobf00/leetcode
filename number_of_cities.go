package leetcode

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    visited := make([]bool, n)
    provinces := 0
    
    var dfs func(city int)
    dfs = func(city int) {
        visited[city] = true
        for neighbor := 0; neighbor < n; neighbor++ {
            if isConnected[city][neighbor] == 1 && !visited[neighbor] {
                dfs(neighbor)
            }
        }
    }
    
    for city := 0; city < n; city++ {
        if !visited[city] {
            provinces++
            dfs(city)
        }
    }
    
    return provinces
}
