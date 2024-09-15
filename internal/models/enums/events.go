package enums

// Matches events from the log file.
const (
	EventInitGame              = "InitGame"
	EventClientConnect         = "ClientConnect"
	EventClientUserinfoChanged = "ClientUserinfoChanged"
	EventClientBegin           = "ClientBegin"
	EventKill                  = "Kill"
)

// EventLengths is a map of event names to their lengths.
const (
	EventInitGameLength              = 9
	EventClientConnectLength         = 14
	EventClientUserinfoChangedLength = 22
	EventClientBeginLength           = 12
	EventKillLength                  = 5
)
