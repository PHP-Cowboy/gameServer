package core

import "fmt"

// AOI管理模块
type AOIManager struct {
	MinX  int           //区域左边界坐标
	MaxX  int           //区域右边界坐标
	CntX  int           //x方向格子的数量
	MinY  int           //区域上边界坐标
	MaxY  int           //区域下边界坐标
	CntY  int           //y方向的格子数量
	grids map[int]*Grid //当前区域中都有哪些格子，key=格子ID， value=格子对象
}

// 初始化一个AOI区域
func NewAOIManager(minX, maxX, cntX, minY, maxY, cntY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntX:  cntX,
		MinY:  minY,
		MaxY:  maxY,
		CntY:  cntY,
		grids: make(map[int]*Grid),
	}

	//给AOI初始化区域中所有的格子
	for y := 0; y < cntY; y++ {
		for x := 0; x < cntX; x++ {
			//计算格子ID
			//格子编号：id = idy *nx + idx  (利用格子坐标得到格子编号)
			gid := y*cntX + x

			//初始化一个格子放在AOI中的map里，key是当前格子的ID
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength())
		}
	}

	return aoiMgr

}

// 得到每个格子在x轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxY - m.MinX) / m.CntY
}

// 得到每个格子在x轴方向的长度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntY
}

// 根据格子的gID得到当前周边的九宫格信息
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否存在
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//将当前gid添加到九宫格中
	grids = append(grids, m.grids[gID])

	//根据gid得到当前格子所在的X轴编号
	idx := gID % m.CntX

	//判断当前idx左边是否还有格子
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	//判断当前的idx右边是否还有格子
	if idx < m.CntX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	//将x轴当前的格子都取出，进行遍历，再分别得到每个格子的上下是否有格子

	//得到当前x轴的格子id集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.Gid)
	}

	//遍历x轴格子
	for _, v := range gidsX {
		//计算该格子处于第几列
		idy := v / m.CntX

		//判断当前的idy上边是否还有格子
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntX])
		}
		//判断当前的idy下边是否还有格子
		if idy < m.CntY-1 {
			grids = append(grids, m.grids[v+m.CntX])
		}
	}

	return
}

// 通过横纵坐标获取对应的格子Id
func (m *AOIManager) GetGidByPos(x, y float32) int {
	gx := (int(x) - m.MinX) / m.gridWidth()
	gy := (int(x) - m.MinY) / m.gridLength()

	return gy*m.CntX + gx
}

// 通过横纵坐标得到周边九宫格内的全部PlayerIDs
func (m *AOIManager) GetPIdsByPos(x, y float32) (playerIds []int) {
	//根据横纵坐标得到当前坐标属于哪个格子ID
	gid := m.GetGidByPos(x, y)
	//根据格子ID得到周边九宫格的信息
	grids := m.GetSurroundGridsByGid(gid)

	for _, grid := range grids {
		playerIds = append(playerIds, grid.GetPlayerIds()...)
	}

	return
}

// 通过GID获取当前格子的全部playerId
func (m *AOIManager) GetPIdsByGid(gId int) (playerIds []int) {
	return m.grids[gId].GetPlayerIds()
}

// 移除一个格子中的PlayerId
func (m *AOIManager) RemovePidFromGrid(pId, gId int) {
	m.grids[gId].Remove(pId)
	return
}

// 添加一个PlayerId到一个格子中
func (m *AOIManager) AddPidToGrid(pId, gId int) {
	m.grids[gId].Add(pId)
	return
}

// 通过横纵坐标添加一个Player到一个格子中
func (m *AOIManager) AddToGridByPos(pId int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	m.grids[gid].Add(pId)
}

// 通过横纵坐标把一个Player从对应的格子中删除
func (m *AOIManager) RemoveFromGridByPos(pId int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	m.grids[gid].Remove(pId)
}

// 打印信息方法
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManagr:\nminX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOI Manager:\n",
		m.MinX, m.MaxX, m.CntX, m.MinY, m.MaxY, m.CntY)

	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}
