// Code generated by "stringer -type=ClientMessageType"; DO NOT EDIT.

package vnc

import "fmt"

const (
	_ClientMessageType_name_0 = "SetPixelFormatMsgType"
	_ClientMessageType_name_1 = "SetEncodingsMsgTypeFramebufferUpdateRequestMsgTypeKeyEventMsgTypePointerEventMsgTypeClientCutTextMsgType"
)

var (
	_ClientMessageType_index_0 = [...]uint8{0, 21}
	_ClientMessageType_index_1 = [...]uint8{0, 19, 50, 65, 84, 104}
)

func (i ClientMessageType) String() string {
	switch {
	case i == 0:
		return _ClientMessageType_name_0
	case 2 <= i && i <= 6:
		i -= 2
		return _ClientMessageType_name_1[_ClientMessageType_index_1[i]:_ClientMessageType_index_1[i+1]]
	default:
		return fmt.Sprintf("ClientMessageType(%d)", i)
	}
}