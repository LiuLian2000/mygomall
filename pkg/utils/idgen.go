package utils

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	// 全局唯一的 Node 实例
	globalNode *snowflake.Node
	// 确保 Node 只初始化一次
	nodeOnce sync.Once

	nodeId int64 = 1
)

// InitializeNode 初始化 Snowflake 节点
func InitializeNode(nodeID int64) error {
	var initErr error

	nodeOnce.Do(func() {
		node, err := snowflake.NewNode(nodeID)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize snowflake node: %w", err)
			return
		}
		globalNode = node
	})

	return initErr
}

// GenerateID 生成唯一 ID
func GenerateID() (int64, error) {
	if globalNode == nil {
		InitializeNode(nodeId)
	}
	return globalNode.Generate().Int64(), nil
}
