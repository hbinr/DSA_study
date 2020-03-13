#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/13 15:59
# @Author  : duanhaobin
# @File    : my_stack.py
# @Software: PyCharm
# @Desc    : 设计自定义栈类,模拟入栈,出栈,判空,判满及改变栈大小等操作


class Stack:
    """ 实现自定义栈,以list末尾为栈顶"""

    # 构造方法
    def __init__(self, maxlen=10):
        self._content = []
        self._size = maxlen
        self._current = 0

    # 析构方法,释放列表控件
    def __del__(self):
        del self._content

    # 清空栈中元素
    def clear(self):
        self._content = []
        self._size = 0

    # 测试栈是否为空
    def isEmpty(self):
        return not self._content

    # 修改栈的大小
    def setSize(self, size):
        # 新的栈大小必须大于已有元素数量
        if size < self._current:
            print(f'new size must >={str(self._current)}')
            return
        self._size = size

    # 测试栈是否已满
    def isFull(self):
        return self._current == self._size

    # 入栈
    def push(self, v):
        # 栈未满的情况下才可入栈
        if self._current < self._size:
            # 列表尾部追加元素
            self._content.append(v)
            # 栈中元素+1
            self._current += 1
        else:
            print('Stack is full!')

    # 出栈,删除栈顶元素
    def pop(self):
        # 栈为空的时候无法出栈
        if self._content:
            # 栈中元素-1
            self._current -= 1
            # 弹出并返回列表尾部元素
            return self._content.pop()
        else:
            print("Stack is empty!")

    # 返回栈顶元素,不会删除
    def peek(self):
        # 栈为空的时候无法出栈
        if self._content:
            # 弹出并返回列表尾部元素
            return self._content[self._current - 1]
        else:
            print("Stack is empty!")

    def __str__(self):
        return f'Stack({str(self._content)},maxlen={str(self._size)})'

    # 复用__str__方法
    __repr__ = __str__