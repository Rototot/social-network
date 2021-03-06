// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package endpoints

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	services "social-network/pkg/users/services"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson24ffd8d2DecodeSocialNetworkPkgUsersEndpoints(in *jlexer.Lexer, out *PostLogoutRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "SessionId":
			out.SessionId = services.SessionId(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson24ffd8d2EncodeSocialNetworkPkgUsersEndpoints(out *jwriter.Writer, in PostLogoutRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"SessionId\":"
		out.RawString(prefix[1:])
		out.String(string(in.SessionId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostLogoutRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson24ffd8d2EncodeSocialNetworkPkgUsersEndpoints(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostLogoutRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson24ffd8d2EncodeSocialNetworkPkgUsersEndpoints(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostLogoutRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson24ffd8d2DecodeSocialNetworkPkgUsersEndpoints(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostLogoutRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson24ffd8d2DecodeSocialNetworkPkgUsersEndpoints(l, v)
}
