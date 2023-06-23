package gbc

import "github.com/sorakoro/gb-emulator/pkg/cart"

type ROM struct {
	bank   byte
	buffer [256][0x4000]byte
}

type RAM struct {
	bank   byte
	buffer [256][0x2000]byte
}

type CurInst struct {
	Opcode byte
	PC     uint16
}

type GBC struct {
	Reg       Register
	Inst      CurInst
	ROM       ROM
	RAM       RAM
	Cartridge *cart.Cartridge
}

func New(rom []byte) *GBC {
	return &GBC{
		Cartridge: cart.New(rom),
	}
}

func (g *GBC) read8(addr uint16) (value byte) {
	switch {
	case addr < 0x4000: // ROM bank0
		break
	case addr >= 0x4000 && addr < 0x8000: // ROM bank1...256
		break
	case addr >= 0x8000 && addr < 0xa000: // VRAM
		break
	case addr >= 0xa000 && addr < 0xc000: // RTC or RAM
		break
	case addr >= 0xc000 && addr < 0xd000: // WRAM bank0
		break
	case addr >= 0xd000 && addr < 0xe000: // WRAM bank1...7
		break
	case addr >= 0xfe00 && addr < 0xfea0: // OAM
		break
	case addr >= 0xff00: // IO, HRAM, IE
		break
	}
	return 0
}

func (g *GBC) write8(addr uint16, value byte) {
	switch {
	case addr >= 0x8000 && addr < 0xa000: // VRAM
		break
	case addr >= 0xa000 && addr < 0xc000: // RTC or RAM
		break
	case addr >= 0xc000 && addr < 0xd000: // WRAM bank0
		break
	case addr >= 0xd000 && addr < 0xe000: // WRAM bank1...7
		break
	case addr >= 0xfe00 && addr <= 0xfe9f: // OAM
		break
	case addr >= 0xff00: // IO, HRAM, IE
		break
	}
}

func (g *GBC) d8Fetch() byte {
	value := g.read8(g.Inst.PC + 1)
	g.Reg.PC++
	return value
}

func (g *GBC) a16Fetch() uint16 {
	lower, upper := uint16(g.read8(g.Inst.PC+1)), uint16(g.read8(g.Inst.PC+2))
	g.Reg.PC += 2
	return (upper << 8) | lower
}

func (g *GBC) TransferROM(rom []byte) {
	switch g.Cartridge.Type {
	case 0x00:
		break
	case 0x01:
		break
	case 0x02, 0x03:
		break
	case 0x05, 0x06:
		break
	case 0x08, 0x09:
		break
	case 0x0f, 0x10, 0x11, 0x12, 0x13:
		break
	case 0x19, 0x1a, 0x1b:
		break
	default:
		panic("")
	}
}
