package gbc

const (
	INS_NOP = iota
	INS_LD
	INS_INC
	INS_DEC
	INS_RLCA
	INS_ADD
	INS_RRCA
	INS_STOP
	INS_RLA
	INS_JR
	INS_RRA
	INS_DAA
	INS_CPL
	INS_SCF
	INS_CCF
)

const (
	OP_a16_PAREN = iota
	OP_SP
	OP_HL_PAREN
	OP_d8
)

type Instruction struct {
	id                 int
	Operand1, Operand2 int
	Handler            func(*GBC, int, int)
}

var instructions = [256]Instruction{
	/* 0x0x */ {INS_NOP, 0, 0, nop}, {INS_LD, BC, 0, ld16i}, {INS_LD, BC, A, ldm16r}, {INS_INC, BC, 0, inc16}, {INS_INC, B, 0, inc8}, {INS_DEC, B, 0, dec8}, {INS_LD, B, 0, ld8i}, {INS_RLCA, 0, 0, rlca}, {INS_LD, OP_a16_PAREN, OP_SP, op0x08}, {INS_ADD, HL, BC, addHL}, {INS_LD, A, BC, ld8m}, {INS_DEC, BC, 0, dec16}, {INS_INC, C, 0, inc8}, {INS_DEC, C, 0, dec8}, {INS_LD, C, 0, ld8i}, {INS_RRCA, 0, 0, rrca},
	/* 0x1x */ {INS_STOP, 0, 0, stop}, {INS_LD, DE, 0, ld16i}, {INS_LD, DE, A, ldm16r}, {INS_INC, DE, 0, inc16}, {INS_INC, D, 0, inc8}, {INS_DEC, D, 0, dec8}, {INS_LD, D, 0, ld8i}, {INS_RLA, 0, 0, rla}, {INS_JR, 0, 0, jr}, {INS_ADD, HL, DE, addHL}, {INS_LD, A, DE, ld8m}, {INS_DEC, DE, 0, dec16}, {INS_INC, E, 0, inc8}, {INS_DEC, E, 0, dec8}, {INS_LD, E, 0, ld8i}, {INS_RRA, 0, 0, rra},
	/* 0x2x */ {INS_JR, flagZ, 0, jrncc}, {INS_LD, HL, 0, ld16i}, {INS_LD, HLI, A, ldm16r}, {INS_INC, HL, 0, inc16}, {INS_INC, H, 0, inc8}, {INS_DEC, H, 0, dec8}, {INS_LD, H, 0, ld8i}, {INS_DAA, 0, 0, daa}, {INS_JR, flagZ, 0, jrcc}, {INS_ADD, HL, SP, addHL}, {INS_LD, A, HLD, ld8m}, {INS_DEC, SP, 0, dec16}, {INS_INC, A, 0, inc8}, {INS_DEC, A, 0, dec8}, {INS_LD, A, 0, ld8i}, {INS_CPL, 0, 0, cpl},
	/* 0x3x */ {INS_JR, flagC, 0, jrncc}, {INS_LD, SP, 0, ld16i}, {INS_LD, HLD, A, ldm16r}, {INS_INC, SP, 0, inc16}, {INS_INC, 0, 0, incHL}, {INS_DEC, 0, 0, decHL}, {INS_LD, OP_HL_PAREN, OP_d8, op0x36}, {INS_SCF, 0, 0, scf}, {INS_JR, flagC, 0, jrcc}, {INS_ADD, HL, SP, addHL}, {INS_LD, A, HLD, ld8m}, {INS_DEC, SP, 0, dec16}, {INS_INC, A, 0, inc8}, {INS_DEC, A, 0, dec8}, {INS_LD, A, 0, ld8i}, {INS_CCF, 0, 0, ccf},
}
