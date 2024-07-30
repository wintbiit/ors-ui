package internal

import (
	"github.com/rs/zerolog/log"
	"github.com/tebeka/atexit"
	"github.com/wintbiit/ors-proto"
	"github.com/wintbiit/ors-proto/proto"
)

var server *ors.Server

func bootServer() {
	server = ors.
		NewServer(Config.RoboMaster.Address,
			73,
			proto.S1StuMainJudgeClientTId,
			proto.S1StuMainJudgeClientTeamId)

	server.WithLogger(newServerLogger("transport"))

	server.Debug = DEBUG

	server.WithAnyHandler(onRecvProto)

	server.WithHandler(proto.ProtoIDS1ProtoLoginAck, onRecvLoginAck)

	server.WithHandler(proto.ProtoIDS1ProtoHeartBeatAck, onRecvHeartBeatAck)

	err := server.Connect()
	if err != nil {
		log.Fatal().Msgf("server connect failed: %v", err)
	}

	err = server.Login(proto.S1StuMainJudgeLoginPass)

	if err != nil {
		log.Fatal().Msgf("login failed: %v", err)
	}

	atexit.Register(func() {
		server.Close()
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
