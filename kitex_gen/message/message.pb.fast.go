// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package message

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Message) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Message[number], err)
}

func (x *Message) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Message) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Message) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FromUserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Message) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CreateTime = &tmp
	return offset, err
}

func (x *MessageChatRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MessageChatRequest[number], err)
}

func (x *MessageChatRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromUserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *MessageChatRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *MessageChatRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PreMsgTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *MessageChatResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MessageChatResponse[number], err)
}

func (x *MessageChatResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *MessageChatResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *MessageChatResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v Message
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.MessageList = append(x.MessageList, &v)
	return offset, nil
}

func (x *MessageActionRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MessageActionRequest[number], err)
}

func (x *MessageActionRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromUserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MessageActionRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *MessageActionRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *MessageActionRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MessageActionResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MessageActionResponse[number], err)
}

func (x *MessageActionResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *MessageActionResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *Message) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *Message) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Message) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetToUserId())
	return offset
}

func (x *Message) fastWriteField3(buf []byte) (offset int) {
	if x.FromUserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetFromUserId())
	return offset
}

func (x *Message) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetContent())
	return offset
}

func (x *Message) fastWriteField5(buf []byte) (offset int) {
	if x.CreateTime == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetCreateTime())
	return offset
}

func (x *MessageChatRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *MessageChatRequest) fastWriteField1(buf []byte) (offset int) {
	if x.FromUserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetFromUserId())
	return offset
}

func (x *MessageChatRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetToUserId())
	return offset
}

func (x *MessageChatRequest) fastWriteField3(buf []byte) (offset int) {
	if x.PreMsgTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetPreMsgTime())
	return offset
}

func (x *MessageChatResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *MessageChatResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *MessageChatResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *MessageChatResponse) fastWriteField3(buf []byte) (offset int) {
	if x.MessageList == nil {
		return offset
	}
	for i := range x.GetMessageList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetMessageList()[i])
	}
	return offset
}

func (x *MessageActionRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *MessageActionRequest) fastWriteField1(buf []byte) (offset int) {
	if x.FromUserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFromUserId())
	return offset
}

func (x *MessageActionRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetToUserId())
	return offset
}

func (x *MessageActionRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetActionType())
	return offset
}

func (x *MessageActionRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetContent())
	return offset
}

func (x *MessageActionResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MessageActionResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *MessageActionResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *Message) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *Message) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetId())
	return n
}

func (x *Message) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetToUserId())
	return n
}

func (x *Message) sizeField3() (n int) {
	if x.FromUserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetFromUserId())
	return n
}

func (x *Message) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetContent())
	return n
}

func (x *Message) sizeField5() (n int) {
	if x.CreateTime == nil {
		return n
	}
	n += fastpb.SizeString(5, x.GetCreateTime())
	return n
}

func (x *MessageChatRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *MessageChatRequest) sizeField1() (n int) {
	if x.FromUserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetFromUserId())
	return n
}

func (x *MessageChatRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetToUserId())
	return n
}

func (x *MessageChatRequest) sizeField3() (n int) {
	if x.PreMsgTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetPreMsgTime())
	return n
}

func (x *MessageChatResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *MessageChatResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *MessageChatResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *MessageChatResponse) sizeField3() (n int) {
	if x.MessageList == nil {
		return n
	}
	for i := range x.GetMessageList() {
		n += fastpb.SizeMessage(3, x.GetMessageList()[i])
	}
	return n
}

func (x *MessageActionRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *MessageActionRequest) sizeField1() (n int) {
	if x.FromUserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFromUserId())
	return n
}

func (x *MessageActionRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetToUserId())
	return n
}

func (x *MessageActionRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetActionType())
	return n
}

func (x *MessageActionRequest) sizeField4() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetContent())
	return n
}

func (x *MessageActionResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MessageActionResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetStatusCode())
	return n
}

func (x *MessageActionResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

var fieldIDToName_Message = map[int32]string{
	1: "Id",
	2: "ToUserId",
	3: "FromUserId",
	4: "Content",
	5: "CreateTime",
}

var fieldIDToName_MessageChatRequest = map[int32]string{
	1: "FromUserId",
	2: "ToUserId",
	3: "PreMsgTime",
}

var fieldIDToName_MessageChatResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "MessageList",
}

var fieldIDToName_MessageActionRequest = map[int32]string{
	1: "FromUserId",
	2: "ToUserId",
	3: "ActionType",
	4: "Content",
}

var fieldIDToName_MessageActionResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}
