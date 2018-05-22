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
		MELEEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Warrior/Artist",
		MELEEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Warrior/Performer",
		MELEEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Warrior/Hawker",
		MELEEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Warrior/Orator",
		MELEEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Warrior/Hunter",
		MELEEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Warrior/Merchant",
		MELEEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Warrior/Adviser",
		MELEEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Warrior/Scholar",
		MELEEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Warrior/Engineer",
		MELEEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Paladin/Artist",
		MELEEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Paladin/Performer",
		MELEEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Paladin/Hawker",
		MELEEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Paladin/Orator",
		MELEEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Paladin/Hunter",
		MELEEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Paladin/Merchant",
		MELEEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Paladin/Adviser",
		MELEEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Paladin/Scholar",
		MELEEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Paladin/Engineer",
		MAGICPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Cleric/Artist",
		MAGICPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Cleric/Performer",
		MAGICPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Cleric/Hawker",
		MAGICPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Cleric/Orator",
		MAGICPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Cleric/Hunter",
		MAGICPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Cleric/Merchant",
		MAGICPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Cleric/Adviser",
		MAGICPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Cleric/Scholar",
		MAGICPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Cleric/Engineer",
		MAGICPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Mage/Artist",
		MAGICPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Mage/Performer",
		MAGICPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Mage/Hawker",
		MAGICPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Mage/Orator",
		MAGICPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Mage/Hunter",
		MAGICPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Mage/Merchant",
		MAGICPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Mage/Adviser",
		MAGICPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Mage/Scholar",
		MAGICPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Mage/Engineer",
		MAGICPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Warlock/Artist",
		MAGICPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Warlock/Performer",
		MAGICPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Warlock/Hawker",
		MAGICPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Warlock/Orator",
		MAGICPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Warlock/Hunter",
		MAGICPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Warlock/Merchant",
		MAGICPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Warlock/Adviser",
		MAGICPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Warlock/Scholar",
		MAGICPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Warlock/Engineer",
		RANGEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Caster/Artist",
		RANGEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Caster/Performer",
		RANGEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Caster/Hawker",
		RANGEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Caster/Orator",
		RANGEPRIMARY | MAGICSECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Caster/Hunter",
		RANGEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Caster/Merchant",
		RANGEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Caster/Adviser",
		RANGEPRIMARY | MAGICSECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Caster/Scholar",
		RANGEPRIMARY | MAGICSECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Caster/Engineer",
		RANGEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Sniper/Artist",
		RANGEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Sniper/Performer",
		RANGEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Sniper/Hawker",
		RANGEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Sniper/Orator",
		RANGEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Sniper/Hunter",
		RANGEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Sniper/Merchant",
		RANGEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Sniper/Adviser",
		RANGEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Sniper/Scholar",
		RANGEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Sniper/Engineer",
		RANGEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Archer/Artist",
		RANGEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Archer/Performer",
		RANGEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Archer/Hawker",
		RANGEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Archer/Orator",
		RANGEPRIMARY | MELEESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Archer/Hunter",
		RANGEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Archer/Merchant",
		RANGEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Archer/Adviser",
		RANGEPRIMARY | MELEESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Archer/Scholar",
		RANGEPRIMARY | MELEESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Archer/Engineer",
		MELEEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CREATIVESECONDARY: "Ranger/Artist",
		MELEEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | CUNNINGSECONDARY:  "Ranger/Performer",
		MELEEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CREATIVESECONDARY:  "Ranger/Hawker",
		MELEEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | CUNNINGSECONDARY:   "Ranger/Orator",
		MELEEPRIMARY | RANGESECONDARY | CUNNINGPRIMARY | ORDERLYSECONDARY:   "Ranger/Hunter",
		MELEEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CUNNINGSECONDARY:   "Ranger/Merchant",
		MELEEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | ORDERLYSECONDARY:   "Ranger/Adviser",
		MELEEPRIMARY | RANGESECONDARY | ORDERLYPRIMARY | CREATIVESECONDARY:  "Ranger/Scholar",
		MELEEPRIMARY | RANGESECONDARY | CREATIVEPRIMARY | ORDERLYSECONDARY:  "Ranger/Engineer"}
}
