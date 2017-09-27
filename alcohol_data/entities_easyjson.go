// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package alcohol_data

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3e8ab7adDecodeAlchiveAlcoholData(in *jlexer.Lexer, out *SpiritsWOCategory) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(SpiritsWOCategory, 0, 1)
			} else {
				*out = SpiritsWOCategory{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 SpiritWOCategory
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3e8ab7adEncodeAlchiveAlcoholData(out *jwriter.Writer, in SpiritsWOCategory) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v SpiritsWOCategory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeAlchiveAlcoholData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SpiritsWOCategory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeAlchiveAlcoholData(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SpiritsWOCategory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeAlchiveAlcoholData(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SpiritsWOCategory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeAlchiveAlcoholData(l, v)
}
func easyjson3e8ab7adDecodeAlchiveAlcoholData1(in *jlexer.Lexer, out *Spirits) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Spirits, 0, 1)
			} else {
				*out = Spirits{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 Spirit
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3e8ab7adEncodeAlchiveAlcoholData1(out *jwriter.Writer, in Spirits) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Spirits) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeAlchiveAlcoholData1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Spirits) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeAlchiveAlcoholData1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Spirits) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeAlchiveAlcoholData1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Spirits) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeAlchiveAlcoholData1(l, v)
}
func easyjson3e8ab7adDecodeAlchiveAlcoholData2(in *jlexer.Lexer, out *SpiritWOCategory) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "bars":
			if in.IsNull() {
				in.Skip()
				out.Bars = nil
			} else {
				in.Delim('[')
				if out.Bars == nil {
					if !in.IsDelim(']') {
						out.Bars = make([]string, 0, 4)
					} else {
						out.Bars = []string{}
					}
				} else {
					out.Bars = (out.Bars)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Bars = append(out.Bars, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.Categories = append(out.Categories, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "isCategory":
			out.IsCategory = bool(in.Bool())
		case "preferredUnits":
			if in.IsNull() {
				in.Skip()
				out.PreferredUnits = nil
			} else {
				in.Delim('[')
				if out.PreferredUnits == nil {
					if !in.IsDelim(']') {
						out.PreferredUnits = make([]PreferredUnit, 0, 1)
					} else {
						out.PreferredUnits = []PreferredUnit{}
					}
				} else {
					out.PreferredUnits = (out.PreferredUnits)[:0]
				}
				for !in.IsDelim(']') {
					var v9 PreferredUnit
					(v9).UnmarshalEasyJSON(in)
					out.PreferredUnits = append(out.PreferredUnits, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "imageURL":
			out.ImageURL = string(in.String())
		case "productURL":
			out.ProductURL = string(in.String())
		case "volume":
			out.Volume = float32(in.Float32())
		case "value":
			out.Value = float32(in.Float32())
		case "caloricity":
			out.Caloricity = float32(in.Float32())
		case "description":
			out.Description = string(in.String())
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
func easyjson3e8ab7adEncodeAlchiveAlcoholData2(out *jwriter.Writer, in SpiritWOCategory) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if len(in.Bars) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bars\":")
		if in.Bars == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.Bars {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	if len(in.Categories) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"categories\":")
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Categories {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	if in.IsCategory {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isCategory\":")
		out.Bool(bool(in.IsCategory))
	}
	if len(in.PreferredUnits) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"preferredUnits\":")
		if in.PreferredUnits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.PreferredUnits {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.ImageURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"imageURL\":")
		out.String(string(in.ImageURL))
	}
	if in.ProductURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"productURL\":")
		out.String(string(in.ProductURL))
	}
	if in.Volume != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"volume\":")
		out.Float32(float32(in.Volume))
	}
	if in.Value != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"value\":")
		out.Float32(float32(in.Value))
	}
	if in.Caloricity != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"caloricity\":")
		out.Float32(float32(in.Caloricity))
	}
	if in.Description != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SpiritWOCategory) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeAlchiveAlcoholData2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SpiritWOCategory) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeAlchiveAlcoholData2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SpiritWOCategory) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeAlchiveAlcoholData2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SpiritWOCategory) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeAlchiveAlcoholData2(l, v)
}
func easyjson3e8ab7adDecodeAlchiveAlcoholData3(in *jlexer.Lexer, out *Spirit) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "bars":
			if in.IsNull() {
				in.Skip()
				out.Bars = nil
			} else {
				in.Delim('[')
				if out.Bars == nil {
					if !in.IsDelim(']') {
						out.Bars = make([]string, 0, 4)
					} else {
						out.Bars = []string{}
					}
				} else {
					out.Bars = (out.Bars)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Bars = append(out.Bars, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]Spirit, 0, 1)
					} else {
						out.Categories = []Spirit{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v17 Spirit
					(v17).UnmarshalEasyJSON(in)
					out.Categories = append(out.Categories, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "isCategory":
			out.IsCategory = bool(in.Bool())
		case "preferredUnits":
			if in.IsNull() {
				in.Skip()
				out.PreferredUnits = nil
			} else {
				in.Delim('[')
				if out.PreferredUnits == nil {
					if !in.IsDelim(']') {
						out.PreferredUnits = make([]PreferredUnit, 0, 1)
					} else {
						out.PreferredUnits = []PreferredUnit{}
					}
				} else {
					out.PreferredUnits = (out.PreferredUnits)[:0]
				}
				for !in.IsDelim(']') {
					var v18 PreferredUnit
					(v18).UnmarshalEasyJSON(in)
					out.PreferredUnits = append(out.PreferredUnits, v18)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "imageURL":
			out.ImageURL = string(in.String())
		case "productURL":
			out.ProductURL = string(in.String())
		case "volume":
			out.Volume = float32(in.Float32())
		case "value":
			out.Value = float32(in.Float32())
		case "caloricity":
			out.Caloricity = float32(in.Float32())
		case "description":
			out.Description = string(in.String())
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
func easyjson3e8ab7adEncodeAlchiveAlcoholData3(out *jwriter.Writer, in Spirit) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.ID))
	}
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if len(in.Bars) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"bars\":")
		if in.Bars == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v19, v20 := range in.Bars {
				if v19 > 0 {
					out.RawByte(',')
				}
				out.String(string(v20))
			}
			out.RawByte(']')
		}
	}
	if len(in.Categories) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"categories\":")
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v21, v22 := range in.Categories {
				if v21 > 0 {
					out.RawByte(',')
				}
				(v22).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.IsCategory {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"isCategory\":")
		out.Bool(bool(in.IsCategory))
	}
	if len(in.PreferredUnits) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"preferredUnits\":")
		if in.PreferredUnits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v23, v24 := range in.PreferredUnits {
				if v23 > 0 {
					out.RawByte(',')
				}
				(v24).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.ImageURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"imageURL\":")
		out.String(string(in.ImageURL))
	}
	if in.ProductURL != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"productURL\":")
		out.String(string(in.ProductURL))
	}
	if in.Volume != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"volume\":")
		out.Float32(float32(in.Volume))
	}
	if in.Value != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"value\":")
		out.Float32(float32(in.Value))
	}
	if in.Caloricity != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"caloricity\":")
		out.Float32(float32(in.Caloricity))
	}
	if in.Description != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"description\":")
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Spirit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeAlchiveAlcoholData3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Spirit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeAlchiveAlcoholData3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Spirit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeAlchiveAlcoholData3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Spirit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeAlchiveAlcoholData3(l, v)
}
func easyjson3e8ab7adDecodeAlchiveAlcoholData4(in *jlexer.Lexer, out *PreferredUnit) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = float32(in.Float32())
		case "uom":
			out.UOM = string(in.String())
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
func easyjson3e8ab7adEncodeAlchiveAlcoholData4(out *jwriter.Writer, in PreferredUnit) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"name\":")
		out.String(string(in.Name))
	}
	if in.Value != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"value\":")
		out.Float32(float32(in.Value))
	}
	if in.UOM != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"uom\":")
		out.String(string(in.UOM))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PreferredUnit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3e8ab7adEncodeAlchiveAlcoholData4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PreferredUnit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3e8ab7adEncodeAlchiveAlcoholData4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PreferredUnit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3e8ab7adDecodeAlchiveAlcoholData4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PreferredUnit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3e8ab7adDecodeAlchiveAlcoholData4(l, v)
}
