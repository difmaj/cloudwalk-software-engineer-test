package enums

var DeathMeans = map[DeathMeansID]string{
	Unknown:       "MOD_UNKNOWN",
	Shotgun:       "MOD_SHOTGUN",
	Gauntlet:      "MOD_GAUNTLET",
	Machinegun:    "MOD_MACHINEGUN",
	Grenade:       "MOD_GRENADE",
	GrenadeSplash: "MOD_GRENADE_SPLASH",
	Rocket:        "MOD_ROCKET",
	RocketSplash:  "MOD_ROCKET_SPLASH",
	Plasma:        "MOD_PLASMA",
	PlasmaSplash:  "MOD_PLASMA_SPLASH",
	Railgun:       "MOD_RAILGUN",
	Lightning:     "MOD_LIGHTNING",
	Bfg:           "MOD_BFG",
	BfgSplash:     "MOD_BFG_SPLASH",
	Water:         "MOD_WATER",
	Slime:         "MOD_SLIME",
	Lava:          "MOD_LAVA",
	Crush:         "MOD_CRUSH",
	Telefrag:      "MOD_TELEFRAG",
	Falling:       "MOD_FALLING",
	Suicide:       "MOD_SUICIDE",
	TargetLaser:   "MOD_TARGET_LASER",
	TriggerHurt:   "MOD_TRIGGER_HURT",
}

type DeathMeansID int

const (
	Unknown DeathMeansID = iota
	Shotgun
	Gauntlet
	Machinegun
	Grenade
	GrenadeSplash
	Rocket
	RocketSplash
	Plasma
	PlasmaSplash
	Railgun
	Lightning
	Bfg
	BfgSplash
	Water
	Slime
	Lava
	Crush
	Telefrag
	Falling
	Suicide
	TargetLaser
	TriggerHurt
)

func (d DeathMeansID) String() string {
	st := DeathMeans[d]
	if st == "" {
		return "MOD_UNKNOWN"
	}
	return st
}
