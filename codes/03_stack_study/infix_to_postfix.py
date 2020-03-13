#!/usr/bin/env python 
# -*- coding: utf-8 -*-
# @Time    : 2020/3/13 21:10
# @Author  : duanhaobin
# @File    : infix_to_postfix.py
# @Software: PyCharm
# @Desc    : infix to postfix,中缀表达式转后缀表达式算法实现


"""
    中缀表达式:  A + B * C
    对应的后缀表达式: A B C * +

    操作数ABC的顺序没有改变,操作符(*+)的出现顺序,在后缀表达式中反转了
    由于*的优先级比+高,所以后缀表达式中操作符的出现顺序于运算次序一致

    因此,操作符要比操作数晚输出,所以在扫描到对应的第二个操作数之前,需要把操作符保存起来

    而这些暂存的操作符,由于优先级的规则,还可能要反转次序输出.如A+B*C
    由于这种反转特性,我们可以考虑使用栈的特性来保存未处理的操作符

    带括号的表达式:
    中缀表达式:  (A+B)*C
    对应的后缀表达式: AB+C*
    这里+的输出要比*早,这是因为()的存在,使+的优先级高于*了.

    对于带括号的表达式,后缀表达式中的操作符应该出现在左括号对应的右括号位置(操作符替换右括号即可),然后删除左括号
    所以,遇到左括号,要标记下,其后对应的操作符优先级就相当于提升了,一旦扫描到对应的右括号,就可马上输出这个操作符

    算法约定:
    操作符包括:+-*/()
    操作数:A-Z,0-9

    创建空栈opstack用来存操作符,空表用来保存后缀表达式
"""
from my_stack import Stack


def infixToPostfix(infixExpr):
    """
    中缀表达式转后缀表达式算法实现
    @infixExpr: 中缀表达式,格式如:A + B * C,必须含空格,且操作数只包含A-Z 0-9其中的值
    @return: 转换成功的后缀表达式
    """
    # 记录操作符优先级
    prec = {}
    prec["*"] = 3
    prec["/"] = 3
    prec["+"] = 2
    prec["-"] = 2
    prec["("] = 1

    op_stack = Stack()
    possible_str = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
    post_fix_lst = []
    token_lst = infixExpr.split()
    for token in token_lst:
        # 如果单词是 操作数, 则直接添加到后缀表达式列表的末尾
        if token in possible_str:
            post_fix_lst.append(token)

        # 如果单词是左'(', 则压入栈顶,标记下
        elif token == '(':
            op_stack.push(token)

        # 如果单词是右')',则反复弹出栈顶操作符,加入到输出列表末尾,直到遇到左括号
        elif token == ')':
            top_token = op_stack.pop()
            while top_token != '(':
                post_fix_lst.append(top_token)
                top_token = op_stack.pop()
        # 如果单词是操作符, 则压入栈顶,
        # 但在压之前, 要比较其与栈顶操作符的优先级
        # 如果栈顶操作符的优先级高于它, 就要反复弹出栈顶操作符, 加入到输出列表末尾
        # 直到栈顶操作符的优先级低于它

        else:
            # 注意条件prec[op_stack.peek()] >= prec[token], 不是 > ,是 >=,因为+- 或*/是同一级
            while (not op_stack.isEmpty()) and (prec[op_stack.peek()] >= prec[token]):
                post_fix_lst.append(op_stack.pop())
            op_stack.push(token)

    # 扫描结束后,把opstack栈中的所有剩余操作符依次弹出,加入到输出列表末尾
    # 把输出列表再用join()合并成后缀表达式字符串
    while not op_stack.isEmpty():
        post_fix_lst.append(op_stack.pop())

    return " ".join(post_fix_lst)


if __name__ == '__main__':
    infixExpr = "A + B * C"

    result = infixToPostfix(infixExpr)
    print('转换后的结果:', result)

    infixExpr = "( A + B ) * C"
    result = infixToPostfix(infixExpr)
    print('转换后的结果:', result)

    infixExpr = "9 - 9 - 8"
    result = infixToPostfix(infixExpr)
    print('转换后的结果:', result)