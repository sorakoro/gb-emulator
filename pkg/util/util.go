package util

func SetMSB(value byte, b bool) byte {
	if b {
		value |= 0x80
	} else {
		value &= 0x7f
	}
	return value
}

func SetLSB(value byte, b bool) byte {
	if b {
		value |= 1
	} else {
		value &= 0xfe
	}
	return value
}
