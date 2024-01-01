package consts

const (
	SyncPid          = 1   //Server 同步玩家本次登录的ID
	Talk             = 2   //Client 世界聊天
	Position         = 3   //Client 移动
	BroadCast        = 200 //Server 广播消息(Tp 1 世界聊天 2 坐标(出生点同步) 3 动作 4 移动之后坐标信息更新)
	BroadCastSyncPid = 201 //Server 广播消息 掉线/aoi消失在视野
	SyncPlayers      = 202 //Server 同步周围的人位置信息(包括自己)
)
