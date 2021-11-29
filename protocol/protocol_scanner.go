package protocol

import (
	"github.com/Weldeasy/onionscan/config"
	"github.com/Weldeasy/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
