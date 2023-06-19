<h1 align=center>
⚠️THIS IS ONLY FOR MY LEARNING PURPOSE⚠️

Any contribution to this projects are welcomed.

# UrlShortener

## Are you tired of copying long url? Here is your solution!

### What is url shortener you asking 🤔?

- Copy your very long and ugly URL and shortener it beatiful short URL.

- For example: from https://verylongurlthatisnotbeatifull.com/who/even/can/remember/it/what/is/it?long=1&ugly=1 to http://localhost:4000/URL/shortAndBeatifull

## Build Using

### Backend
- Golang
- PostqreSQL
### Frontend
- Bootstrap

## about project

- This project is heavily inspired by book [Lets go](https://lets-go.alexedwards.net/). 
- I tried to using only standard libraries only.

# Install
### requriements:
- Golang 1.20+
- Docker or PostqreSQL
```bash
docker-compose -f services.yml up -d

#running migration on db
docker exec -it db_shorten psql -U dev
#run inside script from ./migration/migration_XXXX.sql

go mod download
go run ./cmd/api 
```

# TODO
- [ ] Lock urls by passwords
- [ ] Unit testing
