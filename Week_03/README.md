# 学习笔记

## 第7课	泛型递归、树的递归

1. 递归

    计算n!

    n! = 1 * 2 * 3 * ... * n

    ```python
    def Factorial(n)
    	if n <= 1:
            return 1
        return n * Factorial(n-1)
    ```

    

2. 递归代码模板

    ```python
    def recursion(level, param1, param2, ...):
        #recursion terminator	递归终结条件
        if level >MAX_LEVEL:
            process_result
            return
        
        #process logic in current level	处理当前层逻辑
        process(level, data...)
        
        #drill down				下探到下一层
        self.recursion(level+1, p1, ...)
        
        #reverse the current level status if neeeded	必要时清理当前层状态
        
    ```

    

3. 思维要点

    (1) 不要进行人肉递归

    (2) 找到最近最简方法，将其折解成可重复解决的问题（重复子问题）

    (3) 数学归纳法思想



## 第8课	分治、回溯

1. 分治代码模板

    ```python
    def divide_conquer(problem, param1, param2, ...):
        #recursion terminator:
        if problem is None:
            print_result
            return
        
        #prepare data
        data = prepare_data(problem)
        subproblems = split_problem(problem, data)
        
        #conquer subproblems
        subresult1 = self.divide_conquer(subproblems[0], p1, ...)
        subresult2 = self.divide_conquer(subproblems[1], p1, ...)
        subresult3 = self.divide_conquer(subproblems[2], p1, ...)
        ...
        
        #process and generate the final result
        result = process_result(subresult1, subresult2, subresult3, ...)
        
        #revert the current level status
        
    ```

    

2. 回溯

    (1) 采用试错的思想，尝试分步解决一个问题

    (2) 在分布解决问题的过程中，

    ​	 当现有分步答案不能得到有效的、正确的答案时，

    ​	 它将取消上一步甚至是上几步的计算，

    ​	 再通过其他的可能的分步解答再次寻找问题的答案。

    (3) 回溯法通常用最简单的递归方法来实现。



























