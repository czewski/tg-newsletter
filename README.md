
# tg-newsletter &middot; [![Go](https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/) [![Mongo](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)](https://www.mongodb.com/pt-br) [![TG](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://core.telegram.org/bots/api) 

This is an Telegram BOT that send news based on custom queries that were added to it.

I created this project with learning purposes. 

## Installing / Getting started

There's no fancy installation, just create a shell script and run.

```shell
go run main.go
```

When executing, it will search for new messages with Telegram API. 

## How the bot works
- 

## Developing

### Built With
- [Go 1.17.6](https://go.dev/)
- [MongoDB Driver](go.mongodb.org/mongo-driver/mongo)

## Configuration

For this bot to work, you need to set a config/config.json file with the following struct: 

```json
{
"botKey": "",
"newsKey": "",
"mongoPwd": ""
}
```

Where:
- "botKey" represents the TelegramAPI bot access key.
- "newsKey" represents the NewsAPI access key.
- "mongoPwd" represents the MongoDB password. 


## Api Reference

- [Telegram API](https://core.telegram.org/bots/api)
- [News API](https://newsapi.org/docs)


## Database

For this projected I wanted to learn a bit about [MongoDB](https://www.mongodb.com/pt-br) and non-relational databases.  
