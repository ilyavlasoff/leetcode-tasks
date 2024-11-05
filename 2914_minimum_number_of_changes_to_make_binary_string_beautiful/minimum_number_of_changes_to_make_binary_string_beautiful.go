package main

/*
Source: https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful

Вам дана 0-индексированная двоичная строка s, имеющая четную длину.
Строка "красивая", если ее можно разбить на одну или несколько подстрок таким образом, что:
- Каждая подстрока имеет четную длину.
- Каждая подстрока содержит только 1 или только 0.
- Вы можете изменить любой символ в s на 0 или 1.
Верните минимальное количество изменений, необходимых для того, чтобы сделать строку s "красивой".

Идея решения заключается в применении жадного алгоритма:
1) Так как количество символов в строке четное, значит есть возможность разбить 
строку на несколько четных последовательностей
2) Для поиска повторяющихся последовательностей проходим цикл по всем символам
 строки пока не встретим символ, прерывающий последовательность
3) Если последовательность символов до прерывания четная, то просто
 переходим к следующей последовательности
4) Если последовательность нечетная, то изменяем один следующий за ней символ на
 символ данной последовательности (инкрементируя количество изменений).
*/

func minChanges(s string) int {
    var changeCounter, curCounter int

    for i := 0; i != len(s); i++ {
        curCounter++
        if i != len(s)-1 && s[i] == s[i+1] {
            continue
        }

        isEven := curCounter % 2 == 0
        curCounter = 0

        if isEven {
            continue
        }

        changeCounter++
        i++
    }

    return changeCounter
}
