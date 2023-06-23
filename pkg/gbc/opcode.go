package gbc

import "github.com/sorakoro/gb-emulator/pkg/util"

func nop(g *GBC, _, _ int) {
	// no operation
}

func stop(g *GBC, _, _ int) {
	// wip
}

func jr(g *GBC, _, _ int) {
	// wip
}

func jrcc(g *GBC, cc, _ int) {
	// wip
}

func jrncc(g *GBC, cc, _ int) {
	// wip
}

func daa(g *GBC, _, _ int) {
	//
}

func cpl(g *GBC, _, _ int) {
	//
}

func incHL(g *GBC, _, _ int) {
	//
}

func decHL(g *GBC, _, _ int) {
	//
}

func op0x36(g *GBC, operand1, operand2 int) {
	//
}

func scf(g *GBC, _, _ int) {
	//
}

func ccf(g *GBC, _, _ int) {
	//
}

func rla(g *GBC, _, _ int) {
	carry, value, bit7 := g.f(flagC), g.Reg.R[A], g.Reg.R[A]>>7 == 1
	g.Reg.R[A] = util.SetLSB(value<<1, carry)
	g.setZNHC(false, false, false, bit7)
}

func rra(g *GBC, _, _ int) {
	oldC, value := g.f(flagC), g.Reg.R[A]
	newC := value&1 == 1
	g.Reg.R[A] = util.SetMSB(value>>1, oldC)
	g.setZNHC(false, false, false, newC)
}

func rlca(g *GBC, _, _ int) {
	value, bit7 := g.Reg.R[A], g.Reg.R[A]>>7 == 1
	g.Reg.R[A] = util.SetLSB(value<<1, bit7)
	g.setZNHC(false, false, false, bit7)
}

func rrca(g *GBC, _, _ int) {
	value, bit0 := g.Reg.R[A], g.Reg.R[A]&1 == 1
	g.Reg.R[A] = util.SetMSB(value>>1, bit0)
	g.setZNHC(false, false, false, bit0)
}

func ld8i(g *GBC, r8, _ int) {
	g.Reg.R[r8] = g.d8Fetch()
}

func ld8m(g *GBC, r8, r16 int) {
	g.Reg.R[r8] = g.read8(g.Reg.R16(r16))
}

func inc8(g *GBC, r8, _ int) {
	value := g.Reg.R[r8] + 1
	carryBits := g.Reg.R[r8] ^ 1 ^ value
	g.Reg.R[r8] = value
	g.setZNH(value == 0, false, (carryBits&1<<4) == 0x10)
}

func dec8(g *GBC, r8, _ int) {
	value := g.Reg.R[r8] - 1
	carryBits := g.Reg.R[r8] ^ 1 ^ value
	g.Reg.R[r8] = value
	g.setZNH(value == 0, true, (carryBits&1<<4) == 0x00)
}

func ld16i(g *GBC, r16, _ int) {
	g.Reg.setR16(r16, g.a16Fetch())
}

func ldm16r(g *GBC, r16, r8 int) {
	g.write8(g.Reg.R16(r16), g.Reg.R[r8])
}

func dec16(g *GBC, r16, _ int) {
	g.Reg.setR16(r16, g.Reg.R16(r16)-1)
}

func inc16(g *GBC, r16, _ int) {
	g.Reg.setR16(r16, g.Reg.R16(r16)+1)
}

func addHL(g *GBC, _, r16 int) {
	lhs, rhs := g.Reg.HL(), g.Reg.R16(r16)
	value := uint32(lhs) + uint32(rhs)
	carryBits := uint32(lhs) ^ uint32(rhs) ^ value
	g.Reg.setHL(uint16(value))
	g.setNHC(false, (carryBits&1<<12) == 0x1000, (carryBits&1<<16) == 0x10000)
}

func op0x08(g *GBC, operand1, operand2 int) {
	addr := g.a16Fetch()
	upper, lower := byte(g.Reg.SP>>8), byte(g.Reg.SP)
	g.write8(addr, lower)
	g.write8(addr+1, upper)
}
