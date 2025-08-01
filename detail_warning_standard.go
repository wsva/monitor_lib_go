package mlib

import "time"

//warning standard
const (
	WarnStdCPU         = 85
	WarnStdMemory      = 85
	WarnStdSwap        = 30
	WarnStdDisk        = 85
	WarnStdNTPOffset   = 10 * time.Second //时间差超过10秒
	WarnStdASM         = 85
	WarnStdArchiveLog  = 75 //suggested by LuoBiao
	WarnStdTableSpace  = 85
	WarnSTDJVMHeapUsed = 85
	WarnSTDQueueLength = 10
)
