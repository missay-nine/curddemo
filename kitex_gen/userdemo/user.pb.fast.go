// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package userdemo

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMessage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ServiceTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateUserRequest[number], err)
}

func (x *CreateUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateUserRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateUserResponse[number], err)
}

func (x *CreateUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *MGetUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MGetUserRequest[number], err)
}

func (x *MGetUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.UserIds = append(x.UserIds, v)
			return offset, err
		})
	return offset, err
}

func (x *MGetUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MGetUserResponse[number], err)
}

func (x *MGetUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Users = append(x.Users, &v)
	return offset, nil
}

func (x *MGetUserResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *CheckUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckUserRequest[number], err)
}

func (x *CheckUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckUserRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckUserResponse[number], err)
}

func (x *CheckUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CheckUserResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMessage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMessage())
	return offset
}

func (x *BaseResp) fastWriteField3(buf []byte) (offset int) {
	if x.ServiceTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetServiceTime())
	return offset
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUserName())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAvatar())
	return offset
}

func (x *CreateUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateUserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserName())
	return offset
}

func (x *CreateUserRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *CreateUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *MGetUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *MGetUserRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.UserIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetUserIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetUserIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *MGetUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MGetUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Users == nil {
		return offset
	}
	for i := range x.GetUsers() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUsers()[i])
	}
	return offset
}

func (x *MGetUserResponse) fastWriteField2(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetBaseResp())
	return offset
}

func (x *CheckUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CheckUserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserName())
	return offset
}

func (x *CheckUserRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *CheckUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CheckUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CheckUserResponse) fastWriteField2(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetBaseResp())
	return offset
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetStatusCode())
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.StatusMessage == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMessage())
	return n
}

func (x *BaseResp) sizeField3() (n int) {
	if x.ServiceTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetServiceTime())
	return n
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUserName())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAvatar())
	return n
}

func (x *CreateUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateUserRequest) sizeField1() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserName())
	return n
}

func (x *CreateUserRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *CreateUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateUserResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *MGetUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *MGetUserRequest) sizeField1() (n int) {
	if len(x.UserIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetUserIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetUserIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *MGetUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MGetUserResponse) sizeField1() (n int) {
	if x.Users == nil {
		return n
	}
	for i := range x.GetUsers() {
		n += fastpb.SizeMessage(1, x.GetUsers()[i])
	}
	return n
}

func (x *MGetUserResponse) sizeField2() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetBaseResp())
	return n
}

func (x *CheckUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CheckUserRequest) sizeField1() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserName())
	return n
}

func (x *CheckUserRequest) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *CheckUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CheckUserResponse) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *CheckUserResponse) sizeField2() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetBaseResp())
	return n
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMessage",
	3: "ServiceTime",
}

var fieldIDToName_User = map[int32]string{
	1: "UserId",
	2: "UserName",
	3: "Avatar",
}

var fieldIDToName_CreateUserRequest = map[int32]string{
	1: "UserName",
	2: "Password",
}

var fieldIDToName_CreateUserResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_MGetUserRequest = map[int32]string{
	1: "UserIds",
}

var fieldIDToName_MGetUserResponse = map[int32]string{
	1: "Users",
	2: "BaseResp",
}

var fieldIDToName_CheckUserRequest = map[int32]string{
	1: "UserName",
	2: "Password",
}

var fieldIDToName_CheckUserResponse = map[int32]string{
	1: "UserId",
	2: "BaseResp",
}