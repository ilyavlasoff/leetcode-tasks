package main

/*
Source: https://leetcode.com/problems/successful-pairs-of-spells-and-potions

Вам даны два положительных целочисленных массива spells и potions, длиной n и m соответственно, где spells[i] представляет силу i-го заклинания, а potions[j] представляет силу j-го зелья.
Вам также дан целочисленный успех. Пара заклинания и зелья считается успешной, если произведение их сил составляет по крайней мере успех.
Верните целочисленный массив pairs длиной n, где pairs[i] — это количество зелий, которые образуют успешную пару с i-м заклинанием.
*/

func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions) // отсортируем значения для бинарного поиска

    var products []int

    for _, s := range spells {
        m := success / int64(s)

        f := findIndex(potions, m)
        for f > 0 && int64(s * potions[f]) == success {
            f-- // если в результате бин поиска нашлось несколько элементов, идем к первому из группы
        }
        for f < len(potions) && int64(s * potions[f]) < success {
            f++ // если деление выше было не ровным, то пройдем вперед до тех пор пока spells[i]*potions[j] >= success
        }

        products = append(products, len(potions)-f)
    }

    return products
}

func findIndex(values []int, ref int64) int {
	var s, e = 0, len(values) - 1

    if int64(values[s]) == ref {
        return s
    }
    if int64(values[e]) == ref {
        return e
    }

	for s < e {
		m := (s + e) / 2
		if int64(values[m]) == ref {
			return m
		} else if int64(values[m]) < ref {
			s = m + 1
		} else {
			e = m
		}
	}

	return s
}