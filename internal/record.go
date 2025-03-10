package internal

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog/log"
	"github.com/tebeka/atexit"
	"github.com/wintbiit/ors-proto/proto"
	"gorm.io/gorm"
)

var (
	recordDb  *gorm.DB
	recordChn chan *Record
)

type Record struct {
	ProtoID    uint16
	ProtoName  string
	DataLen    uint32
	ProtoType  uint16
	AckType    byte
	SequenceId byte
	Data       string

	*gorm.Model
}

func bootRecorder() {
	if !Config.RecordProto {
		return
	}

	var err error

	dsn := fmt.Sprintf("file:logs/record.db?cache=shared&mode=rwc")
	recordDb, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: &gormLogger{},
	})

	if err != nil {
		log.Fatal().Err(err).Msg("failed to open record db")
	}

	err = recordDb.AutoMigrate(&Record{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to migrate record db")
	}

	recordChn = make(chan *Record, 1000)

	atexit.Register(func() {
		close(recordChn)
	})

	go recordCoroutine()
}

func recordCoroutine() {
	for rec := range recordChn {
		err := recordDb.Create(rec).Error
		if err != nil {
			log.Error().Err(err).Uint16("proto id", rec.ProtoID).Msg("failed to record proto")
		}
	}
}

func writeRecord(ctx *proto.S1ProtoContext) {
	if !Config.RecordProto {
		return
	}

	protoName, ok := proto.ProtoIdMap[ctx.Header.ProtoId]
	if !ok {
		protoName = "unknown"
	}

	inst := proto.CreateProtoInstance(ctx.Header.ProtoId)

	instJ, err := sonic.ConfigFastest.MarshalToString(inst)
	if err != nil {
		log.Error().Err(err).Uint16("protoId", ctx.Header.ProtoId).Msg("failed write proto record")
	}

	recordChn <- &Record{
		ProtoID:    ctx.Header.ProtoId,
		ProtoName:  protoName,
		DataLen:    ctx.Header.DataLen,
		ProtoType:  ctx.Header.ProtoType,
		AckType:    ctx.Header.AckType,
		SequenceId: ctx.Header.SequenceId,
		Data:       instJ,
	}
}
