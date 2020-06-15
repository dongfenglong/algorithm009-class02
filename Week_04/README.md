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