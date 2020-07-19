学习笔记
### 深度优先搜索和广度优先搜索
- dfs
```go
func dfs(root *TreeNode, level int, result [][]int)  {
	if root == nil {
		return
	}
	result[level] =  append(result[level], root.Val)
	if root.Left != nil{
		dfs(root.Left, level+1, result)
	}
	if root.Right != nil{
		dfs(root.Right, level+1, result)
	}
}

```
- bfs

```go

func bfs(root *TreeNode )  []int{
    result := make([]int, 0) 
	queue := []*TreeNode{root}

	if len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if root.Left != nil{
			queue = append(queue, root.Left)
		}

		if node.Right != nil{
			queue = append(queue, root.Right)
		}
	}
	return result
}
```
###贪心算法
> 一种在每一步选择中都采取在当前状态最好或最优的选择，从而希望导致结果是全局最好或最优的算法

### 二分法
- 模板
```go
 func binarySearch(nums []int, target int)  int {
 	left, right, mid := 0, len(nums) - 1, math.MaxInt32
 	for left <= right{
 		mid = (left + right) / 2
 		if nums[mid] == target{
 			return mid
 		}else if nums[mid] < target{
 			left = mid + 1
 		}else {
 			right = mid -1
 		}
 
 
 	}
 
 	return -1
 }
```

