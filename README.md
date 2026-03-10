# gator

## dev dependecies

For application we need to:

- posgresql
- go
- sqlc
- goose

All dependecies included in `devbox.json`. You simply install [Create a Dev Environment with Devbox - Jetify Docs](https://www.jetify.com/docs/devbox/quickstart)

## configuration

Gator used own configuration where you should setup connection to posgresql:

```json
{
  "db_url":"postgres://vit:password@localhost:5432/gator?sslmode=disable",
  "current_user_name":"test"
}
```

Current user will be setup by gator

## usage

`gator register user_name` - register user in database
`gator addfeed "feed_name" "feed_url"` - add rss feed to database
`gator follow "feed_url"` - follow existing feed
