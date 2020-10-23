package main

func numJewelsInStones(jewels string, stones string) (count int) {
	stoneInventory := make(map[rune]int)

	for _, stone := range stones {
		stoneInventory[stone]++
	}

	for _, jewel := range jewels {
		if stoneCount, ok := stoneInventory[jewel]; ok {
			count += stoneCount
		}
	}

	return
}
