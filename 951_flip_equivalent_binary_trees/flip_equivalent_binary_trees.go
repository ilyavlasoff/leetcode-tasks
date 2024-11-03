package main

/*
Source: https://leetcode.com/problems/flip-equivalent-binary-trees

For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.
A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.
Given the roots of two binary trees root1 and root2, return true if the two trees are flip equivalent or false otherwise.
*/

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    var compareFlip func(root1 *TreeNode, root2 *TreeNode) bool

    compareFlip = func(root1 *TreeNode, root2 *TreeNode) bool {
        isNilR1, isNilR2 := root1==nil, root2==nil

        if isNilR1 || isNilR2 {
            return isNilR1 && isNilR2
        }

        if root1.Val != root2.Val {
            return false
        }

        return (compareFlip(root1.Left, root2.Left) && compareFlip(root1.Right, root2.Right) || compareFlip(root1.Left, root2.Right) && compareFlip(root1.Right, root2.Left))
    }

    return compareFlip(root1, root2)
}