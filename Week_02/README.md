# 学习笔记

## 第5课	哈希表、映射、集合

​	哈希表，也叫散列表，是根据关键码值（key、value）而直接进行访问的数据结构。

​	它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。

​	这个函数叫作映射函数，也叫散列函数（Hash Function）。

​	存放记录的数组叫做哈希表，或散列表。

​	Hash Collisions（哈希碰撞），导致数组加挂链表。



​	Hash Table时间复杂度

|      | 查找 | 插入 | 删除 |
| :--: | :--: | :--: | :--: |
| 平均 | O(1) | O(1) | O(1) |
| 最坏 | O(n) | O(n) | O(n) |



​	工程应用：

​		Map：key-value，key不重复。

​		Set：  不重复元素的集合。



## 第6课（上）	树、二叉树、二叉搜索树

1. Linked List是特殊化的Tree
    Tree是特殊化的图

2. 二叉树遍历

    前序遍历	Pre-Order：	<u>根</u>左右

    中序遍历	In-Order：	  左<u>根</u>右

    后序遍历	Post-Order：  左右<u>根</u>

3. Binary Search Tree

    二叉搜索树，是指一棵空树或者具有下列性质的二叉树：

    ​	左子树 <u>所有结点</u> 的值均小于它的根结点的值；

    ​	右子树 <u>所有结点</u> 的值均大于它的根结点的值；

    ​	以此类推，左右子树也分别为二叉查找树。

    二叉搜索树的中序遍历是升序遍历。

    常见操作：

    ​	查询

    ​	插入新结点

    ​	删除







## 第6课（下）	堆、二叉堆、图

1. Heap

    ①堆，一种数据结构，可以迅速找到一堆数中的最大值或最小值。

    

    ②将根结点最大的堆叫做大顶堆、大根堆；

    ​	将根结点最小的堆叫做小顶堆、小根堆；

    ​	常见的堆有二叉堆、斐波那契堆等。

    

    ③假设是大顶堆，则常见的操作（API）：

    ​	find-max：	O(1)

    ​	delete-max：O(logN)

    ​	insert（create）：O(logN)或O(1)

    

    ④二叉堆性质

    ​	通过完全二叉树来实现（注意：跟二叉搜索树没有关系）。

    ​	完全二叉树，除了最下层结点，剩下的所有结点都是丰满的。

    

    ​	二叉堆（大顶）满足以下性质：

    ​		性质一：是一棵完全树

    ​		性质二：树中任意结点的值总是 >= 其子结点的值

    

    ⑤实现细节

    ​	二叉堆一般都是通过“数组”来实现

    ​	假设第一个元素在数组中的索引为 0 的话，则父结点和子结点的关系如下：

    ​		01.  索引为 i 的结点的左孩子索引是（2i+1）

    ​		02. 索引为 i 的结点的右孩子索引是（2i+2）

    ​		03. 索引为 i 的结点的父节点的索引是floor($\frac {n-1}{2}$)

    

    ​	Insert 插入操作：

    ​		01.  新元素一律插入到堆的尾部

    ​		02. 依次向上调整整个堆的结构，一直到根即可	HeapifyUp

    

    ​	Delete Max 操作：	O(logN)

    ​		01. 将堆尾元素替换到顶部，即堆顶元素被替代删除掉

    ​		02. 依次从根部向下调整整个堆的结构，一直到堆尾即可

    

    ​	注意：

    ​		二叉堆是堆（优先队列 Priority Queue）的一种常见且简单的实现，但是并不是最优的实现。

    

    

2. 图的实现和特性

    ①什么是图

    ​	由点和边所构成的图形就是图。

    ②图的属性

    ​	Graph(V, E)

    ​	V - vertex: 点

    ​		度：入度、出度

    ​		点与点之间：连通与否

    ​	E - edge:   边

    ​		有向和无向，对应入度、出度，无向边只有度

    ​		权重：边长

    ③分类

    ​	无向无权图

    ​	有向无权图

    ​	无向有权图

    ​	有向有权图

    ④表示图的方式

    ​	邻接矩阵、

    ​	邻接表

    ⑤基于图的常见算法——DFS

    ```python
    visited = set()
    def dfs(node, visited):
        if node in visited: #已经访问过
            return
        visited.add(node)	#和树的DFS的最大区别
        #process current node here
        ...
        
        for next_node in node.children():
            if not next_node in visited:
                dfs(next_node, visited)
    ```

    

    ⑥基于图的常见算法——BFS

    ```python
    def bfs(graph, start, end):
        queue = []
        queue.append([start])
        visited = set()		#和树的BFS的最大区别
        
        while queue:
            node = queue.pop()
            visited.add(node)
            process(node)
            nodes = generate_related_nodes(node)
            queue.push(nodes)
    ```



​		⑦图的高级算法

​			连通图个数

​			拓扑排序

​			最短路径

​			最小生成树

































