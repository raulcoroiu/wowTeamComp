package models

type ApiResponse struct {
	Rankings       []Ranking `json:"rankings"`
	LeaderboardURL string    `json:"leaderboard_url"`
	Params         Params    `json:"params"`
}

type Params struct {
	Season  Season `json:"season"`
	Region  string `json:"region"`
	Dungeon string `json:"dungeon"`
	Page    int64  `json:"page"`
}

type Ranking struct {
	Rank  int64   `json:"rank"`
	Score float64 `json:"score"`
	Run   Run     `json:"run"`
}

type Run struct {
	Season             Season           `json:"season"`
	Status             Status           `json:"status"`
	Dungeon            Dungeon          `json:"dungeon"`
	KeystoneRunID      int64            `json:"keystone_run_id"`
	MythicLevel        int64            `json:"mythic_level"`
	ClearTimeMS        int64            `json:"clear_time_ms"`
	KeystoneTimeMS     int64            `json:"keystone_time_ms"`
	CompletedAt        string           `json:"completed_at"`
	NumChests          int64            `json:"num_chests"`
	TimeRemainingMS    int64            `json:"time_remaining_ms"`
	LoggedRunID        *int64           `json:"logged_run_id"`
	Videos             []interface{}    `json:"videos"`
	WeeklyModifiers    []WeeklyModifier `json:"weekly_modifiers"`
	NumModifiersActive int64            `json:"num_modifiers_active"`
	Faction            RunFaction       `json:"faction"`
	DeletedAt          interface{}      `json:"deleted_at"`
	KeystonePlatoonID  interface{}      `json:"keystone_platoon_id"`
	KeystoneTeamID     int64            `json:"keystone_team_id"`
	Roster             []Roster         `json:"roster"`
	Platoon            interface{}      `json:"platoon"`
}

type Dungeon struct {
	ID                     int64            `json:"id"`
	Name                   DungeonName      `json:"name"`
	ShortName              DungeonShortName `json:"short_name"`
	Slug                   DungeonSlug      `json:"slug"`
	ExpansionID            int64            `json:"expansion_id"`
	IconURL                string           `json:"icon_url"`
	Patch                  Patch            `json:"patch"`
	KeystoneTimerMS        int64            `json:"keystone_timer_ms"`
	NumBosses              int64            `json:"num_bosses"`
	GroupFinderActivityIDS []int64          `json:"group_finder_activity_ids"`
}

type Roster struct {
	Character    Character   `json:"character"`
	OldCharacter interface{} `json:"oldCharacter"`
	IsTransfer   bool        `json:"isTransfer"`
	Role         Role        `json:"role"`
}

type Character struct {
	ID                  int64         `json:"id"`
	PersonaID           int64         `json:"persona_id"`
	Name                string        `json:"name"`
	Class               Class         `json:"class"`
	Race                Class         `json:"race"`
	Faction             ClassFaction  `json:"faction"`
	Level               int64         `json:"level"`
	Spec                Class         `json:"spec"`
	Path                string        `json:"path"`
	Realm               Realm         `json:"realm"`
	Region              Region        `json:"region"`
	Stream              *Stream       `json:"stream"`
	RecruitmentProfiles []interface{} `json:"recruitmentProfiles"`
}

type Class struct {
	ID      int64         `json:"id"`
	Name    string        `json:"name"`
	Slug    string        `json:"slug"`
	Faction *ClassFaction `json:"faction,omitempty"`
}

type Realm struct {
	ID               int64    `json:"id"`
	ConnectedRealmID int64    `json:"connectedRealmId"`
	Name             string   `json:"name"`
	AltName          *AltName `json:"altName"`
	Slug             string   `json:"slug"`
	AltSlug          string   `json:"altSlug"`
	Locale           Locale   `json:"locale"`
	IsConnected      bool     `json:"isConnected"`
}

type Region struct {
	Name      RegionName      `json:"name"`
	Slug      RegionSlug      `json:"slug"`
	ShortName RegionShortName `json:"short_name"`
}

type Stream struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	UserID       string        `json:"user_id"`
	GameID       string        `json:"game_id"`
	Type         string        `json:"type"`
	Title        string        `json:"title"`
	CommunityIDS []interface{} `json:"community_ids"`
	ViewerCount  int64         `json:"viewer_count"`
	StartedAt    string        `json:"started_at"`
	Language     string        `json:"language"`
	ThumbnailURL string        `json:"thumbnail_url"`
}

type WeeklyModifier struct {
	ID          int64  `json:"id"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Season string

const (
	SeasonDf1 Season = "season-df-1"
)

type DungeonName string

const (
	CourtOfStars            DungeonName = "Court of Stars"
	ShadowmoonBurialGrounds DungeonName = "Shadowmoon Burial Grounds"
)

type Patch string

const (
	The601 Patch = "6.0.1"
	The70  Patch = "7.0"
)

type DungeonShortName string

const (
	Cos DungeonShortName = "COS"
	Sbg DungeonShortName = "SBG"
)

type DungeonSlug string

const (
	SlugCourtOfStars            DungeonSlug = "court-of-stars"
	SlugShadowmoonBurialGrounds DungeonSlug = "shadowmoon-burial-grounds"
)

type RunFaction string

const (
	Mixed          RunFaction = "mixed"
	PurpleAlliance RunFaction = "alliance"
)

type ClassFaction string

const (
	FluffyAlliance ClassFaction = "alliance"
	Horde          ClassFaction = "horde"
)

type AltName string

const (
	РевущийФьорд AltName = "Ревущий фьорд"
	尖石           AltName = "尖石"
	阿薩斯          AltName = "阿薩斯"
)

type Locale string

const (
	DeDE Locale = "de_DE"
	EnGB Locale = "en_GB"
	EnUS Locale = "en_US"
	RuRU Locale = "ru_RU"
	ZhTW Locale = "zh_TW"
)

type RegionName string

const (
	Europe              RegionName = "Europe"
	Taiwan              RegionName = "Taiwan"
	UnitedStatesOceania RegionName = "United States & Oceania"
)

type RegionShortName string

const (
	Eu RegionShortName = "EU"
	Tw RegionShortName = "TW"
	Us RegionShortName = "US"
)

type RegionSlug string

const (
	SlugEu RegionSlug = "eu"
	SlugTw RegionSlug = "tw"
	SlugUs RegionSlug = "us"
)

type Role string

const (
	Dps    Role = "dps"
	Healer Role = "healer"
	Tank   Role = "tank"
)

type Status string

const (
	Finished Status = "finished"
)