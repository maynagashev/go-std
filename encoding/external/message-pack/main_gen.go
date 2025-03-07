package main

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AccountBalance) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "account_id_hash":
			z.AccountIdHash, err = dc.ReadBytes(z.AccountIdHash)
			if err != nil {
				err = msgp.WrapError(err, "AccountIdHash")
				return
			}
		case "amounts":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Amounts")
				return
			}
			if cap(z.Amounts) >= int(zb0002) {
				z.Amounts = (z.Amounts)[:zb0002]
			} else {
				z.Amounts = make([]CurrencyAmount, zb0002)
			}
			for za0001 := range z.Amounts {
				var zb0003 uint32
				zb0003, err = dc.ReadMapHeader()
				if err != nil {
					err = msgp.WrapError(err, "Amounts", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						err = msgp.WrapError(err, "Amounts", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Amount":
						z.Amounts[za0001].Amount, err = dc.ReadInt64()
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001, "Amount")
							return
						}
					case "Decimals":
						z.Amounts[za0001].Decimals, err = dc.ReadInt8()
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001, "Decimals")
							return
						}
					case "Symbol":
						z.Amounts[za0001].Symbol, err = dc.ReadString()
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001, "Symbol")
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001)
							return
						}
					}
				}
			}
		case "is_blocked":
			z.IsBlocked, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "IsBlocked")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AccountBalance) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "account_id_hash"
	err = en.Append(0x83, 0xaf, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.AccountIdHash)
	if err != nil {
		err = msgp.WrapError(err, "AccountIdHash")
		return
	}
	// write "amounts"
	err = en.Append(0xa7, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Amounts)))
	if err != nil {
		err = msgp.WrapError(err, "Amounts")
		return
	}
	for za0001 := range z.Amounts {
		// map header, size 3
		// write "Amount"
		err = en.Append(0x83, 0xa6, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74)
		if err != nil {
			return
		}
		err = en.WriteInt64(z.Amounts[za0001].Amount)
		if err != nil {
			err = msgp.WrapError(err, "Amounts", za0001, "Amount")
			return
		}
		// write "Decimals"
		err = en.Append(0xa8, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73)
		if err != nil {
			return
		}
		err = en.WriteInt8(z.Amounts[za0001].Decimals)
		if err != nil {
			err = msgp.WrapError(err, "Amounts", za0001, "Decimals")
			return
		}
		// write "Symbol"
		err = en.Append(0xa6, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c)
		if err != nil {
			return
		}
		err = en.WriteString(z.Amounts[za0001].Symbol)
		if err != nil {
			err = msgp.WrapError(err, "Amounts", za0001, "Symbol")
			return
		}
	}
	// write "is_blocked"
	err = en.Append(0xaa, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsBlocked)
	if err != nil {
		err = msgp.WrapError(err, "IsBlocked")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AccountBalance) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "account_id_hash"
	o = append(o, 0x83, 0xaf, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, z.AccountIdHash)
	// string "amounts"
	o = append(o, 0xa7, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Amounts)))
	for za0001 := range z.Amounts {
		// map header, size 3
		// string "Amount"
		o = append(o, 0x83, 0xa6, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74)
		o = msgp.AppendInt64(o, z.Amounts[za0001].Amount)
		// string "Decimals"
		o = append(o, 0xa8, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73)
		o = msgp.AppendInt8(o, z.Amounts[za0001].Decimals)
		// string "Symbol"
		o = append(o, 0xa6, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c)
		o = msgp.AppendString(o, z.Amounts[za0001].Symbol)
	}
	// string "is_blocked"
	o = append(o, 0xaa, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64)
	o = msgp.AppendBool(o, z.IsBlocked)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccountBalance) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "account_id_hash":
			z.AccountIdHash, bts, err = msgp.ReadBytesBytes(bts, z.AccountIdHash)
			if err != nil {
				err = msgp.WrapError(err, "AccountIdHash")
				return
			}
		case "amounts":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Amounts")
				return
			}
			if cap(z.Amounts) >= int(zb0002) {
				z.Amounts = (z.Amounts)[:zb0002]
			} else {
				z.Amounts = make([]CurrencyAmount, zb0002)
			}
			for za0001 := range z.Amounts {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Amounts", za0001)
					return
				}
				for zb0003 > 0 {
					zb0003--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "Amounts", za0001)
						return
					}
					switch msgp.UnsafeString(field) {
					case "Amount":
						z.Amounts[za0001].Amount, bts, err = msgp.ReadInt64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001, "Amount")
							return
						}
					case "Decimals":
						z.Amounts[za0001].Decimals, bts, err = msgp.ReadInt8Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001, "Decimals")
							return
						}
					case "Symbol":
						z.Amounts[za0001].Symbol, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001, "Symbol")
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							err = msgp.WrapError(err, "Amounts", za0001)
							return
						}
					}
				}
			}
		case "is_blocked":
			z.IsBlocked, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IsBlocked")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AccountBalance) Msgsize() (s int) {
	s = 1 + 16 + msgp.BytesPrefixSize + len(z.AccountIdHash) + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Amounts {
		s += 1 + 7 + msgp.Int64Size + 9 + msgp.Int8Size + 7 + msgp.StringPrefixSize + len(z.Amounts[za0001].Symbol)
	}
	s += 11 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CurrencyAmount) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Amount":
			z.Amount, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Amount")
				return
			}
		case "Decimals":
			z.Decimals, err = dc.ReadInt8()
			if err != nil {
				err = msgp.WrapError(err, "Decimals")
				return
			}
		case "Symbol":
			z.Symbol, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Symbol")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z CurrencyAmount) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Amount"
	err = en.Append(0x83, 0xa6, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Amount)
	if err != nil {
		err = msgp.WrapError(err, "Amount")
		return
	}
	// write "Decimals"
	err = en.Append(0xa8, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.Decimals)
	if err != nil {
		err = msgp.WrapError(err, "Decimals")
		return
	}
	// write "Symbol"
	err = en.Append(0xa6, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.Symbol)
	if err != nil {
		err = msgp.WrapError(err, "Symbol")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z CurrencyAmount) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Amount"
	o = append(o, 0x83, 0xa6, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.Amount)
	// string "Decimals"
	o = append(o, 0xa8, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73)
	o = msgp.AppendInt8(o, z.Decimals)
	// string "Symbol"
	o = append(o, 0xa6, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c)
	o = msgp.AppendString(o, z.Symbol)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CurrencyAmount) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Amount":
			z.Amount, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Amount")
				return
			}
		case "Decimals":
			z.Decimals, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Decimals")
				return
			}
		case "Symbol":
			z.Symbol, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Symbol")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z CurrencyAmount) Msgsize() (s int) {
	s = 1 + 7 + msgp.Int64Size + 9 + msgp.Int8Size + 7 + msgp.StringPrefixSize + len(z.Symbol)
	return
}
