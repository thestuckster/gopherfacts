gopherfacts
===

Go sdk for [Artifacts MMO](https://artifactsmmo.com). Hand built with love.

Artifact MMO Client: https://artifactsmmo.com/client

## Getting Started

```
token := "MyToken"
character := "MyCharacter"
client := clients.NewClient(&token)

err := client.MyCharacterClient.Move(character, 0, 0)
```


## APIs 

```
🚧 Server Status
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

🚧 Equip Item
🚧 Unequip Item
🚧 Deposit Bank Gold
🚧 Recycle
🚧 Withdraw Bank Gold
🚧 Grand Exchange Sell
🚧 Accept task
🚧 Complete Task
🚧 Task Exchange
🚧 Delete Item
```

### Account

```
🚧 Change Password
🚧 Create Character
🚧 Create Account
🚧 Create Token
```

### Meta / Character Info

```
✅ Get All Characters Info
✅ Get single character info
✅ Get Bank Gold

🚧 Get All Characters Logs
🚧 Get Specific Character Info
🚧 Get Map
🚧 Get All Map
🚧 Get Item
🚧 Get All Items
🚧 ... same for monsters, resources, events, GE, 
```