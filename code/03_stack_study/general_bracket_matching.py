#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/13 17:03
# @Author  : duanhaobin
# @File    : general_bracket_matching.py
# @Software: PyCharm
# @Desc    : 用栈来实现通用括号匹配算法


"""
    (),{},[], (({}[]))......
    以上括号都是想匹配的,左括号和右括号是想对应的,不仅开闭相对应,类型也必须相对应
    错误的示范:
        ()), ({)), {[]) 这类都是错误的,不符合括号匹配原则
    算法构思:
        从左到右扫描括号,最新打开的左括号,应该匹配最先遇到的右括号
        这样,第一个左括号(最先打开),就应该匹配最后一个右括号(最后遇到)
        这种次序反转的识别,正好符合栈的特性
"""

from my_stack import Stack


def parChecker(symbol_string):
    """
        先写一个只有'()'类型的算法
    @symbol_string: 要匹配的括号字符串
    @return: 返回匹配结果,bool类型
    """
    stack = Stack()
    index = 0
    balanced_flag = True  # 两边括号数量平衡标识

    while index < len(symbol_string) and balanced_flag:
        symbol = symbol_string[index]
        if symbol == '(':
            stack.push(symbol)
        else:
            if stack.isEmpty():
                balanced_flag = False
            else:
                stack.pop()
        index += 1
    if balanced_flag and stack.isEmpty():
        balanced_flag = True
    else:
        balanced_flag = False
    return balanced_flag


def parCheckerMore(symbol_string):
    """
        通用括号匹配,支持 (, [, { 这三种
    @symbol_string: 要匹配的括号字符串
    @return: 返回匹配结果,bool类型
    """
    stack = Stack()
    index = 0
    balanced_flag = True  # 两边括号数量平衡标识

    while index < len(symbol_string) and balanced_flag:
        symbol = symbol_string[index]
        # 判断左边括号类型需要调整
        if symbol in '([{':
            stack.push(symbol)
        else:
            if stack.isEmpty():
                balanced_flag = False
            else:
                top = stack.pop()
                # 栈中移除的括号要和进入栈的括号匹配类型
                if not matches(top, symbol):
                    balanced_flag = False
        index += 1
    if balanced_flag and stack.isEmpty():
        balanced_flag = True
    else:
        balanced_flag = False
    return balanced_flag


def matches(open, close):
    """
        判断左右括号类型是否匹配,如果下标值相同,则匹配.反之,则未匹配
    @open: 左括号
    @close: 右括号
    @retur: 匹配结果,bool值
    """
    opens = '([{'
    closes = ')]}'
    return opens.index(open) == closes.index(close)


if __name__ == '__main__':
    test_symbol = '()()()()()'
    result = parChecker(test_symbol)
    print('测试结果:', result)

    test_symbol_more = '()(){}()({]]]})([])'
    result = parCheckerMore(test_symbol_more)
    print('测试结果:', result)
