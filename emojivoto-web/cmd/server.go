package main

import (
	"log"
	"os"

	"github.com/buoyantio/emojivoto/db"
	pb "github.com/buoyantio/emojivoto/emojivoto-web/gen/proto"
	"github.com/buoyantio/emojivoto/emojivoto-web/web"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
)

var (
	webPort              = os.Getenv("WEB_PORT")
	emojisvcHost         = os.Getenv("EMOJISVC_HOST")
	votingsvcHost        = os.Getenv("VOTINGSVC_HOST")
	indexBundle          = os.Getenv("INDEX_BUNDLE")
	webpackDevServerHost = os.Getenv("WEBPACK_DEV_SERVER")
)

func main() {

	if webPort == "" || emojisvcHost == "" || votingsvcHost == "" {
		log.Fatalf("WEB_PORT (currently [%s]) EMOJISVC_HOST (currently [%s]) and VOTINGSVC_HOST (currently [%s]) INDEX_BUNDLE (currently [%s]) environment variables must me set.", webPort, emojisvcHost, votingsvcHost, indexBundle)
	}

	votingSvcConn := openGrpcClientConnection(votingsvcHost)
	votingClient := pb.NewVotingServiceClient(votingSvcConn)
	defer votingSvcConn.Close()

	emojiSvcConn := openGrpcClientConnection(emojisvcHost)
	emojiSvcClient := pb.NewEmojiServiceClient(emojiSvcConn)
	defer emojiSvcConn.Close()

	database, err := db.NewClient()
	if err != nil {
		panic("failed to connect database")
	}

	web.StartServer(webPort, webpackDevServerHost, indexBundle, database, emojiSvcClient, votingClient)
}

func openGrpcClientConnection(host string) *grpc.ClientConn {
	log.Printf("Connecting to [%s]", host)
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}
