#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/13 17:45
# @Author  : duanhaobin
# @File    : 10num_to_2num.py
# @Software: PyCharm
# @Desc    : 利用栈的特性来求解10进制数转成2进制


"""
    先回顾一下,10进制数如何转成2进制数
    通常,我们使用'除以2'的方法来得到2进制数.
    在'除以2'的过程中,得到的余数是从从上到下,而输出2进制数是反着来的,从下到上输出
    所以可以利用栈的特性来解决
"""
from my_stack import Stack


def divideBy2(decNumber):
    """
    10进制数转2进制
    @num: 需要转换的数
    @return: 转换后的结果
    """
    remainder_stack = Stack()

    while decNumber > 0:
        remainder = decNumber % 2
        remainder_stack.push(remainder)
        decNumber //= 2  # 整除

    result = ''
    while not remainder_stack.isEmpty():
        result += str(remainder_stack.pop())

    return result


def baseConvert(decNumber, base):
    """
    10进制数转成任意16进制以下的数

    @decNumber: 需要转换的数
    @base: 进制
    @return: 转换后的结果
    """
    digits = '0123456789ABCDEF'  # 定义数范围(最大16进制)

    remainder_stack = Stack()

    while decNumber > 0:
        remainder = decNumber % base
        remainder_stack.push(remainder)
        decNumber //= base  # 整除

    result = ''
    while not remainder_stack.isEmpty():
        result += digits[remainder_stack.pop()]

    return result



def Dec2Bin(num):
    result = ''
    if num:
        result = Dec2Bin(num // 2)
        print(result)
        return result + str(num % 2)
    else:
        return result

if __name__ == '__main__':
    result = divideBy2(35)
    print('10进制数转2进制,结果:', result)

    # Python内置函数:10进制转2进制 bin(),返回结果开头带0b,表示二进制数
    result_bin = bin(35)
    print('bin内置函数10进制数转2进制,结果:', result_bin)  # 结果一致

    result = baseConvert(35, 2)
    print('35转2进制:', result)

    result = baseConvert(1000, 8)
    print('1000转8进制:', result)

    print(Dec2Bin(10))