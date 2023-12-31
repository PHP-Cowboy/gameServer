package iface

import "net"

// 定义连接接⼝口
type Connection interface {
	//启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接状态M
	Stop()
	//获取当前连接ID
	GetConnId() uint32
	//从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获取远程客户端地址信息
	RemoteAddr() net.Addr
	//直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendMsg(msgId uint32, data []byte) (err error)
	//直接将Message数据发送给远程的TCP客户端(有缓冲)
	SendBuffMsg(msgId uint32, data []byte) (err error)
	//设置链接属性
	SetProperty(key string, value interface{})
	//获取链接属性
	GetProperty(key string) (interface{}, error)
	//移除链接属性
	RemoveProperty(key string)
}

type HandleFunc func(conn *net.TCPConn, data []byte, cnt int) (err error)
