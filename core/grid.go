package core

import (
	"fmt"
	"sync"
)

// 一个地图中的格子类
type Grid struct {
	Gid       int              //格子id
	MinX      int              //格子x轴最小坐标
	MaxX      int              //格子x轴最大坐标
	MinY      int              //格子y轴最小坐标
	MaxY      int              //格子y轴最大坐标
	playerIds map[int]struct{} //当前格子内的玩家或者物体成员ID
	pIdLock   sync.RWMutex     //playerIDs的保护map的锁
}

// 初始化一个格子
func NewGrid(gId, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		Gid:       gId,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIds: make(map[int]struct{}),
	}
}

// 向当前格子中添加一个玩家
func (g *Grid) Add(playerId int) {
	g.pIdLock.Lock()
	defer g.pIdLock.Unlock()
	g.playerIds[playerId] = struct{}{}
}

// 从格子中删除一个玩家
func (g *Grid) Remove(playerId int) {
	g.pIdLock.Lock()
	defer g.pIdLock.Unlock()
	delete(g.playerIds, playerId)
}

// 得到当前格子中所有的玩家
func (g *Grid) GetPlayerIds() (playerIds []int) {
	g.pIdLock.RLock()
	defer g.pIdLock.RUnlock()
	playerIds = make([]int, 0, len(g.playerIds))

	for id, _ := range g.playerIds {
		playerIds = append(playerIds, id)
	}

	return
}

// 打印信息方法
func (g *Grid) String() string {
	return fmt.Sprintf(
		"Grid id: %d, minX:%d, maxX:%d, minY:%d, maxY:%d, playerIDs:%v",
		g.Gid,
		g.MinX,
		g.MaxX,
		g.MinY,
		g.MaxY,
		g.playerIds,
	)
}
