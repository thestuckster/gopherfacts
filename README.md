gopherfacts
===

Go sdk for [Artifacts MMO](https://artifactsmmo.com). Hand built with love.

Artifact MMO Client: https://artifactsmmo.com/client

## Getting Started

**Install gopherfacts**
```bash
go get github.com/thestuckster/gopherfacts
```

**Using gopherfacts**
```go
import "github.com/thestuckster/gopherfacts/pkg/clients"

func main() {

    //check the artifacts server status
    statusInfo, err := sdk.CheckServerStatus()
    if err != nil {
        logger.Error().Err(err).Msg("Artifacts Server isnt healthy, check back later.")
        os.Exit(1)
    }
    
    token := "YOUR_TOKEN"
    character := "YOUR_CHARACTER"
    
    // create the client
    gopherfacts := clients.NewClient(&token)
    
    //Character client is the basic functionality client.
    res, err := gopherfacts.CharacterClient.Move(character, 1, 1)
    if err != nil {
        panic(err)
    }
	//CharacterClient will NOT handle cooldowns, you must manage them yourself.
    time.Sleep(time.Duration(res.Cooldown.RemainingSeconds) * time.Second)
    
    //EasyClient provides easier to use functionality and WILL handle the cooldowns before returning its result
    res, err := gopherfacts.EasyClient.MoveToBank(character)
    if err != nil {
        panic(err)
    }
}
```

## Implemented Artifacts APIs

```
✅ Server Status
```

### Actions

```
✅ Moving
✅ Fighting
✅ Gathering
✅ Crafting
✅ Bank Deposit
✅ Withdraw Bank Item
✅ Grand Exchange Buy
✅ Grand Exchange Sell
✅ Equip Item
✅ Unequip Item
✅ Delete Item
✅ Deposit Bank Gold
✅ Withdraw Bank Gold
✅ Recycle
✅ Accept task
✅ Complete Task
✅ Task Cancel
✅ Task Exchange
✅ Delete Item
```

### Account

```
✅ Create Account
✅ Get Account Info
🚧 Change Password
🚧 Create Character
🚧 Create Token
```

### Meta / Character Info

```
✅ Get All Characters Info
✅ Get single character info
✅ Get Bank Gold
✅ Buy bank expansion
✅ Get All Map
✅ Get All Items
✅ Get All Monsters

🚧 Get All Characters Logs
🚧 Get Specific Character Info
🚧 Get Specific Map Tile
🚧 Get Specific Item
🚧 ... same for monsters, resources, events, GE, 
```