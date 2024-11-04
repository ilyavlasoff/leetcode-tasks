package main

/*
Source: https://leetcode.com/problems/binary-tree-paths

Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.
*/

func binaryTreePaths(root *TreeNode) []string {
    var dfs func(*TreeNode)
    var paths []string
    var curPath []int

    if root == nil {
        return paths
    }

    dfs = func(root *TreeNode) {
        curPath = append(curPath, root.Val)
        defer func() {
            curPath = curPath[:len(curPath)-1]
        }()

        if root.Right == nil && root.Left == nil {
            paths = append(paths, formatPath(curPath))

            return
        }

        if root.Right != nil {
            dfs(root.Right)
        }
        if root.Left != nil {
            dfs(root.Left)
        }
    }

    dfs(root)

    return paths
}

func formatPath(paths []int) string {
    var formatted string

    for i, p := range paths {
        if i != 0 {
            formatted += "->"
        }
        formatted += strconv.Itoa(p)
    }

    return formatted
}