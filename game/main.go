package main

import (
	"log"
	"strings"
)

type Character struct {
	backpack          bool
	door              bool
	backpackList      map[string]bool
	CharacterLocation *Location
}
type Items struct {
	name string
}
type Location struct {
	Name            string
	LocationMessage string
	Items           map[string]bool
	neighbourList   []*Location
}

var Player = new(Character)

func initGame() { //Функци инициализирует Мир, т.е. создаёт обьекты и даёт им начальные свойства.
	var room = new(Location)
	var kitchen = new(Location)
	var street = new(Location)
	var hallway = new(Location)

	var backpack = new(Items)
	var tea = new(Items)
	var keys = new(Items)
	var docs = new(Items)

	Player.CharacterLocation = kitchen
	Player.backpack = false
	Player.door = false
	Player.backpackList = make(map[string]bool)

	kitchen.Name = "кухня"
	kitchen.LocationMessage = "кухня, ничего интересного"
	kitchen.neighbourList = append(kitchen.neighbourList, hallway)
	kitchen.Items = make(map[string]bool)
	kitchen.Items["чай"] = true

	room.Name = "комната"
	room.LocationMessage = "ты в своей комнате. "
	room.neighbourList = append(room.neighbourList, hallway)
	room.Items = make(map[string]bool)
	room.Items["рюкзак"] = true
	room.Items["конспекты"] = true
	room.Items["ключи"] = true

	hallway.Name = "коридор"
	hallway.LocationMessage = "ничего интересного. "
	hallway.neighbourList = append(hallway.neighbourList, kitchen, room, street)

	street.Name = "улица"
	street.LocationMessage = "на улице весна. можно пройти - домой"
	street.neighbourList = append(street.neighbourList, hallway)

	tea.name = "чай"
	docs.name = "конспекты"
	backpack.name = "рюкзак"
	keys.name = "ключи"

	log.Println("Мир инициализирован")
}

func HandleCommand(command string) string {

	var words = strings.Split(command, " ")
	switch words[0] {
	case "осмотреться":
		return Player.LookAround()
	case "идти":
		return Player.ChangeLocation(words[1])
	case "применить":
		return Player.Use(words[1], words[2])
	case "взять":
		return Player.Take(words[1])
	case "надеть":
		return Player.Take(words[1])

	default:
		return "неизвестная команда"
	}

}

func main() {
}

func (Player *Character) LookAround() string {
	var answer string
	switch Player.CharacterLocation.Name {
	case "кухня":
		answer += "ты находишься на кухне, "
		if Player.CharacterLocation.Items["чай"] == true {
			answer += "на столе: чай, "
		}
		if Player.backpackList["конспекты"] != true || Player.backpackList["ключи"] != true {
			answer += "надо собрать рюкзак и идти в универ"
		} else {
			answer += "надо идти в универ"
		}
		return answer + Player.CharacterLocation.AbilityToGo()
	case "комната":
		if Player.CharacterLocation.Items["конспекты"] != false && Player.CharacterLocation.Items["ключи"] != false && Player.CharacterLocation.Items["рюкзак"] != false {
			answer = "на столе: ключи, конспекты, на стуле: рюкзак" + Player.CharacterLocation.AbilityToGo()
			return answer
		}
		if Player.CharacterLocation.Items["конспекты"] != false && Player.CharacterLocation.Items["ключи"] != false && Player.CharacterLocation.Items["рюкзак"] != true {
			answer = "на столе: ключи, конспекты" + Player.CharacterLocation.AbilityToGo()
			return answer
		}
		if Player.CharacterLocation.Items["конспекты"] != true && Player.CharacterLocation.Items["ключи"] != true {
			answer = "пустая комната" + Player.CharacterLocation.AbilityToGo()
			return answer
		} else {
			answer += "на столе:"
		}
		if Player.CharacterLocation.Items["конспекты"] == false && Player.CharacterLocation.Items["ключи"] == true {
			answer += " ключи,"
		} else {
			answer += " конспекты"
		}
		return answer + Player.CharacterLocation.AbilityToGo()

	default:
		return "ничего интересного" + Player.CharacterLocation.AbilityToGo()
	}

}

func (loc *Location) AbilityToGo() string {
	var answer string
	for i, loc := range loc.neighbourList {
		if i > len(loc.neighbourList)-1 {
			answer += ", "
		}
		answer += loc.Name
	}
	return ". можно пройти - " + answer
}

func (Player *Character) Take(object string) string {
	//убрать значение в команте (вещь пропала)
	var answer string

	if object == "рюкзак" {
		answer = "вы надели: рюкзак"
		Player.backpack = true
		Player.CharacterLocation.Items[object] = false
		return answer
	}

	if Player.backpack == true {
		if Player.CharacterLocation.Items[object] == true {
			Player.backpackList[object] = true
			Player.CharacterLocation.Items[object] = false
			answer = "предмет добавлен в инвентарь: " + object
			return answer
		}

	} else {
		answer = "некуда класть"
		return answer
	}
	answer = "нет такого"
	return answer
}

func (Player *Character) ChangeLocation(nextLocation string) string {
	var answer string
	if nextLocation == "улица" && Player.CharacterLocation.Name == "коридор" && Player.door != true {
		answer = "дверь закрыта"
		return answer
	} else {
		if nextLocation == "улица" && Player.CharacterLocation.Name == "коридор" && Player.door == true {
			answer = "на улице весна. можно пройти - домой"
			return answer
		}
	}

	for _, v := range Player.CharacterLocation.neighbourList {

		if v.Name == nextLocation && nextLocation == "комната" {
			Player.CharacterLocation = v
			answer = "ты в своей комнате. можно пройти - коридор"
			return answer
		}

		if v.Name == nextLocation && nextLocation == "кухня" {
			Player.CharacterLocation = v
			answer = Player.CharacterLocation.LocationMessage + Player.CharacterLocation.AbilityToGo()
			return answer
		}
		if v.Name == nextLocation {
			Player.CharacterLocation = v
			answer = Player.LookAround()
			return answer
		}
	}
	answer = "нет пути в " + nextLocation
	return answer
}

func (Player *Character) Use(object, place string) string {
	var answer string
	if place == "дверь" {
		if object == "ключи" && Player.backpackList["ключи"] == true {
			answer = "дверь открыта"
			Player.door = true
			return answer
		} else {
			if object == "ключи" && Player.backpackList["ключи"] == false {
				answer = "нет предмета в инвентаре - ключи"
				return answer
			}
		}
	} else {
		if object == "ключи" {
			answer = "не к чему применить"
			return answer
		} else {
			answer = "нет предмета в инвентаре - " + object
			return answer
		}
	}
	return answer
}
