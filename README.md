# Gator

Gator is a simple command-line RSS feed aggregator written in Go. It lets you register users, follow RSS feeds, fetch posts, and browse them from your terminal using a PostgreSQL database.

---

## Requirements

To run Gator, you need the following installed:

- **Go** (https://go.dev/doc/install)
- **PostgreSQL** (https://www.postgresql.org/download/)

Make sure PostgreSQL is running and that you can connect to it locally.

---

## Installation

Install the `gator` CLI using `go install`:

```bash
go install github.com/SpollaL/gator@latest
````

Ensure your Go bin directory is in your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

After this, you should be able to run:

```bash
gator
```

---

## Configuration

Gator uses a config file located at:

```bash
~/.gatorconfig.json
```

Create the file:

```bash
touch ~/.gatorconfig.json
```

Example contents:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username`, `password`, and database name with your own values.

---

## Database Setup

Create a PostgreSQL database:

```bash
createdb gator
```

Run migrations if needed (using goose):

```bash
goose postgres "$DB_URL" up
```

---

## Usage

### Register and login

```bash
gator register alice
gator login alice
```

### Add and follow feeds

```bash
gator addfeed https://example.com/rss
gator follow https://example.com/rss
```

### Fetch posts

Fetch new posts on a schedule (for example, every 5 minutes):

```bash
gator agg 5m
```

### Browse posts

```bash
gator browse
gator browse 10
```

(The default limit is 2 posts.)

---

## Summary

* Install Go and PostgreSQL
* Install Gator with `go install`
* Set up `~/.gatorconfig.json`
* Register a user, add feeds, fetch posts, and browse them from the CLI

```
