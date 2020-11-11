package src

import (
	"golang-starter-project/src/infra"
	"golang-starter-project/src/infra/starters"
)

// 将所有Starter预先载入启动容器
func init() {
	infra.Register(&starters.LogStarter{})
}
