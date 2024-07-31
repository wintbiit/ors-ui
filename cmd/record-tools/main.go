package main

import "flag"

func main() {
	var runAnalyze bool
	var analyzeRecordPath string
	var analyzeExportFile string

	flag.BoolVar(&runAnalyze, "analyze", false, "Run analyze record")
	flag.StringVar(&analyzeRecordPath, "record", "", "Record path")
	flag.StringVar(&analyzeExportFile, "export", "", "Export analyze result to file")

	flag.Parse()

	if runAnalyze {
		commandAnalyzeRecord(&AnalyzeOptions{
			RecordPath: analyzeRecordPath,
			ExportFile: analyzeExportFile,
		})

		return
	}
}
