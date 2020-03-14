#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/14 20:25
# @Author  : duanhaobin
# @File    : joseph_ring.py
# @Software: PyCharm
# @Desc    : 解决约瑟夫环问题


"""
    约瑟夫环（约瑟夫问题）是一个数学的应用问题：
    编号为 1-N 的 N 个士兵围坐在一起形成一个圆圈，从编号为 1 的士兵开始依次报数（1，2，3…这样依次报），
    数到 m 的 士兵会被杀死出列，之后的士兵再从 1 开始报数。直到最后剩下一士兵，求这个士兵的编号。

    因为约瑟夫环的规则，所以在进行人员出列的时候，剩下的人员的排序是个问题：
    1、如果是静态的重排，那么就需要列尾最后一个人报完后从列头从新报数，是否可以将报完数的人排到列尾。
    2、如果不进行重新排序，是否可以通过计数来标记出列的人（或存活的人）
"""


def queueSolution(n, m):
    """
        解决方案一:
        可以使用队列的思想,每次从队首(序号为1)开始报数,凡是没有报到m的人出列,并直接排到队尾,然后接着报数
        直到报到m的人出现,此时直接出列. 然后再从头(序号为1)开始报道,一直循环,直到最后只剩下1个人
        这样就动态形成了一个环

    @param n: 总人数n
    @param m: 报到m的人出列
    @return: 被'淘汰'的人列表
    """
    eliminate_lst = []
    people_lst = list(range(1, n + 1))

    # 最后只留1个人
    while len(people_lst) > 1:
        i = 1  # 如果报到m出列,要从1开始报数
        while i < m:
            people_lst.append(people_lst.pop(0))  # 没有报到m的人出列,并排到队尾
            i += 1
        eliminate_lst.append(people_lst.pop(0))  # 报到m的人出列后单独保存到被淘汰列表

    return eliminate_lst


def iterationSolution(n, m):
    """
    参考:https://blog.csdn.net/dreamispossible/article/details/96360338?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task
    迭代思路:
    每次我们删除了某一个士兵之后，我们就对这些士兵重新编号，然后我们的难点就是找出删除前和删除后士兵编号的映射关系。
    定义递归函数 f(n，m) 的返回结果是存活士兵的编号，显然当 n = 1 时，f(n, m) = 1。

    假如我们能够找出 f(n，m) 和 f(n-1，m) 之间的关系的话，我们就可以用递归的方式来解决了。我们假设人员数为 n, 报数到 m 的人就自杀。则刚开始的编号为:
    1 2 3 4 ... m-2 m-1 m m+1 m+2 ... n

    进行了一次删除之后，删除了编号为 m 的节点。删除之后，就只剩下 n - 1 个节点了，删除前和删除之后的编号转换关系为：
    删除前 — 删除后
    1       ...
    2       ...
    3       ...
    ...     ...
    m-2     n-2
    m-1     n-1
    m   ---- 无(因为编号被删除了)
    m+1 ---- 1
    m+2 ---- 2
    m+3 ---- 3
    n   ----...
    新的环中只有 n - 1 个节点。
    观察发现:删除前编号为 m + 1, m + 2, m + 3 的节点成了删除后编号为 1， 2， 3 的节点。

    假设 old 为删除之前的节点编号， new 为删除了一个节点之后的编号，则 old 与 new 之间的关系为:
        old = (new + m - 1) % n + 1。
    这样，我们就得出 f(n, m) 与 f(n - 1, m)之间的关系了,又因为:
        f(1, m) = 1
        m + 1 = 1
    带入公式验证:
        m + 1 = (1 + m - 1)%1 + 1
    得:
        m + 1 = m + 1
    这样我们就可以用递归得思路了

    @param n: 总人数n
    @param m: 报到m的人出列
    @return: 幸存者
    """
    # if n == 1:
    #     return n
    # else:
    #     return (iterationSolution(n - 1, m) + m - 1) % n + 1

    return n if n == 1 else (iterationSolution(n - 1, m) + m - 1) % n + 1

if __name__ == '__main__':
    # 总共有10个人,报到3的人出列,返回淘汰人序号列表
    result = queueSolution(10, 3)
    print(result)

    # 总共有10个人,报到3的人出列,返回幸存者序号
    result = iterationSolution(10, 3)
    print(result)