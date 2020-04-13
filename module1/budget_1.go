package module1

// Budget stores budget information

type Budget struct {
	float32 Max
	Items   []Item
}

// Item stores item information

type Item struct {
	string  Description
	float32 Price
}
