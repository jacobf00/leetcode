package leetcode

func minReorder(n int, connections [][]int) int {
    // Build adjacency list with both original and reverse edges
    adj := make([][]pair, n)
    for i := 0; i < n; i++ {
        adj[i] = make([]pair, 0)
    }
    
    // Store original edges with direction flag
    for _, conn := range connections {
        from, to := conn[0], conn[1]
        adj[from] = append(adj[from], pair{to, 1})    // Original edge (costs 1 to reverse)
        adj[to] = append(adj[to], pair{from, 0})      // Reverse edge (costs 0 to use)
    }
    
    visited := make([]bool, n)
    
    var dfs func(city int) int
    dfs = func(city int) int {
        changes := 0
        visited[city] = true
        
        for _, next := range adj[city] {
            if !visited[next.city] {
                changes += next.cost + dfs(next.city)
            }
        }
        
        return changes
    }
    
    return dfs(0)
}

type pair struct {
    city int
    cost int
}
