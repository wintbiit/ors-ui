package internal

import (
	"log"

	"github.com/tebeka/atexit"
	"github.com/wintbiit/ors-proto"
	"github.com/wintbiit/ors-proto/proto"
)

var server *ors.Server

func bootServer() {
	server = ors.NewServer(config.RoboMaster.Address, proto.S1StuMainJudgeClientId+1, proto.S1StuMainJudgeClientTId, proto.S1StuMainJudgeClientTeamId)

	err := server.Connect()
	if err != nil {
		log.Fatalf("server connect failed: %v", err)
	}

	err = server.Login(proto.S1StuMainJudgeLoginPass)

	if err != nil {
		log.Fatalf("login failed: %v", err)
	}

	atexit.Register(func() {
		server.Close()
	})
}
