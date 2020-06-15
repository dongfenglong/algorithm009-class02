# 学习笔记

## 第9课 深度优先搜索和广度优先搜索

1. 搜索-遍历

   ① 每个节点都要访问一次

   ② 每个节点仅访问一次

   ③ 对于节点的访问顺序不限

   	- 深度优先搜索
   	- 广度优先搜索

2. 深度优先搜索示例代码

   ```python
   def dfs(node):
       if node in visited:
           # already visited
           return
       visited.add(node)
       
       # process current node
       # ... # logic here
       dfs(node.left)
       dfs(node.right)
   ```

   

   #递归写法

   ```python
   visited = set()
   def dfs(node, visited):
       if node in visited: # terminator
           # already visited
           return
       
       visited.add(node)
       # process current node here.
       ...
       for next_node in node.children():
           if not next_node in visited:
               dfs(next_node, visited)
   ```

   

   #非递归写法

   ```python
   def dfs(self, tree):
       if tree.root is None:
           return []
       
       visited, stack = [], [tree.root]
       
       while stack:
           node = stack.pop()
           visited.add(node)
           
           process(node)
           nodes = generate_related_nodes(node)
           stack.push(nodes)
           
       # other processing work
       ...
   ```

   

3. 广度优先搜索示例代码

   ```python
   def bfs(graph, start, end):
       queue = []
       queue.append([start])
       visited.add(start)
       
       while queue:
           node = queue.pop()
           visited.add(node)
           
           process(node)
           nodes = generate_related_nodes(node)
           queue.push(nodes)
           
       # other processing work
       ...
   ```



## 第10课 贪心算法

1. 贪心算法是一种在每一步选择中都采取在当前状态下最好或最优（即最有利）的选择，从而希望导致结果是全局最好或最优的算法。

2. 贪心算法与动态规划的不同，在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。

3. 贪心、回溯、动态规划

   贪心：当下做局部最优判断

   回溯：能够回退

   动态规划：最优判断+回退

4. 贪心算法可以解决一些最优化问题，如：求图中的最小生成树、求哈夫曼编码等。然而对于工程和生活中的问题，贪心法一般不能得到我们所要求的答案。

5. 一旦一个问题可以通过贪心法来解决，那么贪心法一般是解决这个问题的最好办法。由于贪心法的高效性以及其所求得的答案比较接近最优结果，贪心法也可以用作辅助算法或者直接解决一些要求结果不是特别精确的问题。

6. 使用贪心算法的场景

   简单地说，问题能够分解成子问题来解决，子问题的最优解能递推到最终问题的最优解。这种子问题最优解称为最优子结构。

7. 贪心算法与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。



## 第11课 二分查找

1. 二分查找的前提

   - 目标函数单调性（单调递增或单调递减）
   - 存在上下界（bounded）
   - 能够通过索引访问（index accessible）

2. 代码模板

   ```python
   left, right = 0, len(array)-1
   while left <= right:
       mid = (left+right)/2
       if array[mid] == target:
           # find the target!!!
           break or return result
       elif array[mid] < target:
           left = mid + 1
       else:
           right = mid - 1
   ```

   

3. 

4. 