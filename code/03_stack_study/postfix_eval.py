#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/13 22:12
# @Author  : duanhaobin
# @File    : postfix_eval.py
# @Software: PyCharm
# @Desc    : 后缀表达式求值


"""
    从左到右扫描后缀表达式,先得到是操作数,后得到操作符
    所以要暂存操作数,在碰到操作符的时候,再将暂存的两个操作数进行实际的计算
    操作符只作用与离它最近的两个数,栈顶的前两个数,先弹出的是右操作数,后弹出的是左操作数

    但是在扫描操作符的时候,要注意'-/'  因为4-5和5-4是不一样的

    因此为了后续的计算,需要把中间计算的结果先压入栈顶,然后继续扫描后面的符号

    当所有的操作符都处理完毕,栈中只留一个操作数,就是表达式的值

    注意:数字必须是0-9
"""
from my_stack import Stack

def postfixEval(postfixExpr):
    """
        后缀表达式求值
    @postfixExpr: 后缀表达式
    @return: 计算结果
    """
    token_lst = postfixExpr.split()
    operand_stack = Stack()

    for token in token_lst:
        # 如果是操作数,则压入栈顶
        if token in '123456789':
            operand_stack.push(int(token))

        # 否则取操作数开始计算
        else:
            # 栈的特性,后进先出
            operand2 = operand_stack.pop()
            operand1 = operand_stack.pop()
            result = math_eval(token, operand1, operand2)
            operand_stack.push(result)

    return operand_stack.pop()


def math_eval(op, op_num1, op_num2):
    if op == '*':
        return op_num1 * op_num2
    elif op == '/':
        return op_num1 / op_num2
    elif op == '+':
        return op_num1 + op_num2
    elif op == '-':
        return op_num1 - op_num2
    else:
        return '操作符错误,后缀表达式是否正确!'


if __name__ == '__main__':
    from infix_to_postfix import infixToPostfix

    postfixExpr = infixToPostfix('8 - 1 - 9')
    print('后缀表达式:', postfixExpr)

    result = postfixEval(postfixExpr)
    print('计算结果:', result)