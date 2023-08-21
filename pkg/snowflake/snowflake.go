package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	// 实例
	sonyFlake *sonyflake.Sonyflake
	// 机器ID
	sonyMachineID uint16
)

// getMachineID 返回全局定义的机器ID
func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传入当前的机器ID
func Init(machineId uint16, startTime string) (err error) {
	sonyMachineID = machineId
	// 初始化一个开始的时间
	t, _ := time.Parse("2006-01-02", startTime)
	// 生成全局配置
	settings := sonyflake.Settings{
		StartTime: t,
		// 指定机器ID
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回生成的id值
func GetID() (id int64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	snowId, err := sonyFlake.NextID()
	id = int64(snowId)
	return
}
