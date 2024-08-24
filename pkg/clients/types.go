package clients

type DestinationData struct {
	Name    string     `json:"name"`
	Skin    string     `json:"skin"`
	X       int        `json:"x"`
	Y       int        `json:"y"`
	Content MapContent `json:"content"`
}

type MapContent struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

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

type CharacterSchema struct {
	Name                    string          `json:"name"`
	Skin                    string          `json:"skin"`
	Level                   int             `json:"level"`
	XP                      int             `json:"xp"`
	MaxXP                   int             `json:"max_xp"`
	TotalXP                 int             `json:"total_xp"`
	Gold                    int             `json:"gold"`
	Speed                   int             `json:"speed"`
	MiningLevel             int             `json:"mining_level"`
	MiningXP                int             `json:"mining_xp"`
	MiningMaxXP             int             `json:"mining_max_xp"`
	WoodcuttingLevel        int             `json:"woodcutting_level"`
	WoodcuttingXP           int             `json:"woodcutting_xp"`
	WoodcuttingMaxXP        int             `json:"woodcutting_max_xp"`
	FishingLevel            int             `json:"fishing_level"`
	FishingXP               int             `json:"fishing_xp"`
	FishingMaxXP            int             `json:"fishing_max_xp"`
	WeaponcraftingLevel     int             `json:"weaponcrafting_level"`
	WeaponcraftingXP        int             `json:"weaponcrafting_xp"`
	WeaponcraftingMaxXP     int             `json:"weaponcrafting_max_xp"`
	GearcraftingLevel       int             `json:"gearcrafting_level"`
	GearcraftingXP          int             `json:"gearcrafting_xp"`
	GearcraftingMaxXP       int             `json:"gearcrafting_max_xp"`
	JewelrycraftingLevel    int             `json:"jewelrycrafting_level"`
	JewelrycraftingXP       int             `json:"jewelrycrafting_xp"`
	JewelrycraftingMaxXP    int             `json:"jewelrycrafting_max_xp"`
	CookingLevel            int             `json:"cooking_level"`
	CookingXP               int             `json:"cooking_xp"`
	CookingMaxXP            int             `json:"cooking_max_xp"`
	HP                      int             `json:"hp"`
	Haste                   int             `json:"haste"`
	CriticalStrike          int             `json:"critical_strike"`
	Stamina                 int             `json:"stamina"`
	AttackFire              int             `json:"attack_fire"`
	AttackEarth             int             `json:"attack_earth"`
	AttackWater             int             `json:"attack_water"`
	AttackAir               int             `json:"attack_air"`
	DamageFire              int             `json:"dmg_fire"`
	DamageEarth             int             `json:"dmg_earth"`
	DamageWater             int             `json:"dmg_water"`
	DamageAir               int             `json:"dmg_air"`
	ResistanceFire          int             `json:"res_fire"`
	ResistanceEarth         int             `json:"res_earth"`
	ResistanceWater         int             `json:"res_water"`
	ResistanceAir           int             `json:"res_air"`
	X                       int             `json:"x"`
	Y                       int             `json:"y"`
	Cooldown                int             `json:"cooldown"`
	CooldownExpiration      string          `json:"cooldown_expiration"`
	WeaponSlot              string          `json:"weapon_slot"`
	ShieldSlot              string          `json:"shield_slot"`
	HelmetSlot              string          `json:"helmet_slot"`
	BodyArmorSlot           string          `json:"body_armor_slot"`
	LegArmorSlot            string          `json:"leg_armor_slot"`
	BootsSlot               string          `json:"boots_slot"`
	Ring1Slot               string          `json:"ring1_slot"`
	Ring2Slot               string          `json:"ring2_slot"`
	AmuletSlot              string          `json:"amulet_slot"`
	Artifact1Slot           string          `json:"artifact1_slot"`
	Artifact2Slot           string          `json:"artifact2_slot"`
	Artifact3Slot           string          `json:"artifact3_slot"`
	Consumable1Slot         string          `json:"consumable1_slot"`
	Consumable1SlotQuantity int             `json:"consumable1_slot_quantity"`
	Consumable2Slot         string          `json:"consumable2_slot"`
	Consumable2SlotQuantity int             `json:"consumable2_slot_quantity"`
	Task                    string          `json:"task"`
	TaskType                string          `json:"task_type"`
	TaskProgress            int             `json:"task_progress"`
	TaskTotal               int             `json:"task_total"`
	InventoryMaxItems       int             `json:"inventory_max_items"`
	Inventory               []InventorySlot `json:"inventory"`
}

type Item struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type InventorySlot struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
	Slot     int    `json:"slot"`
}

type CraftDetails struct {
	XpGained int    `json:"xp"`
	Items    []Item `json:"items"`
}
