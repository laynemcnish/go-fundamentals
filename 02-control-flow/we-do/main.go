package main

func FruitOrVegatable(v string) string {
	switch v {
	case "Apple", "Orange", "Banana":
		return "Fruit"
	case "Onion", "Broccoli", "Celery":
		return "Vegatable"
	default:
		return "Neither"
	}

}
