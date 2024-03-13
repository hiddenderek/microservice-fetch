package service

import (
	"cmp"
	"fetch-service/internal/database"
	"fetch-service/internal/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func GetPoints(id string) (models.PointResponse, error, bool) {
	item, ok := database.GetData[models.Receipt](id)
	if ok {
		points, err := accumulatePoints(item)
		pointResponse := models.PointResponse{
			Points: points,
		}
		if err != nil {
			return pointResponse, err, true
		} else {
			return pointResponse, nil, true
		}
	} else {
		pointResponse := models.PointResponse{
			Points: 0,
		}
		return pointResponse, nil, false
	}
}

func accumulatePoints(item models.Receipt) (points int, err error) {
	score1 := scoreRetailerName(item.Retailer)
	score2 := scoreReceiptLength(item.Items)
	score3, err1 := scoreTotalPrice(item.Total)
	score4, err2 := scoreItemDescriptions(item.Items)
	score5, err3 := scorePurchaseDate(item.PurchaseDate)
	score6, err4 := scorePurchaseTime(item.PurchaseTime)
	err = cmp.Or(err1, err2, err3, err4)
	if err != nil {
		return 0, err
	} else {
		points += score1
		points += score2
		points += score3
		points += score4
		points += score5
		points += score6
	}
	fmt.Println("Total Points", points)
	return points, nil
}

func scoreRetailerName(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	fmt.Println("Number of alphanumeric characters: ", count)
	return count
}

func scoreTotalPrice(s string) (int, error) {
	total, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// this should not happen
		return 0, fmt.Errorf("error parsing total")
	} else {
		score := 0
		// accumulate points based on the number
		if math.Mod(total, 1) == 0 {
			fmt.Println("Total is an integer")
			score += 50
		}
		// both conditions should be able to accumulate if they are both met
		if math.Mod(total, .25) == 0 {
			fmt.Println("Total is divisible by .25")
			score += 25
		}
		return score, nil
	}
}

func scoreReceiptLength(items []models.Item) int {
	var itemLength = len(items)
	var score = int(itemLength/2) * 5
	fmt.Printf("Total receipt length score: %d\n", score)
	return score
}

func scoreItemDescriptions(items []models.Item) (int, error) {
	score := 0
	for _, item := range items {
		trimmedLength := len(strings.Trim(item.ShortDescription, " "))
		// if the trimmed length is divisible by three, calculate the new points
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				// this should not happen
				return 0, fmt.Errorf("error parsing price")
			}
			points := int(math.Ceil(price * .2))
			fmt.Println("Trimed receipt description is divisible by three, new points: ", points)
			score += points
		}
	}
	return score, nil
}

func scorePurchaseDate(dateString string) (int, error) {
	parsedDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		// this should not happen
		return 0, fmt.Errorf("error parsing purchase date")
	}

	day := parsedDate.Day()
	if day%2 == 0 {
		fmt.Println("Purchase day is even: ", day)
		return 0, nil
	} else {
		fmt.Println("Purchase day is odd: ", day)
		return 6, nil
	}
}

func scorePurchaseTime(timeString string) (int, error) {
	time := strings.Split(timeString, ":")
	hour, err := strconv.Atoi(time[0])
	if err != nil {
		// this should not happen
		return 0, fmt.Errorf("error parsing hour")
	}

	minute, err := strconv.Atoi(time[1])
	if err != nil {
		// this should not happen
		return 0, fmt.Errorf("error parsing minute")
	}
	adjHour := hour - 12
	// Check the hours and minutes. Has to be in between 2 and 4.
	if adjHour >= 2 && adjHour <= 4 && (minute != 0 || adjHour == 3) {
		fmt.Println("Purchase time is between 2 and 4 PM")
		return 10, nil
	} else {
		fmt.Println("Purchase time is not between 2 and 4 PM")
		return 0, nil
	}
}
