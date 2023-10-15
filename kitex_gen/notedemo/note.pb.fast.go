// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package notedemo

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

func (x *Note) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Note[number], err)
}

func (x *Note) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.NoteId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Note) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Note) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Note) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UserAvatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Note) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Note) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Note) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.CreateTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateNoteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateNoteRequest[number], err)
}

func (x *CreateNoteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateNoteRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateNoteRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateNoteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateNoteResponse[number], err)
}

func (x *CreateNoteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *DeleteNoteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteNoteRequest[number], err)
}

func (x *DeleteNoteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.NoteId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteNoteRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteNoteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteNoteResponse[number], err)
}

func (x *DeleteNoteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *UpdateNoteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateNoteRequest[number], err)
}

func (x *UpdateNoteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.NoteId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateNoteRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateNoteRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateNoteRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateNoteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateNoteResponse[number], err)
}

func (x *UpdateNoteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *MGetNoteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MGetNoteRequest[number], err)
}

func (x *MGetNoteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.NoteIds = append(x.NoteIds, v)
			return offset, err
		})
	return offset, err
}

func (x *MGetNoteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MGetNoteResponse[number], err)
}

func (x *MGetNoteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Note
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Notes = append(x.Notes, &v)
	return offset, nil
}

func (x *MGetNoteResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *QueryNoteRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 4:
		offset, err = x.fastReadField4(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryNoteRequest[number], err)
}

func (x *QueryNoteRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *QueryNoteRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SearchKey, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *QueryNoteRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Offset, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *QueryNoteRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Limit, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *QueryNoteResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryNoteResponse[number], err)
}

func (x *QueryNoteResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Note
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Notes = append(x.Notes, &v)
	return offset, nil
}

func (x *QueryNoteResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *QueryNoteResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
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

func (x *Note) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *Note) fastWriteField1(buf []byte) (offset int) {
	if x.NoteId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetNoteId())
	return offset
}

func (x *Note) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *Note) fastWriteField3(buf []byte) (offset int) {
	if x.UserName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserName())
	return offset
}

func (x *Note) fastWriteField4(buf []byte) (offset int) {
	if x.UserAvatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetUserAvatar())
	return offset
}

func (x *Note) fastWriteField5(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetTitle())
	return offset
}

func (x *Note) fastWriteField6(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetContent())
	return offset
}

func (x *Note) fastWriteField7(buf []byte) (offset int) {
	if x.CreateTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetCreateTime())
	return offset
}

func (x *CreateNoteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateNoteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTitle())
	return offset
}

func (x *CreateNoteRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetContent())
	return offset
}

func (x *CreateNoteRequest) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *CreateNoteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateNoteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *DeleteNoteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteNoteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.NoteId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetNoteId())
	return offset
}

func (x *DeleteNoteRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *DeleteNoteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteNoteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *UpdateNoteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *UpdateNoteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.NoteId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetNoteId())
	return offset
}

func (x *UpdateNoteRequest) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *UpdateNoteRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetTitle())
	return offset
}

func (x *UpdateNoteRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetContent())
	return offset
}

func (x *UpdateNoteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateNoteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetBaseResp())
	return offset
}

func (x *MGetNoteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *MGetNoteRequest) fastWriteField1(buf []byte) (offset int) {
	if len(x.NoteIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetNoteIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetNoteIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *MGetNoteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MGetNoteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Notes == nil {
		return offset
	}
	for i := range x.GetNotes() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetNotes()[i])
	}
	return offset
}

func (x *MGetNoteResponse) fastWriteField2(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetBaseResp())
	return offset
}

func (x *QueryNoteRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *QueryNoteRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *QueryNoteRequest) fastWriteField2(buf []byte) (offset int) {
	if x.SearchKey == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSearchKey())
	return offset
}

func (x *QueryNoteRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Offset == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetOffset())
	return offset
}

func (x *QueryNoteRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Limit == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetLimit())
	return offset
}

func (x *QueryNoteResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *QueryNoteResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Notes == nil {
		return offset
	}
	for i := range x.GetNotes() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetNotes()[i])
	}
	return offset
}

func (x *QueryNoteResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *QueryNoteResponse) fastWriteField3(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetBaseResp())
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

