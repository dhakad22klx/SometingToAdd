//Leetcode link : https://leetcode.com/problems/trapping-rain-water-ii/description/

package priorityQueue

import "container/heap"


type Cell struct {
    height, row, col int
}

// MinHeap to prioritize the cells based on their height
type MinHeap []Cell

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
    return h[i].height < h[j].height // Min heap by height
}
func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Cell))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func trapRainWater(heightMap [][]int) int {
    if len(heightMap) == 0 || len(heightMap[0]) == 0 {
        return 0
    }

    vol := 0
    M, N := len(heightMap), len(heightMap[0])

    // Directions for moving in 4 cardinal directions
    directions := [][2]int{
        {0, 1}, {0, -1}, {1, 0}, {-1, 0},
    }

    // Min-heap and visited matrix
    minHeap := &MinHeap{}
    visited := make([][]bool, M)
    for i := range visited {
        visited[i] = make([]bool, N)
    }

    // Add all boundary cells to the min-heap
    for i := 0; i < N; i++ {
        heap.Push(minHeap, Cell{heightMap[0][i], 0, i})
        heap.Push(minHeap, Cell{heightMap[M-1][i], M - 1, i})
        visited[0][i] = true
        visited[M-1][i] = true
    }

    for i := 0; i < M; i++ {
        heap.Push(minHeap, Cell{heightMap[i][0], i, 0})
        heap.Push(minHeap, Cell{heightMap[i][N-1], i, N - 1})
        visited[i][0] = true
        visited[i][N-1] = true
    }

    // Perform the BFS-like processing
    for minHeap.Len() > 0 {
        curr := heap.Pop(minHeap).(Cell)
        height, row, col := curr.height, curr.row, curr.col

        // Explore the neighboring cells
        for _, dir := range directions {
            r, c := row+dir[0], col+dir[1]
            if r >= 0 && r < M && c >= 0 && c < N && !visited[r][c] {
                visited[r][c] = true

                // If the neighboring cell is lower, water can be trapped
                if heightMap[r][c] < height {
                    vol += height - heightMap[r][c]
                }

                // Push the neighboring cell to the min-heap with updated height
                heap.Push(minHeap, Cell{max(height, heightMap[r][c]), r, c})
            }
        }
    }

    return vol
}

// Helper function to get the maximum value
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}