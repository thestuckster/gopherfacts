package clients

const BASE_URL = "https://api.artifactsmmo.com/"
const BANK = BASE_URL + "my/bank"
const BANK_GOLD = BANK + "/gold"

const CHARACTER = BASE_URL + "my/%s/"
const ACTION = CHARACTER + "action/"
const MOVE = ACTION + "move/"
const FIGHT = ACTION + "fight/"
const GATHER = ACTION + "gathering"
const CRAFT = ACTION + "crafting"
const CHARACTER_BANK = ACTION + "bank/"
const DEPOSIT_CHARACTER_BANK = CHARACTER_BANK + "deposit"

const CHARACTER_INFO = BASE_URL + "characters/%s"

const ITEM_INFO = BASE_URL + "items"
const ALL_MAPS = BASE_URL + "maps"
