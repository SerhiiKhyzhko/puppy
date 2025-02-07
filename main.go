package puppy

var someText = "Вдалого полювання, Сталкере!"
const PI = 3.14159

func CirkuleSquere(r float64) float64{
	return PI * (r * r)
}

func Messege() string {
	return someText
}

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof!, Woof!, Woof!"
}