## Текстовая игра на go

### Цель проекта

В рамках выполнения проекта, научился писать базовый код, поработал со структурами, методами, функциями, if-ами. 

Поупражнялся в моделировании объектов

### Описание
Написана простая игра, которая реагирует на команды игрока.

Игровой мир состоит из комнат, где может происходить какое-то действие.
Так же есть игрок.
Как у игрока, так и у команты есть состояние.

initGame делает нового игрока и задаёт ему начальное состояние.

Команда в handleCommand парсится как:

``` bash
$команда $параметр1 $параметр2 $параметр3
```

Запускать тест через

``` bash
go test -v
```
находясь в папке `game`.
