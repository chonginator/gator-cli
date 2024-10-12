# Gator Blog Aggregator CLI

Gator is a CLI tool for aggregating blog posts built using Go, PostgreSQL, Goose, and sqlc.

Gator streamlines the process of staying up-to-date with your favourite blogs. With Gator, you can add RSS feeds, follow or unfollow feeds based on your interests, list all of your followed feeds, and browse their latest posts. Gator also gives you the ability to fetch the latest posts in a long-running background service, so that you don't have to manually update or refresh feeds.

## Installation

### 1. Install Go 1.23 or later

Gator is written in Go, so you'll need to download and install Go to build Gator from source. You can download and install Go from [the official Go website](https://golang.org/dl/). 

### 2. Install PostgreSQL 14 or later

Gator uses a PostgreSQL database to store posts, so you'll need to download and install it. You can download and install PostgreSQL from [the official PostgreSQL website](https://www.postgresql.org/download/).

### 3. Install the Gator CLI

This command will download, build, and install the Gator CLI in your Go toolchain's `bin` directory:

```bash
go install github.com/chonginator/gator-cli
```

### 4. Set up the config file

After installation, you'll need to set up a configuration file for Gator. Create a file named `gatorconfig.json` in your home directory with the database connection URL in the format:

```json
{"db_url": "postgres://username:password@host:port/database?sslmode=disable"}
```

For example:
```json
{"db_url": "postgres://chonginator@localhost:5432/gator?sslmode=disable"}
```

### 5. Make sure that the PostgreSQL database server is running:

```bash
# For macOS
brew services start postgresql
```

```bash
# For Linux
sudo service postgresql start
```

For more detailed instructions, refer to the [official PostgreSQL documentation on server startup](https://www.postgresql.org/docs/current/server-start.html).

## Usage
Here are some commands you can run with Gator:
- `gator register <name>`: Register a new user
- `gator addfeed <name> <feed_url>`: Add a new feed
- `gator feeds`: List all the feeds added to Gator
- `gator follow <feed_url>`: Follow a feed
- `gator agg <time_between_reqs>`: Starts a long-running background service to fetch the latest posts at a set interval
- `gator browse [limit]`: View all the latest posts from the feeds the current user follows