package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int, len(nums))
	for i, n := range nums {
		if _, ok := numsMap[n]; !ok {
			numsMap[n] = []int{i}
		} else {
			numsMap[n] = append(numsMap[n], i)
		}

	}

	for k, v := range numsMap {
		d := target - k
		if dv, ok := numsMap[d]; ok {
			if len(v) > 1 {
				return v[:2]
			} else if v[0] != dv[0] {
				return []int{v[0], dv[0]}
			}
		}
	}

	return nil
}
