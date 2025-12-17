package util

import (
	snowflake "github.com/yockii/snowflake_ext"
)

var snowflakeWorker *snowflake.Worker

func InitSnowflake(workID uint64) {
	snowflakeWorker, _ = snowflake.NewSnowflakeWithConfig(workID, snowflake.WorkerOption{
		// 基准时间戳用2025-12-01
		BaseEpoch: 1675392000000,
	})
}

func NextSnowflakeID() uint64 {
	if snowflakeWorker == nil {
		panic("snowflake worker not initialized")
	}
	return snowflakeWorker.NextId()
}
