package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/wintbiit/ors-ui/internal"
	"gorm.io/gorm"
	"os"
	"strings"
)

type AnalyzeOptions struct {
	RecordPath string
	ExportFile string
}

type AnalyzeResult struct {
	ProtoID      uint16
	ProtoName    string
	Count        int
	TotalBytes   int
	AvgBytes     int
	AvgFrequency float64
}

func commandAnalyzeRecord(opt *AnalyzeOptions) {
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", opt.RecordPath)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open record db")
		return
	}

	var records []internal.Record
	db.Find(&records)

	result := make(map[uint16]*AnalyzeResult)
	startTimes := make(map[uint16]int64)
	endTimes := make(map[uint16]int64)

	for _, record := range records {
		if _, ok := result[record.ProtoID]; !ok {
			result[record.ProtoID] = &AnalyzeResult{
				ProtoID:   record.ProtoID,
				ProtoName: record.ProtoName,
			}
		}

		result[record.ProtoID].Count++
		result[record.ProtoID].TotalBytes += len(record.Data)

		if _, ok := startTimes[record.ProtoID]; !ok {
			startTimes[record.ProtoID] = record.CreatedAt.UnixNano()
		}

		endTimes[record.ProtoID] = record.CreatedAt.UnixNano()
	}

	for _, r := range result {
		r.AvgBytes = r.TotalBytes / r.Count
		if r.Count <= 1 {
			r.AvgFrequency = 0
		} else {
			r.AvgFrequency = float64(r.Count) / float64((endTimes[r.ProtoID]-startTimes[r.ProtoID])/1e9)
		}
	}

	fmt.Printf("ProtoID\tProtoName\tCount\tTotalBytes(bytes)\tAvgBytes(bytes)\tAvgFrequency(Hz)\n")
	for _, r := range result {
		fmt.Printf("%d\t%s\t%d\t%d\t%d\t%.2f\n", r.ProtoID, r.ProtoName, r.Count, r.TotalBytes, r.AvgBytes, r.AvgFrequency)
	}

	if opt.ExportFile != "" && strings.HasSuffix(opt.ExportFile, ".csv") {
		fmt.Printf("Exporting to %s\n", opt.ExportFile)

		f, err := os.Create(opt.ExportFile)
		if err != nil {
			fmt.Println("Failed to create export file")
			return
		}

		defer f.Close()

		f.WriteString("ProtoID,ProtoName,Count,TotalBytes,AvgBytes,AvgFrequency\n")
		for _, r := range result {
			f.WriteString(fmt.Sprintf("%d,%s,%d,%d,%d,%.2f\n", r.ProtoID, r.ProtoName, r.Count, r.TotalBytes, r.AvgBytes, r.AvgFrequency))
		}

		fmt.Println("Exported")
	}
}
