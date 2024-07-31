package internal

import (
	"github.com/rs/zerolog/log"
	"github.com/tebeka/atexit"
	"github.com/wintbiit/ors-proto"
	"github.com/wintbiit/ors-proto/proto"
	"time"
)

var client *ors.Client

func bootServer() {
	client = ors.
		NewClient(Config.RoboMaster.Address,
			73,
			proto.S1StuMainJudgeClientTId,
			proto.S1StuMainJudgeClientTeamId)

	client.WithLogger(newServerLogger("transport"))

	client.Debug = DEBUG

	client.WithAnyHandler(onRecvProto)

	client.WithHandler(proto.ProtoIDS1ProtoLoginAck, onRecvLoginAck)

	client.WithHandler(proto.ProtoIDS1ProtoHeartBeatAck, onRecvHeartBeatAck)

	err := client.Connect()
	if err != nil {
		log.Fatal().Msgf("client connect failed: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = client.Login(proto.S1StuMainJudgeLoginPass)

	if err != nil {
		log.Fatal().Msgf("login failed: %v", err)
	}

	atexit.Register(func() {
		client.Close()
	})
}

func onRecvProto(ctx *proto.S1ProtoContext) {
	go writeRecord(ctx)
}

func onRecvLoginAck(ctx *proto.S1ProtoContext) {
	var ack proto.S1ProtoLoginAck
	err := ack.Deserialize(ctx.Data)
	if err != nil {
		log.Error().Msgf("login ack deserialize failed: %v", err)
		return
	}

	if ack.ResultId == 1 {
		log.Info().Msg("login success")
	}

	if ack.ResultId > 1 {
		log.Error().Msgf("login failed: %d", ack.ResultId)
	}
}

func onRecvHeartBeatAck(ctx *proto.S1ProtoContext) {
	var ack proto.S1ProtoHeartBeatAck
	err := ack.Deserialize(ctx.Data)
	if err != nil {
		log.Error().Msgf("heartbeat ack deserialize failed: %v", err)
		return
	}
}
