package main

import (
	"fmt"
	"strings"
)

const filledStar = "★" // "●"
const emptyStar = "☆"  // "○"

func calculateRating(score int, thresholds []int) int {
	if score < 0 {
		return 0
	}

	for i, bound := range thresholds {
		if score <= bound {
			return i + 1
		}
	}

	return 0
}

func renderStarRating(rating int, totalStars int) string {
	filled := strings.Repeat(filledStar, rating)
	empty := strings.Repeat(emptyStar, totalStars-rating)
	return filled + empty
}

func drawRating(score int, thresholds []int) string {
	ratingValue := calculateRating(score, thresholds)
	totalStars := len(thresholds)
	return renderStarRating(ratingValue, totalStars)
}

func main() {
	thresholds5Star := []int{20, 40, 60, 80, 100}

	fmt.Println("--- 5-Star System ---")
	fmt.Printf("Score 75: %s\n", drawRating(75, thresholds5Star))   // 4
	fmt.Printf("Score 30: %s\n", drawRating(30, thresholds5Star))   // 2
	fmt.Printf("Score 110: %s\n", drawRating(110, thresholds5Star)) // 0
	fmt.Printf("Score -5: %s\n", drawRating(-5, thresholds5Star))   // 0

	thresholds3Star := []int{30, 70, 100}
	fmt.Println("\n--- 3-Star System ---")
	fmt.Printf("Score 25: %s\n", drawRating(25, thresholds3Star))   // 1
	fmt.Printf("Score 31: %s\n", drawRating(31, thresholds3Star))   // 2
	fmt.Printf("Score 100: %s\n", drawRating(100, thresholds3Star)) // 3

	emptyThresholds := []int{}
	fmt.Println("\n--- Empty System ---")
	fmt.Printf("Score 50: %s\n", drawRating(50, emptyThresholds)) // empty

}
