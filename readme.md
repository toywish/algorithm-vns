# VNS解决0-1背包问题
## VNS介绍

变邻域搜索算法（VNS）就是一种改进型的局部搜索算法。它利用不同的动作构成的邻域结构进行交替搜索，在集中性和疏散性之间达到很好的平衡。其思想可以概括为“变则通”。

变邻域搜索算法**依赖于以下事实：**

1) 一个邻域结构的局部最优解**不一定**是另一个邻域结构的局部最优解。

2) 全局最优解是所有可能邻域的局部最优解。

变邻域搜索算法主要由以下两个部分组成：

**1) VARIABLE NEIGHBORHOOD DESCENT (VND)**

**2) SHAKING PROCEDURE**

参考文献：[干货 | 变邻域搜索算法(Variable Neighborhood Search,VNS)超详细一看就懂 (qq.com)](https://mp.weixin.qq.com/s?__biz=MzU0NzgyMjgwNg==&mid=2247484621&idx=1&sn=f2e92f44c2306b58034cf158647bc737&chksm=fb49c974cc3e406228737e1a986c73368131bc7f0c0251d82b1e64266220df59134ab0a9def1&scene=21#wechat_redirect)

## 0-1背包问题

> 0-1 背包问题：给定 n 种物品和一个容量为 C 的背包，物品 i 的重量是 w_i，其价值为 v_i。

> 问：应该如何选择装入背包的物品，使得装入背包中的物品的总价值最大？

此代码是[干货 | 变邻域搜索算法解决0-1背包问题(Knapsack Problem)代码实例 (qq.com)](https://mp.weixin.qq.com/s/wedhCjmtHwHTqVroM0Q63Q)的代码复现。

原文使用的是C++，这里使用的是Golang，版本1.18.3。


