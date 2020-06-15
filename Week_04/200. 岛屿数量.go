package Week_04

func numIslands(grid [][]byte) int {
	if len(grid) < 1 || len(grid[0]) < 1{
		return 0
	}

	h := len(grid)
	w := len(grid[0])

	count := 0

	for i:= 0; i<w; i++{
		for j:=0; j <h; j++{
			if grid[j][i] == '1'{
				count++
				dfs(grid, j, i)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, line, col int){
	h := len(grid)
	w := len(grid[0])

	grid[line][col] = '0'

	if line-1 >= 0 && grid[line-1][col] == '1'{
		dfs(grid, line-1, col)
	}

	if line+1 < h && grid[line+1][col] == '1'{
		dfs(grid, line+1, col)
	}

	if col-1>=0 && grid[line][col-1] == '1'{
		dfs(grid, line, col-1)
	}

	if col+1 < w && grid[line][col+1] == '1'{
		dfs(grid, line, col+1)
	}
}
