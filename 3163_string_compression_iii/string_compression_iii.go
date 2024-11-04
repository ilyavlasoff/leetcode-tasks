package main

/*
Source: https://leetcode.com/problems/string-compression-iii

Given a string word, compress it using the following algorithm:

Begin with an empty string comp. While word is not empty, use the following operation:
Remove a maximum length prefix of word made of a single character c repeating at most 9 times.
Append the length of the prefix followed by c to comp.
Return the string comp.
*/

func compressedString(word string) string {
    wlen := len(word)
    idx := 0

    output := strings.Builder{}

    for idx < wlen {
        var consCount int
        curChar := word[idx]

        for idx < wlen && word[idx] == curChar && consCount < 9 {
            idx++
            consCount++
        }

        output.WriteString(strconv.Itoa(consCount))
        output.WriteByte(curChar)
    }

    return output.String()
}