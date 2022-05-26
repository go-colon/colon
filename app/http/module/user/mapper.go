/**
 * @author:cb <fastopencn@gmail.com>
 * @date:2022/5/26
 **/
package user

import "github.com/go-colon/colon/app/provider/user"

// ConvertUserToDTO 将user转换为UserDTO
func ConvertUserToDTO(user *user.User) *UserDTO {
	if user == nil {
		return nil
	}
	return &UserDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
	}
}
