#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/11 15:35
# @Author  : duanhaobin
# @File    : 01_anagram_solution.py
# @Software: PyCharm
# @Desc    : 回文词判断练习


"""
    相同字母组成,但是顺序不同的单词,这里练习小写字母,如:
    heart和earth,python和typhon
"""

from collections import Counter


def anagram_solution(s1, s2):
    """
        首先计算每一个字符在字符串中出现的次数。
        由于共有 26 个可能的字符，我们可以利用有 26 个计数器的列表，每个计数器对应一个字符。
        每当我们看到一个字符，就在相对应的计数器上加一.
        最终，如果这两个计数器列表相同，则这两个字符串是变位词
        时间复杂度为:O(n)
    """
    count1 = {}
    count2 = {}
    for i in range(len(s1)):
        pos = ord(s1[i]) - ord('a')  # ord返回unicode编码
        count1[pos] = pos + 1

    for j in range(len(s2)):
        pos = ord(s2[j]) - ord('a')
        count2[pos] = pos + 1
    print('count1:', count1)
    print('count2:', count2)

    m = 0
    still_ok = True
    while m < 26 and still_ok:
        if count1 == count2:
            m += 1
        else:
            still_ok = False

    return still_ok


def anagram_solution_new(s1, s2):
    """
        Python中提供了Counter函数,直接就可以使用
    """
    return Counter(s1) == Counter(s2)


if __name__ == '__main__':
    s1 = 'python'
    s2 = 'typhon'
    print("s1和s2是否为回文词:", anagram_solution(s1, s2))

    print(anagram_solution_new(s1, s2))
