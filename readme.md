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

## Usage

Create a new user:

```bash
bootdev-gator register <name>
```

Add a feed:

```bash
bootdev-gator addfeed <url>
```

Start the aggregator:

```bash
bootdev-gator agg 30s
```

View the posts:

```bash
bootdev-gator browse [limit]
```

There are a few other commands you'll need as well:

- `bootdev-gator login <name>` - Log in as a user that already exists
- `bootdev-gator users` - List all users
- `bootdev-gator feeds` - List all feeds
- `bootdev-gator follow <url>` - Follow a feed that already exists in the database
- `bootdev-gator unfollow <url>` - Unfollow a feed that already exists in the database
