gopherfacts
===

Go sdk for [Artifacts MMO](https://artifactsmmo.com/client). Hand built with love.

Artifact MMO Client: https://artifactsmmo.com/client

## Getting Started

```
token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6InRoZXN0dWNrc3R0ZXIiLCJwYXNzd29yZF9jaGFuZ2VkIjoiIn0.g3qYaJOY40MEjUu2234JeXYcY-1HeboS7LmGQcHhnuk"
character := "MyCharacter"
client := clients.NewClient(&token)

err := client.MyCharacterClient.Move(character, 0, 0)
```