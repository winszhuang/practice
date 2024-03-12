package main

func main() {

}

type Player struct {
	hand           []string
	communityCards []string
}

// 12:25
func isPlayerHandGetOneOrMoreSpecialCards(p *Player, handRankStr string) bool {
	switch handRankStr {
	case "four_of_a_kind":
		allCards := append(p.hand, p.communityCards...)
		main := findMostFrequentRanks(allCards)
		print(main)
		return false
	}
	return true
}

func findMostFrequentRanks(cards []string) []string {
	if len(cards) < 3 {
		return []string{}
	}

	// 轉換成數字
	var points []string
	for _, card := range cards {
		points = append(points, card[:len(card)-1])
	}

	pointsMap := make(map[string]int)
	for i := 0; i < len(points); i++ {
		pointsMap[points[i]] += 1
	}

	maxNum := pointsMap[points[0]]
	for _, num := range pointsMap {
		if num > maxNum {
			maxNum = num
		}
	}

	var mainPoint string
	for pointKey, num := range pointsMap {
		if num == maxNum {
			mainPoint = pointKey
		}
	}

	var result []string
	for i, point := range points {
		if point == mainPoint {
			result = append(result, cards[i])
		}
	}

	return result
}
