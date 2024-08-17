gopherfacts
===

Go sdk for [Artifacts MMO](https://artifactsmmo.com/client). Hand built with love.

Artifact MMO Client: https://artifactsmmo.com/client

## Getting Started

```
token := "MyToken"
character := "MyCharacter"
client := clients.NewClient(&token)

err := client.MyCharacterClient.Move(character, 0, 0)
```