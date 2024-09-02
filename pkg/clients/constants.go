package clients

const BASE_URL = "https://api.artifactsmmo.com/"
const BANK = BASE_URL + "my/bank"
const BANK_GOLD = BANK + "/gold"
const BANK_EXPANSION = BANK + "/buy_expansion"

const CHARACTER = BASE_URL + "my/%s/"
const ACTION = CHARACTER + "action/"
const MOVE = ACTION + "move/"
const FIGHT = ACTION + "fight/"
const GATHER = ACTION + "gathering"
const CRAFT = ACTION + "crafting"
const EQUIP = ACTION + "equip"
const UNEQUIP = ACTION + "unequip"
const DELETE_ITEM = ACTION + "delete"
const RECYCLE = ACTION + "recycle"

const TASK = ACTION + "task/"
const NEW_TASK = TASK + "new"

const CHARACTER_BANK = ACTION + "bank/"
const DEPOSIT_CHARACTER_BANK = CHARACTER_BANK + "deposit"
const WITHDRAW_CHARACTER_BANK = CHARACTER_BANK + "withdraw"
const DEPOSIT_GOLD_BANK = DEPOSIT_CHARACTER_BANK + "/gold"
const WITHDRAW_GOLD_BANK = WITHDRAW_CHARACTER_BANK + "/gold"

const GE = ACTION + "ge/"
const GE_BUY = GE + "buy"
const GE_SELL = GE + "sell"

const CHARACTER_INFO = BASE_URL + "characters/%s"

const ITEM_INFO = BASE_URL + "items"
const ALL_MAPS = BASE_URL + "maps"

// equipment slots
const WEAPON = "weapon"
const SHIELD = "shield"
const HELMET = "helmet"
const BODY = "body_armor"
const LEG = "leg_armor"
const BOOTS = "boots"
const RING = "ring1"
const RING2 = "ring2"
const AMULET = "amulet"
const ARTIFACT1 = "artifact1"
const ARTIFACT2 = "artifact2"
const ARTIFACT3 = "artifact3"
const CONSUMABLE = "consumable1"
const CONSUMABLE2 = "consumable2"
