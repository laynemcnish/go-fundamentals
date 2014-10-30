package main

func FruitOrVegetable(v string) string {
	switch v {
	case "Apple", "Orange", "Banana":
		return "Fruit"
	case "Onion", "Broccoli", "Celery":
		return "Vegetable"
	default:
		return "Neither"
	}
	panic("unreachable")
}
