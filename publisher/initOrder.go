package publisher

import (
	"WB_L0/internal/models"
)

func InitOrder() models.Order {
	delivery := models.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
	payment := models.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    2,
	}
	item := models.Items{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}
	itemArray := make([]models.Items, 0)
	itemArray = append(itemArray, item)

	//input := "2021-11-26T06:22:19Z"
	//layout := "2006-01-02T15:04:05Z"

	//t, _ := time.Parse(layout, input)
	order := models.Order{
		OrderUid:          "b563feb7b2b84b6test",
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Delivery:          delivery,
		Payment:           payment,
		Items:             itemArray,
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       "2006-01-02T15:04:05Z",
		OofShard:          "1",
	}
	return order
}
