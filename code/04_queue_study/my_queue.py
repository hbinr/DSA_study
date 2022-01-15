#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/14 15:23
# @Author  : duanhaobin
# @File    : my_queue.py
# @Software: PyCharm
# @Desc    : 自定义普通队列,模拟入队,出队,判空,判满,改变大小等操作


class Queue:
    """ 实现自定义队列,以list末尾为队尾"""

    # 构造方法
    def __init__(self, maxlen=10):
        self._content = []
        self._size = maxlen
        self._current = 0

    # 析构方法,释放列表控件
    def __del__(self):
        del self._content

    # 清空队列中元素
    def clear(self):
        self._content = []
        self._size = 0

    # 设置队列大小
    def setSize(self, size):
        self._size = size

    # 测试队列是否为空
    def isEmpty(self):
        return not self._content

    # 测试队列是否已满
    def isFull(self):
        return self._current == self._size

    # 入队列,尾部加入元素
    def enqueue(self, v):
        # 队列未满的情况下才可入队
        if self._current < self._size:
            self._content.append(v)
            self._current += 1
        else:
            print("The queue is full and cannot be queued")

    # 出队列,删除队首元素
    def popqueue(self):
        # 队列为空的时候无法出队
        if self._content:
            # 队列元素-1
            self._current -= 1
            return self._content.pop(0)
        else:
            print("Can't leave the queue when the queue is empty")

    def __str__(self):
        return f'Queue({str(self._content)},maxlen={str(self._size)})'

    # 复用__str__方法
    __repr__ = __str__
