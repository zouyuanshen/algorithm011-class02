### 做题的一些问题
- 递归容易忘了在terminal 语句return 
- go一些指针引用问题需要注意

### 泛型递归
- 模版
> recursion terminator --> process logic --> driil down --> reverse level status

### 分治
> 把一个大问题分解成很多小的问题，而且各个小问题的解可以组合成大问题的解

- 分治也是本质上一种递归，是把每个drill down merge在一起

### 回溯
- 回溯也是本质上一种递归
- 把问题的解答一个个去尝试，有种回过头去重新走一遍的感觉
- 要注意在递归的时候去处理递归数据和level status
