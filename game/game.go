package game

import (
	"log"
	"net/http"
	"time"
)

type Unit struct {
	xPos, yPos float64
}

func (unit *Unit) aplyPos(xPos, yPos float64) {
	unit.xPos = xPos
	unit.yPos = yPos
}

const (
	TOP    = 1
	LEFT   = 2
	RIGHT  = 3
	BOTTOM = 4
)

func (unit *Unit) handleAction(action int) (float64, float64) {
	switch action {
	case TOP:
		return unit.xPos, unit.yPos + 1
	case BOTTOM:
		return unit.xPos, unit.yPos - 1
	case RIGHT:
		return unit.xPos + 1, unit.yPos
	case LEFT:
		return unit.xPos - 1, unit.yPos
	default:
		return unit.xPos, unit.yPos
	}
}

type Field struct {
	width, hight float64
	units        []Unit
}

func (field *Field) addUnit(u Unit) {
	field.units = append(field.units, u)
}

func (field *Field) applyAction(action Action) {
	x, y := field.units[action.unitId].handleAction(action.actionType)
	if x > field.width && x < 0 {
		log.Println("X out of bounds: ", x)
		return
	}
	if y > field.hight && y < 0 {
		log.Println("Y out of bounds: ", y)
		return
	}
	field.units[action.unitId].aplyPos(x, y)
}

type Action struct {
	unitId     int
	actionType int
}

type Game struct {
	field Field
}

func (game *Game) GameLoop(actions chan Action, units chan Unit) {
	var action Action
	var unit Unit
	for {
		select {
		case action = <-actions:
			game.field.applyAction(action)
		case unit = <-units:
			game.field.addUnit(unit)
		default:
			log.Println(game.field.units)
		}
	}
}

func gameMain() {
	game := Game{
		field: Field{width: 300, hight: 300, units: make([]Unit, 0)},
	}
	actions := make(chan Action, 4)
	units := make(chan Unit, 4)
	go game.GameLoop(actions, units)
	units <- Unit{100, 100}
	for {
		time.Sleep(200 * time.Microsecond)
		actions <- Action{0, TOP}
	}
}

func HandleResponseToGame(rw http.ResponseWriter, request *http.Request) {
	//log.Println(request)
	//rw.Header().Set("Access-Control-Request-Headers", "*")
	rw.Header().Set("Access-Control-Allow-Origin", "*") // TODO: добавить доступ только для определённых сайтов
	//
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("S H I T"))
	log.Println("Message sended")
	//rw.WriteHeader(http.StatusOK)
}
