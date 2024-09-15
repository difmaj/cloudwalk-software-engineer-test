package parser

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/models/enums"
)

type Game struct {
	InitGame     [][]byte
	ClientEvents []*ClientEvent
	Kills        []*KillEvent
}

type ClientEvent struct {
	EventType []byte
	ClientID  int
	UserName  []byte
}

type KillEvent struct {
	KillerID     int
	VictimID     int
	KillMethodID int
}

type LogData struct {
	Games       []*Game
	currentGame *Game
}

const blankIndex = 6

// Nested map: eventLength -> eventName -> parsing function
var eventHandlers = map[string]func([]byte, *LogData) error{
	enums.EventInitGame:              parseInitGame,
	enums.EventClientConnect:         parseClientEventHandler,
	enums.EventClientUserinfoChanged: parseClientEventHandler,
	enums.EventClientBegin:           parseClientEventHandler,
	enums.EventKill:                  parseKillEventHandler,
}

func ParseLog(filePath string) (*LogData, error) {
	logData := &LogData{
		Games:       make([]*Game, 0),
		currentGame: nil,
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		// Ignore lines starting with '-'
		if line[blankIndex+1] == '-' {
			continue
		}

		var eventIndex int
		for eventIndex = blankIndex + 1; eventIndex < len(line); eventIndex++ {
			if line[eventIndex] == ':' {
				break
			}
		}

		// Check if there are handlers for this event length
		if handler, exists := eventHandlers[string(line[blankIndex+1:eventIndex])]; exists {
			// +3 to ignore the first trhe characters after the event name (': ')
			handler(line[blankIndex+enums.EventInitGameLength+2:], logData)
		} else {
			fmt.Println("No handler for event length", eventIndex-blankIndex)
			fmt.Println(string(line))
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return logData, nil
}

func parseInitGame(line []byte, logData *LogData) error {
	logData.currentGame = &Game{
		InitGame:     bytes.Split(line, []byte("\\")),
		ClientEvents: make([]*ClientEvent, 0),
		Kills:        make([]*KillEvent, 0),
	}
	logData.Games = append(logData.Games, logData.currentGame)
	return nil
}

func parseClientEventHandler(line []byte, logData *LogData) error {
	if logData.currentGame == nil {
		return errors.New("client event without a game")
	}

	parts := bytes.Fields(line)
	clientEvent := &ClientEvent{
		EventType: parts[0],
		ClientID:  parseInt(parts[1]),
	}
	if len(parts) > 2 {
		clientEvent.UserName = parts[2]
	}
	logData.currentGame.ClientEvents = append(logData.currentGame.ClientEvents, clientEvent)
	return nil
}

func parseKillEventHandler(line []byte, logData *LogData) error {
	if logData.currentGame == nil {
		return errors.New("kill event without a game")
	}

	parts := bytes.Fields(line)
	logData.currentGame.Kills = append(logData.currentGame.Kills, &KillEvent{
		KillerID:     parseInt(parts[1]),
		VictimID:     parseInt(parts[2]),
		KillMethodID: parseInt(parts[2]),
	})
	return nil
}

func parseInt(value []byte) int {
	var num int32
	err := binary.Read(bytes.NewReader(value), binary.BigEndian, &num)
	if err != nil {
		return 0
	}
	return int(num)
}
