package models

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required,numeric"`
}

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Total        string `json:"total" binding:"required,numeric"`
	Items        []Item `json:"items" binding:"required"`
}

type IdResponse struct {
	Id string `json:"id" binding:"required,uuid"`
}

type PointResponse struct {
	Points int `json:"points" binding:"required"`
}
