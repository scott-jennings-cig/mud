package mud

var strengthClassNameMap map[byte]string
var skillClassNameMap map[byte]string
var fullTitleMap map[byte]string

// Strengths
const (
	MELEESECONDARY = byte(1)
	RANGESECONDARY = byte(2)
	MAGICSECONDARY = byte(3)
	MELEEPRIMARY   = byte(4)
	RANGEPRIMARY   = byte(8)
	MAGICPRIMARY   = byte(12)
)

// Skills
const (
	CUNNINGSECONDARY  = byte(16)
	ORDERLYSECONDARY  = byte(32)
	CREATIVESECONDARY = byte(48)
	CUNNINGPRIMARY    = byte(64)
	ORDERLYPRIMARY    = byte(128)
	CREATIVEPRIMARY   = byte(192)
)

// Masks for strengths/skills
const (
	SECONDARYSTRENGTHMASK = byte(3)
	PRIMARYSTRENGTHMASK   = byte(12)
	SECONDARYSKILLMASK    = byte(48)
	PRIMARYSKILLMASK      = byte(192)
)

// ClassInfo handles user/NPC class orientation
type ClassInfo interface {
	ClassInfo() byte
	SetClassInfo(byte)

	Strengths() (byte, byte)
	SetStrengths(byte, byte)
	Skills() (byte, byte)
	SetSkills(byte, byte)
}

// GetSubTitles takes strengths and gives a class title
func GetSubTitles(strengthPrimary, strengthSecondary, skillPrimary, skillSecondary byte) (string, string) {
	strS, sklS := "Hippopotamus", "Spaghetti"

	strK, sK := strengthClassNameMap[strengthPrimary|strengthSecondary]

	if sK {
		strS = strK
	}

	sklK, lK := skillClassNameMap[skillPrimary|skillSecondary]

	if lK {
		sklS = sklK
	}

	return strS, sklS
}

// GetTitle takes strengths and gives a class title
func GetTitle(strengthPrimary, strengthSecondary, skillPrimary, skillSecondary byte) string {
	title := "Unworthy"

	newTitle, ok := fullTitleMap[strengthPrimary|strengthSecondary|skillPrimary|skillSecondary]

	if ok {
		title = newTitle
	}

	return title
}

func init() {
	strengthClassNameMap = map[byte]string{
		MELEEPRIMARY | MELEESECONDARY: "Warrior",
		MELEEPRIMARY | MAGICSECONDARY: "Paladin",
		MAGICPRIMARY | MELEESECONDARY: "Cleric",
		MAGICPRIMARY | MAGICSECONDARY: "Mage",
		MAGICPRIMARY | RANGESECONDARY: "Warlock",
		RANGEPRIMARY | MAGICSECONDARY: "Caster",
		RANGEPRIMARY | RANGESECONDARY: "Sniper",
		RANGEPRIMARY | MELEESECONDARY: "Archer",
		MELEEPRIMARY | RANGESECONDARY: "Ranger"}

	skillClassNameMap = map[byte]string{
		CREATIVEPRIMARY | CREATIVESECONDARY: "Artist",
		CREATIVEPRIMARY | CUNNINGSECONDARY:  "Performer",
		CUNNINGPRIMARY | CREATIVESECONDARY:  "Hawker",
		CUNNINGPRIMARY | CUNNINGSECONDARY:   "Orator",
		CUNNINGPRIMARY | ORDERLYSECONDARY:   "Hunter",
		ORDERLYPRIMARY | CUNNINGSECONDARY:   "Merchant",
		ORDERLYPRIMARY | ORDERLYSECONDARY:   "Adviser",
		ORDERLYPRIMARY | CREATIVESECONDARY:  "Scholar",
		CREATIVEPRIMARY | ORDERLYSECONDARY:  "Engineer"}

	fullTitleMap = map[byte]string{
		MELEEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Martial Artist",
		MELEEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Swordslinger",
		MELEEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Swordsknecht",
		MELEEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Strategist",
		MELEEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Tracker",
		MELEEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Pilgrim",
		MELEEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Monk",
		MELEEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Warrior-Poet",
		MELEEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Warrior Sapper",
		MELEEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Paladin-Poet",
		MELEEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Wind Paladin",
		MELEEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Paladin of the Crooks",
		MELEEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Great Paladin",
		MELEEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Paladin of the Wilds",
		MELEEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Paladin of the Bazaar",
		MELEEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Paladin of Wisdom",
		MELEEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Paladin of History",
		MELEEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Paladin-Toolsmith",
		MAGICPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Cleric in Eight-Arms",
		MAGICPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Fox-Cleric",
		MAGICPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Wolf-Cleric",
		MAGICPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Coyote-in-Cloth",
		MAGICPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Ravenscloth",
		MAGICPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Hippocleric of the Trail",
		MAGICPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Spinner-Cleric",
		MAGICPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Cleric of Dancing Flames",
		MAGICPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Eagle-Cleric",
		MAGICPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Weaver Mage",
		MAGICPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Storyteller Mage",
		MAGICPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Mindthief Mage",
		MAGICPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Mage of Flows",
		MAGICPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Mageseeker",
		MAGICPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Mage of Fortune",
		MAGICPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Seer",
		MAGICPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Master Mage",
		MAGICPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Mage of the Flame",
		MAGICPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Mist Warlock",
		MAGICPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Warlock-Enthraller",
		MAGICPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Flickering Warlock",
		MAGICPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Shifting Warlock",
		MAGICPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Warlock of Waters",
		MAGICPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Warlock of Fortune",
		MAGICPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Warlock of the Stone",
		MAGICPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Warlock Dilettante",
		MAGICPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Warlock of Flame",
		RANGEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Acolyte of Clouds",
		RANGEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Acolyte of the Winds",
		RANGEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Acolyte of the Liminal",
		RANGEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Disciple of the Shifing Sands",
		RANGEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Disciple of the Waters",
		RANGEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Disciple of the Trail",
		RANGEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Acolyte of Stone",
		RANGEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Friend of the Sprite",
		RANGEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Acolyte of Flame",
		RANGEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Cloaked-in-Mist",
		RANGEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Fox-in-Wood",
		RANGEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Wolf-in-Shadows",
		RANGEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Trapsetter",
		RANGEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Raven-in-Air",
		RANGEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Spy",
		RANGEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Stone-in-Mountain",
		RANGEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Owl-eyed",
		RANGEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Demolitionist",
		RANGEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Cloud Archer",
		RANGEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Bow Adventurer",
		RANGEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Kastspeerknecht",
		RANGEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Assassin",
		RANGEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Trapper",
		RANGEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Archer of Fortune",
		RANGEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Skill Shot",
		RANGEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Bear Ranger",
		RANGEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Eagle-eyed Ranger",
		MELEEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Vagabond",
		MELEEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Fox-Rogue",
		MELEEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Spearknecht",
		MELEEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Ranger, Acolyte of the Flow",
		MELEEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Ranger, Warden of the Hunt",
		MELEEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Ranger of Fortune",
		MELEEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Scout",
		MELEEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Will-o-Wisp Ranger",
		MELEEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Ranger, Warden of Stone"}
}