func (x *Note) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *Note) sizeField1() (n int) {
	if x.NoteId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetNoteId())
	return n
}

func (x *Note) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserId())
	return n
}

func (x *Note) sizeField3() (n int) {
	if x.UserName == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserName())
	return n
}

func (x *Note) sizeField4() (n int) {
	if x.UserAvatar == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetUserAvatar())
	return n
}

func (x *Note) sizeField5() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetTitle())
	return n
}

func (x *Note) sizeField6() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetContent())
	return n
}

func (x *Note) sizeField7() (n int) {
	if x.CreateTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetCreateTime())
	return n
}

func (x *CreateNoteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateNoteRequest) sizeField1() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTitle())
	return n
}

func (x *CreateNoteRequest) sizeField2() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetContent())
	return n
}

func (x *CreateNoteRequest) sizeField3() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetUserId())
	return n
}

func (x *CreateNoteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateNoteResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *DeleteNoteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteNoteRequest) sizeField1() (n int) {
	if x.NoteId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetNoteId())
	return n
}

func (x *DeleteNoteRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserId())
	return n
}

func (x *DeleteNoteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteNoteResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *UpdateNoteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *UpdateNoteRequest) sizeField1() (n int) {
	if x.NoteId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetNoteId())
	return n
}

func (x *UpdateNoteRequest) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserId())
	return n
}

func (x *UpdateNoteRequest) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetTitle())
	return n
}

func (x *UpdateNoteRequest) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetContent())
	return n
}

func (x *UpdateNoteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateNoteResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetBaseResp())
	return n
}

func (x *MGetNoteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *MGetNoteRequest) sizeField1() (n int) {
	if len(x.NoteIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetNoteIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetNoteIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *MGetNoteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MGetNoteResponse) sizeField1() (n int) {
	if x.Notes == nil {
		return n
	}
	for i := range x.GetNotes() {
		n += fastpb.SizeMessage(1, x.GetNotes()[i])
	}
	return n
}

func (x *MGetNoteResponse) sizeField2() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetBaseResp())
	return n
}

func (x *QueryNoteRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *QueryNoteRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *QueryNoteRequest) sizeField2() (n int) {
	if x.SearchKey == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSearchKey())
	return n
}

func (x *QueryNoteRequest) sizeField3() (n int) {
	if x.Offset == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetOffset())
	return n
}

func (x *QueryNoteRequest) sizeField4() (n int) {
	if x.Limit == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetLimit())
	return n
}

func (x *QueryNoteResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *QueryNoteResponse) sizeField1() (n int) {
	if x.Notes == nil {
		return n
	}
	for i := range x.GetNotes() {
		n += fastpb.SizeMessage(1, x.GetNotes()[i])
	}
	return n
}

func (x *QueryNoteResponse) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *QueryNoteResponse) sizeField3() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetBaseResp())
	return n
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMessage",
	3: "ServiceTime",
}

var fieldIDToName_Note = map[int32]string{
	1: "NoteId",
	2: "UserId",
	3: "UserName",
	4: "UserAvatar",
	5: "Title",
	6: "Content",
	7: "CreateTime",
}

var fieldIDToName_CreateNoteRequest = map[int32]string{
	1: "Title",
	2: "Content",
	3: "UserId",
}

var fieldIDToName_CreateNoteResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_DeleteNoteRequest = map[int32]string{
	1: "NoteId",
	2: "UserId",
}

var fieldIDToName_DeleteNoteResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_UpdateNoteRequest = map[int32]string{
	1: "NoteId",
	2: "UserId",
	3: "Title",
	4: "Content",
}

var fieldIDToName_UpdateNoteResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_MGetNoteRequest = map[int32]string{
	1: "NoteIds",
}

var fieldIDToName_MGetNoteResponse = map[int32]string{
	1: "Notes",
	2: "BaseResp",
}

var fieldIDToName_QueryNoteRequest = map[int32]string{
	1: "UserId",
	2: "SearchKey",
	3: "Offset",
	4: "Limit",
}

var fieldIDToName_QueryNoteResponse = map[int32]string{
	1: "Notes",
	2: "Total",
	3: "BaseResp",
}
