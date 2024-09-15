package tests_test

import (
	"testing"
)

var logLines = [][]byte{
	// \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
	[]byte("  0:00 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\bot_minplayers\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0"),
	[]byte("  0:15 ClientConnect: 2"),
	[]byte("  0:16 ClientUserinfoChanged: 2 n\\Dono da Bola\\t\\0\\model\\sarge\\hmodel\\sarge\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\95\\w\\0\\l\\0\\tt\\0\\tl\\0"),
	[]byte("  0:22 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT"),
}

var blankIndex = 6

var matchByLengthAndStringEventHandlers = map[int]map[string]func([]byte){
	10: {"InitGame": nil},
	13: {"ClientConnect": nil},
	24: {"ClientUserinfoChanged": nil},
	4:  {"Kill": nil},
}

func matchByLengthAndString(line []byte) {

	var eventIndex int
	for eventIndex = blankIndex + 1; eventIndex < len(line); eventIndex++ {
		if line[eventIndex] == ':' {
			break
		}
	}

	if eventMap, exists := matchByLengthAndStringEventHandlers[eventIndex-blankIndex]; exists {
		if len(eventMap) == 1 {
			for _, handler := range eventMap {
				_ = handler
			}
		} else if handler, exists := eventMap[string(line[blankIndex+1:eventIndex])]; exists {
			_ = handler
		}
	}
}

var matchByStringEventHandlers = map[string]func([]byte){
	"InitGame":              nil,
	"ClientConnect":         nil,
	"ClientUserinfoChanged": nil,
	"Kill":                  nil,
}

func matchByString(line []byte) {
	var eventIndex int
	for eventIndex = blankIndex + 1; eventIndex < len(line); eventIndex++ {
		if line[eventIndex] == ':' {
			break
		}
	}

	if handler, exists := matchByStringEventHandlers[string(line[blankIndex+1:eventIndex])]; exists {
		_ = handler
	}
}

// Just checking the performance of the two methods to decide which one to use.
func BenchmarkMatchByString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, line := range logLines {
			matchByString(line)
		}
	}
}

func BenchmarkMatchByLengthAndString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, line := range logLines {
			matchByLengthAndString(line)
		}
	}
}
