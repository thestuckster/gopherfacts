package clients

type Cooldown struct {
	TotalSeconds     int    `json:"total_seconds"`
	RemainingSeconds int    `json:"remaining_seconds"`
	StartedAt        string `json:"started_at"`
	Expiration       string `json:"expiration"`
	Reason           string `json:"reason"`
}

type Drop struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type BlockedHits struct {
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Water int `json:"water"`
	Air   int `json:"air"`
	Total int `json:"total"`
}

type Fight struct {
	Xp                 int         `json:"xp"`
	Gold               int         `json:"gold"`
	Drops              []Drop      `json:"drops"`
	Turns              int         `json:"turns"`
	MonsterBlockedHits BlockedHits `json:"monster_blocked_hits"`
	PlayerBlockedHits  BlockedHits `json:"player_blocked_hits"`
	Logs               []string    `json:"logs"`
	Result             string      `json:"result"`
}

//TODO CharacterStruct
