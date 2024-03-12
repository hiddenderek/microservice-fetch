package service

import (
	"fetch-service/internal/models"
	"testing"
)

func TestAccumulatePointsName(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "WalgreensisCool",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.31419",
		Items:        []models.Item{},
	}
	score := accumulatePoints(testData)
	if score != 15 {
		t.Fatalf(`Must be equal to 15`)
	}
}

func TestAccumulatePointsNameNonAlphaNumeric(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Wal_gree&nsisCool%",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.31419",
		Items:        []models.Item{},
	}
	score := accumulatePoints(testData)
	if score != 15 {
		t.Fatalf(`Must be equal to 15`)
	}
}

func TestAccumulatePointsTotalRandom(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.31419",
		Items:        []models.Item{},
	}
	score := accumulatePoints(testData)
	if score != 9 {
		t.Fatalf(`Must be equal to 9`)
	}
}

func TestAccumulatePointsTotalQuarter(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.25",
		Items:        []models.Item{},
	}
	score := accumulatePoints(testData)
	if score != 34 {
		t.Fatalf(`Must be equal to 34`)
	}
}

func TestAccumulatePointsTotalWhole(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2",
		Items:        []models.Item{},
	}
	score := accumulatePoints(testData)
	if score != 84 {
		t.Fatalf(`Must be equal to 84`)
	}
}

func TestNotAccumulatePointsOneItem(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2",
		Items:        []models.Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}},
	}
	score := accumulatePoints(testData)
	if score != 84 {
		t.Fatalf(`Must be equal to 84`)
	}
}

func TestAccumulatePointsTwoItems(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 89 {
		t.Fatalf(`Must be equal to 89`)
	}
}

func TestAccumulatePointsFiveItems(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 94 {
		t.Fatalf(`Must be equal to 94`)
	}
}

func TestAccumulatePurchaseDateOdd(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-05",
		PurchaseTime: "08:13",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 100 {
		t.Fatalf(`Must be equal to 100`)
	}
}

func TestAccumulatePurchaseTime(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-05",
		PurchaseTime: "14:13",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 110 {
		t.Fatalf(`Must be equal to 110`)
	}
}

func TestAccumulatePurchaseTimeExactlyTwo(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-05",
		PurchaseTime: "14:00",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 100 {
		t.Fatalf(`Must be equal to 100`)
	}
}

func TestAccumulatePurchaseTimeExactlyThree(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-05",
		PurchaseTime: "15:00",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 110 {
		t.Fatalf(`Must be equal to 110`)
	}
}

func TestAccumulateReceiptPrices(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-05",
		PurchaseTime: "15:00",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "Dasani", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "coke", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 111 {
		t.Fatalf(`Must be equal to 111`)
	}
}

func TestAccumulateReceiptPricesTrimmed(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-05",
		PurchaseTime: "15:00",
		Total:        "2",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "  Coke is cool  ", Price: "1.40"},
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
			{ShortDescription: "  Coke is cool  ", Price: "1.40"},
			{ShortDescription: "coke", Price: "1.40"}},
	}
	score := accumulatePoints(testData)
	if score != 112 {
		t.Fatalf(`Must be equal to 112`)
	}
}

func TestProvidedExample1(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}},
	}
	score := accumulatePoints(testData)
	if score != 28 {
		t.Fatalf(`Must be equal to 28`)
	}
}

func TestProvidedExample2(t *testing.T) {
	testData := models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []models.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"}},
	}
	score := accumulatePoints(testData)
	if score != 109 {
		t.Fatalf(`Must be equal to 109`)
	}
}
