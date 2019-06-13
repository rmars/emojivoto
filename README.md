# Emojivoto, for the purposes of exploring GraphQL

This demo is forked from https://github.com/BuoyantIO/emojivoto, and has a
couple more data sources so that we can play around with stitching together data
using GraphQL. I've broken down some basic tasks into separate PRs as tutorials,
which you can [browse in the /pulls
section](https://github.com/rmars/emojivoto/pulls?utf8=%E2%9C%93&q=is%3Apr+is%3Aclosed+tutorial).

The emojivoto application is a microservice application that allows users to
vote for their favorite emoji, and tracks votes received on a leaderboard. May
the best emoji win.

The emojivoto application is composed of the following 3 services:

* [emojivoto-web](emojivoto-web/): Web frontend and REST API
* [emojivoto-emoji-svc](emojivoto-emoji-svc/): gRPC API for finding and listing emoji
* [emojivoto-voting-svc](emojivoto-voting-svc/): gRPC API for voting and leaderboard

![Emojivoto Topology](assets/emojivoto-topology.png "Emojivoto Topology")

In this fork, we've added a sqlite database of users and their favourite emoji,
as well as a web API that fetches additional data about users.

## Checking out and running this code

Because the go packages are imported under the fork's paths, you'll need to check
out the parent repo and point your git remote to this repo in order to run the go code:

```
cd $GOPATH # make sure you're in your gopath
mkdir -p buoyantio
cd buoyantio
git clone https://github.com/BuoyantIO/emojivoto.git
cd emojivoto
git remote add rmarsfork https://github.com/rmars/emojivoto.git
git pull rmarsfork master
git remote -v # you should see the fork's url here
```

## Running
<!---
### In docker-compose
 docker-compose won't work until I add sqlite3 to the docker file

It's possible to run the app with docker-compose.

Build and run:

```
make deploy-to-docker-compose
```

The web app will be running on port 8080 of your docker host.

If you've changed code and want to rebuild the docker images:
```
make build-base-docker-image # build base docker image
make build # build docker images
```
-->

### Emojivoto webapp

This app is written with React and bundled with webpack.
Use the following to run the emojivoto go services and develop on the frontend.

Set up proto files, build apps
```
make build # assumes you have dep installed
```

Start the voting service
```
GRPC_PORT=8081 go run emojivoto-voting-svc/cmd/server.go
```

[In a separate terminal window] Start the emoji service
```
GRPC_PORT=8082 go run emojivoto-emoji-svc/cmd/server.go
```

[In a separate terminal window] Bundle the frontend assets
```
cd emojivoto-web/webapp
yarn install
yarn webpack # one time asset-bundling OR
yarn webpack-dev-server --port 8083 # bundle/serve reloading assets
```

[In a separate terminal window] Start the web service
```
export WEB_PORT=8080
export VOTINGSVC_HOST=localhost:8081
export EMOJISVC_HOST=localhost:8082

# if you ran yarn webpack
export INDEX_BUNDLE=emojivoto-web/webapp/dist/index_bundle.js

# if you ran yarn webpack-dev-server
export WEBPACK_DEV_SERVER=http://localhost:8083

# start the webserver
go run emojivoto-web/cmd/server.go
```

[Optional] Start the vote bot for automatic traffic generation.
```
export WEB_HOST=localhost:8080
go run emojivoto-web/cmd/vote-bot/main.go
```

View emojivoto
```
open http://localhost:8080
```

### Generating some traffic to the emojivoto app

The `VoteBot` service can generate some traffic for you. It votes on emoji
"randomly" as follows:
- It votes for :doughnut: 15% of the time.
- When not voting for :doughnut:, it picks an emoji at random

If you'd like to run the bot:
```
export WEB_HOST=localhost:8080 # replace with your web location
go run emojivoto-web/cmd/vote-bot/main.go
```

# Resources

To play around more with a fully fledged GraphQL API, see
https://graphql.github.io/swapi-graphql/

More resources/tutorials:
- StarWars API playground: https://graphql.github.io/swapi-graphql
- Official Docs: https://graphql.org/learn/
- Tutorial website: https://www.howtographql.com/
- https://github.com/MoonHighway/learning-graphql
- GraphQL libraries: https://graphql.org/code/
- https://github.com/graph-gophers/graphql-go
- https://medium.com/open-graphql/choosing-a-graphql-server-library-in-go-8836f893881b
