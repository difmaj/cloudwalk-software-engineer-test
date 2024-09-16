package parser

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/models/enums"
)

type Game struct {
	InitGame [][]byte
	Clients  []*Client
	Kills    []*KillEvent
}

type Client struct {
	ClientID int
	UserName []byte
}

type KillEvent struct {
	KillerID     int
	VictimID     int
	KillMethodID int
}

type LogData struct {
	Games   []*Game
	Current *Game
}

const blankIndex = 6

// Nested map: eventLength -> eventName -> parsing function
var eventHandlers = map[string]func([]byte, *LogData) error{
	enums.EventInitGame:              ParseInitGameEventHandler,
	enums.EventClientConnect:         ParseClientConectedEventHandler,
	enums.EventClientUserinfoChanged: ParseClientUserinfoChangedEventHandler,
	enums.EventKill:                  ParseKillEventHandler,
}

func ParseLog(filePath string) (*LogData, error) {
	logData := &LogData{
		Games:   make([]*Game, 0),
		Current: nil,
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
			handler(line[eventIndex+2:], logData)
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

func ParseInitGameEventHandler(line []byte, logData *LogData) error {
	info := bytes.Split(line, []byte("\\"))[1:]

	logData.Current = &Game{
		Clients: make([]*Client, 0),
		Kills:   make([]*KillEvent, 0),
	}
	logData.Current.InitGame = make([][]byte, len(info))
	copy(logData.Current.InitGame, info)

	logData.Games = append(logData.Games, logData.Current)
	return nil
}

func ParseClientConectedEventHandler(line []byte, logData *LogData) error {
	if logData.Current == nil {
		return errors.New("client event without a game")
	}

	// check if the client is already in the game
	for _, client := range logData.Current.Clients {
		if client.ClientID == ParseInt(line) {
			return nil
		}
	}

	logData.Current.Clients = append(logData.Current.Clients, &Client{ClientID: ParseInt(line)})
	return nil
}

func ParseClientUserinfoChangedEventHandler(line []byte, logData *LogData) error {
	if logData.Current == nil {
		return errors.New("client event without a game")
	}

	var spaceIndex int
	for spaceIndex = 0; spaceIndex < len(line); spaceIndex++ {
		if line[spaceIndex] == ' ' {
			break
		}
	}

	info := bytes.Split(line[spaceIndex+1:], []byte("\\"))
	if len(info) < 2 {
		return errors.New("invalid ClientUserinfoChanged event")
	}

	// check if the client is already in the game
	for _, client := range logData.Current.Clients {
		if client.ClientID == ParseInt(line[:spaceIndex]) {
			client.UserName = make([]byte, len(info[1]))
			copy(client.UserName, info[1])
		}
	}
	return nil
}

func ParseKillEventHandler(line []byte, logData *LogData) error {
	if logData.Current == nil {
		return errors.New("kill event without a game")
	}

	var colonIndex int
	for colonIndex = 0; colonIndex < len(line); colonIndex++ {
		if line[colonIndex] == ':' {
			break
		}
	}

	parts := bytes.Fields(line[:colonIndex])
	logData.Current.Kills = append(logData.Current.Kills, &KillEvent{
		KillerID:     ParseInt(parts[0]),
		VictimID:     ParseInt(parts[1]),
		KillMethodID: ParseInt(parts[2]),
	})
	return nil
}

func ParseInt(value []byte) int {
	result := 0
	sign := 1
	start := 0

	// Check if the number is negative
	if len(value) > 0 && value[0] == '-' {
		sign = -1
		start = 1
	}

	// Iterate over each byte and calculate the integer value
	for i := start; i < len(value); i++ {
		result = result*10 + int(value[i]-'0')
	}

	return result * sign
}
