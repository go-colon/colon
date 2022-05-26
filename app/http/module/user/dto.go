/**
 * @author:cb <fastopencn@gmail.com>
 * @date:2022/5/26
 **/
package user

import "time"

// UserDTO  表示输出到外部的用户信息
type UserDTO struct {
	ID        int64     `json:"id,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
