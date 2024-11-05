package main

/*
Source: https://leetcode.com/problems/delete-characters-to-make-fancy-string

A fancy string is a string where no three consecutive characters are equal.

Given a string s, delete the minimum possible number of characters from s to make it fancy.

Return the final string after the deletion. It can be shown that the answer will always be unique.
*/

func makeFancyString(s string) string {
    result := strings.Builder{}

    var counter int
    for i := 0 ; i != len(s); i++ {
        counter++
        if i != len(s)-1 && s[i] == s[i+1] {
            continue
        }

        for j := 0; j < 2 && j < counter; j++ {
            result.WriteByte(s[i])
        }
        counter = 0
    }

    return result.String()
}