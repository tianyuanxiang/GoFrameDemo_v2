// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GfUser is the golang structure of table gf_user for DAO operations like Where/Data.
type GfUser struct {
	g.Meta    `orm:"table:gf_user, do:true"`
	Id        interface{} // UID
	Passport  interface{} // 账号
	Password  interface{} // MD5密码
	Nickname  interface{} // 昵称
	Avatar    interface{} // 头像地址
	Status    interface{} // 状态 0:启用 1:禁用
	Gender    interface{} // 性别 0: 未设置 1: 男 2: 女
	CreatedAt *gtime.Time // 注册时间
	UpdatedAt *gtime.Time // 更新时间
}
