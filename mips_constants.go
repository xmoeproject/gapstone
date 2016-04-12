/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst /Users/ben/src/capstone/bindings/python/capstone/
	2016-04-12T16:44:39+09:30

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [mips_const.py]
// Operand type for instruction's operands
const (
	MIPS_OP_INVALID = C.MIPS_OP_INVALID
	MIPS_OP_REG     = C.MIPS_OP_REG
	MIPS_OP_IMM     = C.MIPS_OP_IMM
	MIPS_OP_MEM     = C.MIPS_OP_MEM
)

// MIPS registers
const (
	MIPS_REG_INVALID = C.MIPS_REG_INVALID
)

// General purpose registers
const (
	MIPS_REG_0  = C.MIPS_REG_0
	MIPS_REG_1  = C.MIPS_REG_1
	MIPS_REG_2  = C.MIPS_REG_2
	MIPS_REG_3  = C.MIPS_REG_3
	MIPS_REG_4  = C.MIPS_REG_4
	MIPS_REG_5  = C.MIPS_REG_5
	MIPS_REG_6  = C.MIPS_REG_6
	MIPS_REG_7  = C.MIPS_REG_7
	MIPS_REG_8  = C.MIPS_REG_8
	MIPS_REG_9  = C.MIPS_REG_9
	MIPS_REG_10 = C.MIPS_REG_10
	MIPS_REG_11 = C.MIPS_REG_11
	MIPS_REG_12 = C.MIPS_REG_12
	MIPS_REG_13 = C.MIPS_REG_13
	MIPS_REG_14 = C.MIPS_REG_14
	MIPS_REG_15 = C.MIPS_REG_15
	MIPS_REG_16 = C.MIPS_REG_16
	MIPS_REG_17 = C.MIPS_REG_17
	MIPS_REG_18 = C.MIPS_REG_18
	MIPS_REG_19 = C.MIPS_REG_19
	MIPS_REG_20 = C.MIPS_REG_20
	MIPS_REG_21 = C.MIPS_REG_21
	MIPS_REG_22 = C.MIPS_REG_22
	MIPS_REG_23 = C.MIPS_REG_23
	MIPS_REG_24 = C.MIPS_REG_24
	MIPS_REG_25 = C.MIPS_REG_25
	MIPS_REG_26 = C.MIPS_REG_26
	MIPS_REG_27 = C.MIPS_REG_27
	MIPS_REG_28 = C.MIPS_REG_28
	MIPS_REG_29 = C.MIPS_REG_29
	MIPS_REG_30 = C.MIPS_REG_30
	MIPS_REG_31 = C.MIPS_REG_31
)

// DSP registers
const (
	MIPS_REG_DSPCCOND        = C.MIPS_REG_DSPCCOND
	MIPS_REG_DSPCARRY        = C.MIPS_REG_DSPCARRY
	MIPS_REG_DSPEFI          = C.MIPS_REG_DSPEFI
	MIPS_REG_DSPOUTFLAG      = C.MIPS_REG_DSPOUTFLAG
	MIPS_REG_DSPOUTFLAG16_19 = C.MIPS_REG_DSPOUTFLAG16_19
	MIPS_REG_DSPOUTFLAG20    = C.MIPS_REG_DSPOUTFLAG20
	MIPS_REG_DSPOUTFLAG21    = C.MIPS_REG_DSPOUTFLAG21
	MIPS_REG_DSPOUTFLAG22    = C.MIPS_REG_DSPOUTFLAG22
	MIPS_REG_DSPOUTFLAG23    = C.MIPS_REG_DSPOUTFLAG23
	MIPS_REG_DSPPOS          = C.MIPS_REG_DSPPOS
	MIPS_REG_DSPSCOUNT       = C.MIPS_REG_DSPSCOUNT
)

// ACC registers
const (
	MIPS_REG_AC0 = C.MIPS_REG_AC0
	MIPS_REG_AC1 = C.MIPS_REG_AC1
	MIPS_REG_AC2 = C.MIPS_REG_AC2
	MIPS_REG_AC3 = C.MIPS_REG_AC3
)

// COP registers
const (
	MIPS_REG_CC0 = C.MIPS_REG_CC0
	MIPS_REG_CC1 = C.MIPS_REG_CC1
	MIPS_REG_CC2 = C.MIPS_REG_CC2
	MIPS_REG_CC3 = C.MIPS_REG_CC3
	MIPS_REG_CC4 = C.MIPS_REG_CC4
	MIPS_REG_CC5 = C.MIPS_REG_CC5
	MIPS_REG_CC6 = C.MIPS_REG_CC6
	MIPS_REG_CC7 = C.MIPS_REG_CC7
)

// FPU registers
const (
	MIPS_REG_F0   = C.MIPS_REG_F0
	MIPS_REG_F1   = C.MIPS_REG_F1
	MIPS_REG_F2   = C.MIPS_REG_F2
	MIPS_REG_F3   = C.MIPS_REG_F3
	MIPS_REG_F4   = C.MIPS_REG_F4
	MIPS_REG_F5   = C.MIPS_REG_F5
	MIPS_REG_F6   = C.MIPS_REG_F6
	MIPS_REG_F7   = C.MIPS_REG_F7
	MIPS_REG_F8   = C.MIPS_REG_F8
	MIPS_REG_F9   = C.MIPS_REG_F9
	MIPS_REG_F10  = C.MIPS_REG_F10
	MIPS_REG_F11  = C.MIPS_REG_F11
	MIPS_REG_F12  = C.MIPS_REG_F12
	MIPS_REG_F13  = C.MIPS_REG_F13
	MIPS_REG_F14  = C.MIPS_REG_F14
	MIPS_REG_F15  = C.MIPS_REG_F15
	MIPS_REG_F16  = C.MIPS_REG_F16
	MIPS_REG_F17  = C.MIPS_REG_F17
	MIPS_REG_F18  = C.MIPS_REG_F18
	MIPS_REG_F19  = C.MIPS_REG_F19
	MIPS_REG_F20  = C.MIPS_REG_F20
	MIPS_REG_F21  = C.MIPS_REG_F21
	MIPS_REG_F22  = C.MIPS_REG_F22
	MIPS_REG_F23  = C.MIPS_REG_F23
	MIPS_REG_F24  = C.MIPS_REG_F24
	MIPS_REG_F25  = C.MIPS_REG_F25
	MIPS_REG_F26  = C.MIPS_REG_F26
	MIPS_REG_F27  = C.MIPS_REG_F27
	MIPS_REG_F28  = C.MIPS_REG_F28
	MIPS_REG_F29  = C.MIPS_REG_F29
	MIPS_REG_F30  = C.MIPS_REG_F30
	MIPS_REG_F31  = C.MIPS_REG_F31
	MIPS_REG_FCC0 = C.MIPS_REG_FCC0
	MIPS_REG_FCC1 = C.MIPS_REG_FCC1
	MIPS_REG_FCC2 = C.MIPS_REG_FCC2
	MIPS_REG_FCC3 = C.MIPS_REG_FCC3
	MIPS_REG_FCC4 = C.MIPS_REG_FCC4
	MIPS_REG_FCC5 = C.MIPS_REG_FCC5
	MIPS_REG_FCC6 = C.MIPS_REG_FCC6
	MIPS_REG_FCC7 = C.MIPS_REG_FCC7
)

// AFPR128
const (
	MIPS_REG_W0     = C.MIPS_REG_W0
	MIPS_REG_W1     = C.MIPS_REG_W1
	MIPS_REG_W2     = C.MIPS_REG_W2
	MIPS_REG_W3     = C.MIPS_REG_W3
	MIPS_REG_W4     = C.MIPS_REG_W4
	MIPS_REG_W5     = C.MIPS_REG_W5
	MIPS_REG_W6     = C.MIPS_REG_W6
	MIPS_REG_W7     = C.MIPS_REG_W7
	MIPS_REG_W8     = C.MIPS_REG_W8
	MIPS_REG_W9     = C.MIPS_REG_W9
	MIPS_REG_W10    = C.MIPS_REG_W10
	MIPS_REG_W11    = C.MIPS_REG_W11
	MIPS_REG_W12    = C.MIPS_REG_W12
	MIPS_REG_W13    = C.MIPS_REG_W13
	MIPS_REG_W14    = C.MIPS_REG_W14
	MIPS_REG_W15    = C.MIPS_REG_W15
	MIPS_REG_W16    = C.MIPS_REG_W16
	MIPS_REG_W17    = C.MIPS_REG_W17
	MIPS_REG_W18    = C.MIPS_REG_W18
	MIPS_REG_W19    = C.MIPS_REG_W19
	MIPS_REG_W20    = C.MIPS_REG_W20
	MIPS_REG_W21    = C.MIPS_REG_W21
	MIPS_REG_W22    = C.MIPS_REG_W22
	MIPS_REG_W23    = C.MIPS_REG_W23
	MIPS_REG_W24    = C.MIPS_REG_W24
	MIPS_REG_W25    = C.MIPS_REG_W25
	MIPS_REG_W26    = C.MIPS_REG_W26
	MIPS_REG_W27    = C.MIPS_REG_W27
	MIPS_REG_W28    = C.MIPS_REG_W28
	MIPS_REG_W29    = C.MIPS_REG_W29
	MIPS_REG_W30    = C.MIPS_REG_W30
	MIPS_REG_W31    = C.MIPS_REG_W31
	MIPS_REG_HI     = C.MIPS_REG_HI
	MIPS_REG_LO     = C.MIPS_REG_LO
	MIPS_REG_P0     = C.MIPS_REG_P0
	MIPS_REG_P1     = C.MIPS_REG_P1
	MIPS_REG_P2     = C.MIPS_REG_P2
	MIPS_REG_MPL0   = C.MIPS_REG_MPL0
	MIPS_REG_MPL1   = C.MIPS_REG_MPL1
	MIPS_REG_MPL2   = C.MIPS_REG_MPL2
	MIPS_REG_ENDING = C.MIPS_REG_ENDING
	MIPS_REG_ZERO   = C.MIPS_REG_ZERO
	MIPS_REG_AT     = C.MIPS_REG_AT
	MIPS_REG_V0     = C.MIPS_REG_V0
	MIPS_REG_V1     = C.MIPS_REG_V1
	MIPS_REG_A0     = C.MIPS_REG_A0
	MIPS_REG_A1     = C.MIPS_REG_A1
	MIPS_REG_A2     = C.MIPS_REG_A2
	MIPS_REG_A3     = C.MIPS_REG_A3
	MIPS_REG_T0     = C.MIPS_REG_T0
	MIPS_REG_T1     = C.MIPS_REG_T1
	MIPS_REG_T2     = C.MIPS_REG_T2
	MIPS_REG_T3     = C.MIPS_REG_T3
	MIPS_REG_T4     = C.MIPS_REG_T4
	MIPS_REG_T5     = C.MIPS_REG_T5
	MIPS_REG_T6     = C.MIPS_REG_T6
	MIPS_REG_T7     = C.MIPS_REG_T7
	MIPS_REG_S0     = C.MIPS_REG_S0
	MIPS_REG_S1     = C.MIPS_REG_S1
	MIPS_REG_S2     = C.MIPS_REG_S2
	MIPS_REG_S3     = C.MIPS_REG_S3
	MIPS_REG_S4     = C.MIPS_REG_S4
	MIPS_REG_S5     = C.MIPS_REG_S5
	MIPS_REG_S6     = C.MIPS_REG_S6
	MIPS_REG_S7     = C.MIPS_REG_S7
	MIPS_REG_T8     = C.MIPS_REG_T8
	MIPS_REG_T9     = C.MIPS_REG_T9
	MIPS_REG_K0     = C.MIPS_REG_K0
	MIPS_REG_K1     = C.MIPS_REG_K1
	MIPS_REG_GP     = C.MIPS_REG_GP
	MIPS_REG_SP     = C.MIPS_REG_SP
	MIPS_REG_FP     = C.MIPS_REG_FP
	MIPS_REG_S8     = C.MIPS_REG_S8
	MIPS_REG_RA     = C.MIPS_REG_RA
	MIPS_REG_HI0    = C.MIPS_REG_HI0
	MIPS_REG_HI1    = C.MIPS_REG_HI1
	MIPS_REG_HI2    = C.MIPS_REG_HI2
	MIPS_REG_HI3    = C.MIPS_REG_HI3
	MIPS_REG_LO0    = C.MIPS_REG_LO0
	MIPS_REG_LO1    = C.MIPS_REG_LO1
	MIPS_REG_LO2    = C.MIPS_REG_LO2
	MIPS_REG_LO3    = C.MIPS_REG_LO3
)

// MIPS instruction
const (
	MIPS_INS_INVALID     = C.MIPS_INS_INVALID
	MIPS_INS_ABSQ_S      = C.MIPS_INS_ABSQ_S
	MIPS_INS_ADD         = C.MIPS_INS_ADD
	MIPS_INS_ADDIUPC     = C.MIPS_INS_ADDIUPC
	MIPS_INS_ADDQH       = C.MIPS_INS_ADDQH
	MIPS_INS_ADDQH_R     = C.MIPS_INS_ADDQH_R
	MIPS_INS_ADDQ        = C.MIPS_INS_ADDQ
	MIPS_INS_ADDQ_S      = C.MIPS_INS_ADDQ_S
	MIPS_INS_ADDSC       = C.MIPS_INS_ADDSC
	MIPS_INS_ADDS_A      = C.MIPS_INS_ADDS_A
	MIPS_INS_ADDS_S      = C.MIPS_INS_ADDS_S
	MIPS_INS_ADDS_U      = C.MIPS_INS_ADDS_U
	MIPS_INS_ADDUH       = C.MIPS_INS_ADDUH
	MIPS_INS_ADDUH_R     = C.MIPS_INS_ADDUH_R
	MIPS_INS_ADDU        = C.MIPS_INS_ADDU
	MIPS_INS_ADDU_S      = C.MIPS_INS_ADDU_S
	MIPS_INS_ADDVI       = C.MIPS_INS_ADDVI
	MIPS_INS_ADDV        = C.MIPS_INS_ADDV
	MIPS_INS_ADDWC       = C.MIPS_INS_ADDWC
	MIPS_INS_ADD_A       = C.MIPS_INS_ADD_A
	MIPS_INS_ADDI        = C.MIPS_INS_ADDI
	MIPS_INS_ADDIU       = C.MIPS_INS_ADDIU
	MIPS_INS_ALIGN       = C.MIPS_INS_ALIGN
	MIPS_INS_ALUIPC      = C.MIPS_INS_ALUIPC
	MIPS_INS_AND         = C.MIPS_INS_AND
	MIPS_INS_ANDI        = C.MIPS_INS_ANDI
	MIPS_INS_APPEND      = C.MIPS_INS_APPEND
	MIPS_INS_ASUB_S      = C.MIPS_INS_ASUB_S
	MIPS_INS_ASUB_U      = C.MIPS_INS_ASUB_U
	MIPS_INS_AUI         = C.MIPS_INS_AUI
	MIPS_INS_AUIPC       = C.MIPS_INS_AUIPC
	MIPS_INS_AVER_S      = C.MIPS_INS_AVER_S
	MIPS_INS_AVER_U      = C.MIPS_INS_AVER_U
	MIPS_INS_AVE_S       = C.MIPS_INS_AVE_S
	MIPS_INS_AVE_U       = C.MIPS_INS_AVE_U
	MIPS_INS_BADDU       = C.MIPS_INS_BADDU
	MIPS_INS_BAL         = C.MIPS_INS_BAL
	MIPS_INS_BALC        = C.MIPS_INS_BALC
	MIPS_INS_BALIGN      = C.MIPS_INS_BALIGN
	MIPS_INS_BC          = C.MIPS_INS_BC
	MIPS_INS_BC0F        = C.MIPS_INS_BC0F
	MIPS_INS_BC0FL       = C.MIPS_INS_BC0FL
	MIPS_INS_BC0T        = C.MIPS_INS_BC0T
	MIPS_INS_BC0TL       = C.MIPS_INS_BC0TL
	MIPS_INS_BC1EQZ      = C.MIPS_INS_BC1EQZ
	MIPS_INS_BC1F        = C.MIPS_INS_BC1F
	MIPS_INS_BC1FL       = C.MIPS_INS_BC1FL
	MIPS_INS_BC1NEZ      = C.MIPS_INS_BC1NEZ
	MIPS_INS_BC1T        = C.MIPS_INS_BC1T
	MIPS_INS_BC1TL       = C.MIPS_INS_BC1TL
	MIPS_INS_BC2EQZ      = C.MIPS_INS_BC2EQZ
	MIPS_INS_BC2F        = C.MIPS_INS_BC2F
	MIPS_INS_BC2FL       = C.MIPS_INS_BC2FL
	MIPS_INS_BC2NEZ      = C.MIPS_INS_BC2NEZ
	MIPS_INS_BC2T        = C.MIPS_INS_BC2T
	MIPS_INS_BC2TL       = C.MIPS_INS_BC2TL
	MIPS_INS_BC3F        = C.MIPS_INS_BC3F
	MIPS_INS_BC3FL       = C.MIPS_INS_BC3FL
	MIPS_INS_BC3T        = C.MIPS_INS_BC3T
	MIPS_INS_BC3TL       = C.MIPS_INS_BC3TL
	MIPS_INS_BCLRI       = C.MIPS_INS_BCLRI
	MIPS_INS_BCLR        = C.MIPS_INS_BCLR
	MIPS_INS_BEQ         = C.MIPS_INS_BEQ
	MIPS_INS_BEQC        = C.MIPS_INS_BEQC
	MIPS_INS_BEQL        = C.MIPS_INS_BEQL
	MIPS_INS_BEQZALC     = C.MIPS_INS_BEQZALC
	MIPS_INS_BEQZC       = C.MIPS_INS_BEQZC
	MIPS_INS_BGEC        = C.MIPS_INS_BGEC
	MIPS_INS_BGEUC       = C.MIPS_INS_BGEUC
	MIPS_INS_BGEZ        = C.MIPS_INS_BGEZ
	MIPS_INS_BGEZAL      = C.MIPS_INS_BGEZAL
	MIPS_INS_BGEZALC     = C.MIPS_INS_BGEZALC
	MIPS_INS_BGEZALL     = C.MIPS_INS_BGEZALL
	MIPS_INS_BGEZALS     = C.MIPS_INS_BGEZALS
	MIPS_INS_BGEZC       = C.MIPS_INS_BGEZC
	MIPS_INS_BGEZL       = C.MIPS_INS_BGEZL
	MIPS_INS_BGTZ        = C.MIPS_INS_BGTZ
	MIPS_INS_BGTZALC     = C.MIPS_INS_BGTZALC
	MIPS_INS_BGTZC       = C.MIPS_INS_BGTZC
	MIPS_INS_BGTZL       = C.MIPS_INS_BGTZL
	MIPS_INS_BINSLI      = C.MIPS_INS_BINSLI
	MIPS_INS_BINSL       = C.MIPS_INS_BINSL
	MIPS_INS_BINSRI      = C.MIPS_INS_BINSRI
	MIPS_INS_BINSR       = C.MIPS_INS_BINSR
	MIPS_INS_BITREV      = C.MIPS_INS_BITREV
	MIPS_INS_BITSWAP     = C.MIPS_INS_BITSWAP
	MIPS_INS_BLEZ        = C.MIPS_INS_BLEZ
	MIPS_INS_BLEZALC     = C.MIPS_INS_BLEZALC
	MIPS_INS_BLEZC       = C.MIPS_INS_BLEZC
	MIPS_INS_BLEZL       = C.MIPS_INS_BLEZL
	MIPS_INS_BLTC        = C.MIPS_INS_BLTC
	MIPS_INS_BLTUC       = C.MIPS_INS_BLTUC
	MIPS_INS_BLTZ        = C.MIPS_INS_BLTZ
	MIPS_INS_BLTZAL      = C.MIPS_INS_BLTZAL
	MIPS_INS_BLTZALC     = C.MIPS_INS_BLTZALC
	MIPS_INS_BLTZALL     = C.MIPS_INS_BLTZALL
	MIPS_INS_BLTZALS     = C.MIPS_INS_BLTZALS
	MIPS_INS_BLTZC       = C.MIPS_INS_BLTZC
	MIPS_INS_BLTZL       = C.MIPS_INS_BLTZL
	MIPS_INS_BMNZI       = C.MIPS_INS_BMNZI
	MIPS_INS_BMNZ        = C.MIPS_INS_BMNZ
	MIPS_INS_BMZI        = C.MIPS_INS_BMZI
	MIPS_INS_BMZ         = C.MIPS_INS_BMZ
	MIPS_INS_BNE         = C.MIPS_INS_BNE
	MIPS_INS_BNEC        = C.MIPS_INS_BNEC
	MIPS_INS_BNEGI       = C.MIPS_INS_BNEGI
	MIPS_INS_BNEG        = C.MIPS_INS_BNEG
	MIPS_INS_BNEL        = C.MIPS_INS_BNEL
	MIPS_INS_BNEZALC     = C.MIPS_INS_BNEZALC
	MIPS_INS_BNEZC       = C.MIPS_INS_BNEZC
	MIPS_INS_BNVC        = C.MIPS_INS_BNVC
	MIPS_INS_BNZ         = C.MIPS_INS_BNZ
	MIPS_INS_BOVC        = C.MIPS_INS_BOVC
	MIPS_INS_BPOSGE32    = C.MIPS_INS_BPOSGE32
	MIPS_INS_BREAK       = C.MIPS_INS_BREAK
	MIPS_INS_BSELI       = C.MIPS_INS_BSELI
	MIPS_INS_BSEL        = C.MIPS_INS_BSEL
	MIPS_INS_BSETI       = C.MIPS_INS_BSETI
	MIPS_INS_BSET        = C.MIPS_INS_BSET
	MIPS_INS_BZ          = C.MIPS_INS_BZ
	MIPS_INS_BEQZ        = C.MIPS_INS_BEQZ
	MIPS_INS_B           = C.MIPS_INS_B
	MIPS_INS_BNEZ        = C.MIPS_INS_BNEZ
	MIPS_INS_BTEQZ       = C.MIPS_INS_BTEQZ
	MIPS_INS_BTNEZ       = C.MIPS_INS_BTNEZ
	MIPS_INS_CACHE       = C.MIPS_INS_CACHE
	MIPS_INS_CEIL        = C.MIPS_INS_CEIL
	MIPS_INS_CEQI        = C.MIPS_INS_CEQI
	MIPS_INS_CEQ         = C.MIPS_INS_CEQ
	MIPS_INS_CFC1        = C.MIPS_INS_CFC1
	MIPS_INS_CFCMSA      = C.MIPS_INS_CFCMSA
	MIPS_INS_CINS        = C.MIPS_INS_CINS
	MIPS_INS_CINS32      = C.MIPS_INS_CINS32
	MIPS_INS_CLASS       = C.MIPS_INS_CLASS
	MIPS_INS_CLEI_S      = C.MIPS_INS_CLEI_S
	MIPS_INS_CLEI_U      = C.MIPS_INS_CLEI_U
	MIPS_INS_CLE_S       = C.MIPS_INS_CLE_S
	MIPS_INS_CLE_U       = C.MIPS_INS_CLE_U
	MIPS_INS_CLO         = C.MIPS_INS_CLO
	MIPS_INS_CLTI_S      = C.MIPS_INS_CLTI_S
	MIPS_INS_CLTI_U      = C.MIPS_INS_CLTI_U
	MIPS_INS_CLT_S       = C.MIPS_INS_CLT_S
	MIPS_INS_CLT_U       = C.MIPS_INS_CLT_U
	MIPS_INS_CLZ         = C.MIPS_INS_CLZ
	MIPS_INS_CMPGDU      = C.MIPS_INS_CMPGDU
	MIPS_INS_CMPGU       = C.MIPS_INS_CMPGU
	MIPS_INS_CMPU        = C.MIPS_INS_CMPU
	MIPS_INS_CMP         = C.MIPS_INS_CMP
	MIPS_INS_COPY_S      = C.MIPS_INS_COPY_S
	MIPS_INS_COPY_U      = C.MIPS_INS_COPY_U
	MIPS_INS_CTC1        = C.MIPS_INS_CTC1
	MIPS_INS_CTCMSA      = C.MIPS_INS_CTCMSA
	MIPS_INS_CVT         = C.MIPS_INS_CVT
	MIPS_INS_C           = C.MIPS_INS_C
	MIPS_INS_CMPI        = C.MIPS_INS_CMPI
	MIPS_INS_DADD        = C.MIPS_INS_DADD
	MIPS_INS_DADDI       = C.MIPS_INS_DADDI
	MIPS_INS_DADDIU      = C.MIPS_INS_DADDIU
	MIPS_INS_DADDU       = C.MIPS_INS_DADDU
	MIPS_INS_DAHI        = C.MIPS_INS_DAHI
	MIPS_INS_DALIGN      = C.MIPS_INS_DALIGN
	MIPS_INS_DATI        = C.MIPS_INS_DATI
	MIPS_INS_DAUI        = C.MIPS_INS_DAUI
	MIPS_INS_DBITSWAP    = C.MIPS_INS_DBITSWAP
	MIPS_INS_DCLO        = C.MIPS_INS_DCLO
	MIPS_INS_DCLZ        = C.MIPS_INS_DCLZ
	MIPS_INS_DDIV        = C.MIPS_INS_DDIV
	MIPS_INS_DDIVU       = C.MIPS_INS_DDIVU
	MIPS_INS_DERET       = C.MIPS_INS_DERET
	MIPS_INS_DEXT        = C.MIPS_INS_DEXT
	MIPS_INS_DEXTM       = C.MIPS_INS_DEXTM
	MIPS_INS_DEXTU       = C.MIPS_INS_DEXTU
	MIPS_INS_DI          = C.MIPS_INS_DI
	MIPS_INS_DINS        = C.MIPS_INS_DINS
	MIPS_INS_DINSM       = C.MIPS_INS_DINSM
	MIPS_INS_DINSU       = C.MIPS_INS_DINSU
	MIPS_INS_DIV         = C.MIPS_INS_DIV
	MIPS_INS_DIVU        = C.MIPS_INS_DIVU
	MIPS_INS_DIV_S       = C.MIPS_INS_DIV_S
	MIPS_INS_DIV_U       = C.MIPS_INS_DIV_U
	MIPS_INS_DLSA        = C.MIPS_INS_DLSA
	MIPS_INS_DMFC0       = C.MIPS_INS_DMFC0
	MIPS_INS_DMFC1       = C.MIPS_INS_DMFC1
	MIPS_INS_DMFC2       = C.MIPS_INS_DMFC2
	MIPS_INS_DMOD        = C.MIPS_INS_DMOD
	MIPS_INS_DMODU       = C.MIPS_INS_DMODU
	MIPS_INS_DMTC0       = C.MIPS_INS_DMTC0
	MIPS_INS_DMTC1       = C.MIPS_INS_DMTC1
	MIPS_INS_DMTC2       = C.MIPS_INS_DMTC2
	MIPS_INS_DMUH        = C.MIPS_INS_DMUH
	MIPS_INS_DMUHU       = C.MIPS_INS_DMUHU
	MIPS_INS_DMUL        = C.MIPS_INS_DMUL
	MIPS_INS_DMULT       = C.MIPS_INS_DMULT
	MIPS_INS_DMULTU      = C.MIPS_INS_DMULTU
	MIPS_INS_DMULU       = C.MIPS_INS_DMULU
	MIPS_INS_DOTP_S      = C.MIPS_INS_DOTP_S
	MIPS_INS_DOTP_U      = C.MIPS_INS_DOTP_U
	MIPS_INS_DPADD_S     = C.MIPS_INS_DPADD_S
	MIPS_INS_DPADD_U     = C.MIPS_INS_DPADD_U
	MIPS_INS_DPAQX_SA    = C.MIPS_INS_DPAQX_SA
	MIPS_INS_DPAQX_S     = C.MIPS_INS_DPAQX_S
	MIPS_INS_DPAQ_SA     = C.MIPS_INS_DPAQ_SA
	MIPS_INS_DPAQ_S      = C.MIPS_INS_DPAQ_S
	MIPS_INS_DPAU        = C.MIPS_INS_DPAU
	MIPS_INS_DPAX        = C.MIPS_INS_DPAX
	MIPS_INS_DPA         = C.MIPS_INS_DPA
	MIPS_INS_DPOP        = C.MIPS_INS_DPOP
	MIPS_INS_DPSQX_SA    = C.MIPS_INS_DPSQX_SA
	MIPS_INS_DPSQX_S     = C.MIPS_INS_DPSQX_S
	MIPS_INS_DPSQ_SA     = C.MIPS_INS_DPSQ_SA
	MIPS_INS_DPSQ_S      = C.MIPS_INS_DPSQ_S
	MIPS_INS_DPSUB_S     = C.MIPS_INS_DPSUB_S
	MIPS_INS_DPSUB_U     = C.MIPS_INS_DPSUB_U
	MIPS_INS_DPSU        = C.MIPS_INS_DPSU
	MIPS_INS_DPSX        = C.MIPS_INS_DPSX
	MIPS_INS_DPS         = C.MIPS_INS_DPS
	MIPS_INS_DROTR       = C.MIPS_INS_DROTR
	MIPS_INS_DROTR32     = C.MIPS_INS_DROTR32
	MIPS_INS_DROTRV      = C.MIPS_INS_DROTRV
	MIPS_INS_DSBH        = C.MIPS_INS_DSBH
	MIPS_INS_DSHD        = C.MIPS_INS_DSHD
	MIPS_INS_DSLL        = C.MIPS_INS_DSLL
	MIPS_INS_DSLL32      = C.MIPS_INS_DSLL32
	MIPS_INS_DSLLV       = C.MIPS_INS_DSLLV
	MIPS_INS_DSRA        = C.MIPS_INS_DSRA
	MIPS_INS_DSRA32      = C.MIPS_INS_DSRA32
	MIPS_INS_DSRAV       = C.MIPS_INS_DSRAV
	MIPS_INS_DSRL        = C.MIPS_INS_DSRL
	MIPS_INS_DSRL32      = C.MIPS_INS_DSRL32
	MIPS_INS_DSRLV       = C.MIPS_INS_DSRLV
	MIPS_INS_DSUB        = C.MIPS_INS_DSUB
	MIPS_INS_DSUBU       = C.MIPS_INS_DSUBU
	MIPS_INS_EHB         = C.MIPS_INS_EHB
	MIPS_INS_EI          = C.MIPS_INS_EI
	MIPS_INS_ERET        = C.MIPS_INS_ERET
	MIPS_INS_EXT         = C.MIPS_INS_EXT
	MIPS_INS_EXTP        = C.MIPS_INS_EXTP
	MIPS_INS_EXTPDP      = C.MIPS_INS_EXTPDP
	MIPS_INS_EXTPDPV     = C.MIPS_INS_EXTPDPV
	MIPS_INS_EXTPV       = C.MIPS_INS_EXTPV
	MIPS_INS_EXTRV_RS    = C.MIPS_INS_EXTRV_RS
	MIPS_INS_EXTRV_R     = C.MIPS_INS_EXTRV_R
	MIPS_INS_EXTRV_S     = C.MIPS_INS_EXTRV_S
	MIPS_INS_EXTRV       = C.MIPS_INS_EXTRV
	MIPS_INS_EXTR_RS     = C.MIPS_INS_EXTR_RS
	MIPS_INS_EXTR_R      = C.MIPS_INS_EXTR_R
	MIPS_INS_EXTR_S      = C.MIPS_INS_EXTR_S
	MIPS_INS_EXTR        = C.MIPS_INS_EXTR
	MIPS_INS_EXTS        = C.MIPS_INS_EXTS
	MIPS_INS_EXTS32      = C.MIPS_INS_EXTS32
	MIPS_INS_ABS         = C.MIPS_INS_ABS
	MIPS_INS_FADD        = C.MIPS_INS_FADD
	MIPS_INS_FCAF        = C.MIPS_INS_FCAF
	MIPS_INS_FCEQ        = C.MIPS_INS_FCEQ
	MIPS_INS_FCLASS      = C.MIPS_INS_FCLASS
	MIPS_INS_FCLE        = C.MIPS_INS_FCLE
	MIPS_INS_FCLT        = C.MIPS_INS_FCLT
	MIPS_INS_FCNE        = C.MIPS_INS_FCNE
	MIPS_INS_FCOR        = C.MIPS_INS_FCOR
	MIPS_INS_FCUEQ       = C.MIPS_INS_FCUEQ
	MIPS_INS_FCULE       = C.MIPS_INS_FCULE
	MIPS_INS_FCULT       = C.MIPS_INS_FCULT
	MIPS_INS_FCUNE       = C.MIPS_INS_FCUNE
	MIPS_INS_FCUN        = C.MIPS_INS_FCUN
	MIPS_INS_FDIV        = C.MIPS_INS_FDIV
	MIPS_INS_FEXDO       = C.MIPS_INS_FEXDO
	MIPS_INS_FEXP2       = C.MIPS_INS_FEXP2
	MIPS_INS_FEXUPL      = C.MIPS_INS_FEXUPL
	MIPS_INS_FEXUPR      = C.MIPS_INS_FEXUPR
	MIPS_INS_FFINT_S     = C.MIPS_INS_FFINT_S
	MIPS_INS_FFINT_U     = C.MIPS_INS_FFINT_U
	MIPS_INS_FFQL        = C.MIPS_INS_FFQL
	MIPS_INS_FFQR        = C.MIPS_INS_FFQR
	MIPS_INS_FILL        = C.MIPS_INS_FILL
	MIPS_INS_FLOG2       = C.MIPS_INS_FLOG2
	MIPS_INS_FLOOR       = C.MIPS_INS_FLOOR
	MIPS_INS_FMADD       = C.MIPS_INS_FMADD
	MIPS_INS_FMAX_A      = C.MIPS_INS_FMAX_A
	MIPS_INS_FMAX        = C.MIPS_INS_FMAX
	MIPS_INS_FMIN_A      = C.MIPS_INS_FMIN_A
	MIPS_INS_FMIN        = C.MIPS_INS_FMIN
	MIPS_INS_MOV         = C.MIPS_INS_MOV
	MIPS_INS_FMSUB       = C.MIPS_INS_FMSUB
	MIPS_INS_FMUL        = C.MIPS_INS_FMUL
	MIPS_INS_MUL         = C.MIPS_INS_MUL
	MIPS_INS_NEG         = C.MIPS_INS_NEG
	MIPS_INS_FRCP        = C.MIPS_INS_FRCP
	MIPS_INS_FRINT       = C.MIPS_INS_FRINT
	MIPS_INS_FRSQRT      = C.MIPS_INS_FRSQRT
	MIPS_INS_FSAF        = C.MIPS_INS_FSAF
	MIPS_INS_FSEQ        = C.MIPS_INS_FSEQ
	MIPS_INS_FSLE        = C.MIPS_INS_FSLE
	MIPS_INS_FSLT        = C.MIPS_INS_FSLT
	MIPS_INS_FSNE        = C.MIPS_INS_FSNE
	MIPS_INS_FSOR        = C.MIPS_INS_FSOR
	MIPS_INS_FSQRT       = C.MIPS_INS_FSQRT
	MIPS_INS_SQRT        = C.MIPS_INS_SQRT
	MIPS_INS_FSUB        = C.MIPS_INS_FSUB
	MIPS_INS_SUB         = C.MIPS_INS_SUB
	MIPS_INS_FSUEQ       = C.MIPS_INS_FSUEQ
	MIPS_INS_FSULE       = C.MIPS_INS_FSULE
	MIPS_INS_FSULT       = C.MIPS_INS_FSULT
	MIPS_INS_FSUNE       = C.MIPS_INS_FSUNE
	MIPS_INS_FSUN        = C.MIPS_INS_FSUN
	MIPS_INS_FTINT_S     = C.MIPS_INS_FTINT_S
	MIPS_INS_FTINT_U     = C.MIPS_INS_FTINT_U
	MIPS_INS_FTQ         = C.MIPS_INS_FTQ
	MIPS_INS_FTRUNC_S    = C.MIPS_INS_FTRUNC_S
	MIPS_INS_FTRUNC_U    = C.MIPS_INS_FTRUNC_U
	MIPS_INS_HADD_S      = C.MIPS_INS_HADD_S
	MIPS_INS_HADD_U      = C.MIPS_INS_HADD_U
	MIPS_INS_HSUB_S      = C.MIPS_INS_HSUB_S
	MIPS_INS_HSUB_U      = C.MIPS_INS_HSUB_U
	MIPS_INS_ILVEV       = C.MIPS_INS_ILVEV
	MIPS_INS_ILVL        = C.MIPS_INS_ILVL
	MIPS_INS_ILVOD       = C.MIPS_INS_ILVOD
	MIPS_INS_ILVR        = C.MIPS_INS_ILVR
	MIPS_INS_INS         = C.MIPS_INS_INS
	MIPS_INS_INSERT      = C.MIPS_INS_INSERT
	MIPS_INS_INSV        = C.MIPS_INS_INSV
	MIPS_INS_INSVE       = C.MIPS_INS_INSVE
	MIPS_INS_J           = C.MIPS_INS_J
	MIPS_INS_JAL         = C.MIPS_INS_JAL
	MIPS_INS_JALR        = C.MIPS_INS_JALR
	MIPS_INS_JALRS       = C.MIPS_INS_JALRS
	MIPS_INS_JALS        = C.MIPS_INS_JALS
	MIPS_INS_JALX        = C.MIPS_INS_JALX
	MIPS_INS_JIALC       = C.MIPS_INS_JIALC
	MIPS_INS_JIC         = C.MIPS_INS_JIC
	MIPS_INS_JR          = C.MIPS_INS_JR
	MIPS_INS_JRADDIUSP   = C.MIPS_INS_JRADDIUSP
	MIPS_INS_JRC         = C.MIPS_INS_JRC
	MIPS_INS_JALRC       = C.MIPS_INS_JALRC
	MIPS_INS_LB          = C.MIPS_INS_LB
	MIPS_INS_LBUX        = C.MIPS_INS_LBUX
	MIPS_INS_LBU         = C.MIPS_INS_LBU
	MIPS_INS_LD          = C.MIPS_INS_LD
	MIPS_INS_LDC1        = C.MIPS_INS_LDC1
	MIPS_INS_LDC2        = C.MIPS_INS_LDC2
	MIPS_INS_LDC3        = C.MIPS_INS_LDC3
	MIPS_INS_LDI         = C.MIPS_INS_LDI
	MIPS_INS_LDL         = C.MIPS_INS_LDL
	MIPS_INS_LDPC        = C.MIPS_INS_LDPC
	MIPS_INS_LDR         = C.MIPS_INS_LDR
	MIPS_INS_LDXC1       = C.MIPS_INS_LDXC1
	MIPS_INS_LH          = C.MIPS_INS_LH
	MIPS_INS_LHX         = C.MIPS_INS_LHX
	MIPS_INS_LHU         = C.MIPS_INS_LHU
	MIPS_INS_LL          = C.MIPS_INS_LL
	MIPS_INS_LLD         = C.MIPS_INS_LLD
	MIPS_INS_LSA         = C.MIPS_INS_LSA
	MIPS_INS_LUXC1       = C.MIPS_INS_LUXC1
	MIPS_INS_LUI         = C.MIPS_INS_LUI
	MIPS_INS_LW          = C.MIPS_INS_LW
	MIPS_INS_LWC1        = C.MIPS_INS_LWC1
	MIPS_INS_LWC2        = C.MIPS_INS_LWC2
	MIPS_INS_LWC3        = C.MIPS_INS_LWC3
	MIPS_INS_LWL         = C.MIPS_INS_LWL
	MIPS_INS_LWPC        = C.MIPS_INS_LWPC
	MIPS_INS_LWR         = C.MIPS_INS_LWR
	MIPS_INS_LWUPC       = C.MIPS_INS_LWUPC
	MIPS_INS_LWU         = C.MIPS_INS_LWU
	MIPS_INS_LWX         = C.MIPS_INS_LWX
	MIPS_INS_LWXC1       = C.MIPS_INS_LWXC1
	MIPS_INS_LI          = C.MIPS_INS_LI
	MIPS_INS_MADD        = C.MIPS_INS_MADD
	MIPS_INS_MADDF       = C.MIPS_INS_MADDF
	MIPS_INS_MADDR_Q     = C.MIPS_INS_MADDR_Q
	MIPS_INS_MADDU       = C.MIPS_INS_MADDU
	MIPS_INS_MADDV       = C.MIPS_INS_MADDV
	MIPS_INS_MADD_Q      = C.MIPS_INS_MADD_Q
	MIPS_INS_MAQ_SA      = C.MIPS_INS_MAQ_SA
	MIPS_INS_MAQ_S       = C.MIPS_INS_MAQ_S
	MIPS_INS_MAXA        = C.MIPS_INS_MAXA
	MIPS_INS_MAXI_S      = C.MIPS_INS_MAXI_S
	MIPS_INS_MAXI_U      = C.MIPS_INS_MAXI_U
	MIPS_INS_MAX_A       = C.MIPS_INS_MAX_A
	MIPS_INS_MAX         = C.MIPS_INS_MAX
	MIPS_INS_MAX_S       = C.MIPS_INS_MAX_S
	MIPS_INS_MAX_U       = C.MIPS_INS_MAX_U
	MIPS_INS_MFC0        = C.MIPS_INS_MFC0
	MIPS_INS_MFC1        = C.MIPS_INS_MFC1
	MIPS_INS_MFC2        = C.MIPS_INS_MFC2
	MIPS_INS_MFHC1       = C.MIPS_INS_MFHC1
	MIPS_INS_MFHI        = C.MIPS_INS_MFHI
	MIPS_INS_MFLO        = C.MIPS_INS_MFLO
	MIPS_INS_MINA        = C.MIPS_INS_MINA
	MIPS_INS_MINI_S      = C.MIPS_INS_MINI_S
	MIPS_INS_MINI_U      = C.MIPS_INS_MINI_U
	MIPS_INS_MIN_A       = C.MIPS_INS_MIN_A
	MIPS_INS_MIN         = C.MIPS_INS_MIN
	MIPS_INS_MIN_S       = C.MIPS_INS_MIN_S
	MIPS_INS_MIN_U       = C.MIPS_INS_MIN_U
	MIPS_INS_MOD         = C.MIPS_INS_MOD
	MIPS_INS_MODSUB      = C.MIPS_INS_MODSUB
	MIPS_INS_MODU        = C.MIPS_INS_MODU
	MIPS_INS_MOD_S       = C.MIPS_INS_MOD_S
	MIPS_INS_MOD_U       = C.MIPS_INS_MOD_U
	MIPS_INS_MOVE        = C.MIPS_INS_MOVE
	MIPS_INS_MOVF        = C.MIPS_INS_MOVF
	MIPS_INS_MOVN        = C.MIPS_INS_MOVN
	MIPS_INS_MOVT        = C.MIPS_INS_MOVT
	MIPS_INS_MOVZ        = C.MIPS_INS_MOVZ
	MIPS_INS_MSUB        = C.MIPS_INS_MSUB
	MIPS_INS_MSUBF       = C.MIPS_INS_MSUBF
	MIPS_INS_MSUBR_Q     = C.MIPS_INS_MSUBR_Q
	MIPS_INS_MSUBU       = C.MIPS_INS_MSUBU
	MIPS_INS_MSUBV       = C.MIPS_INS_MSUBV
	MIPS_INS_MSUB_Q      = C.MIPS_INS_MSUB_Q
	MIPS_INS_MTC0        = C.MIPS_INS_MTC0
	MIPS_INS_MTC1        = C.MIPS_INS_MTC1
	MIPS_INS_MTC2        = C.MIPS_INS_MTC2
	MIPS_INS_MTHC1       = C.MIPS_INS_MTHC1
	MIPS_INS_MTHI        = C.MIPS_INS_MTHI
	MIPS_INS_MTHLIP      = C.MIPS_INS_MTHLIP
	MIPS_INS_MTLO        = C.MIPS_INS_MTLO
	MIPS_INS_MTM0        = C.MIPS_INS_MTM0
	MIPS_INS_MTM1        = C.MIPS_INS_MTM1
	MIPS_INS_MTM2        = C.MIPS_INS_MTM2
	MIPS_INS_MTP0        = C.MIPS_INS_MTP0
	MIPS_INS_MTP1        = C.MIPS_INS_MTP1
	MIPS_INS_MTP2        = C.MIPS_INS_MTP2
	MIPS_INS_MUH         = C.MIPS_INS_MUH
	MIPS_INS_MUHU        = C.MIPS_INS_MUHU
	MIPS_INS_MULEQ_S     = C.MIPS_INS_MULEQ_S
	MIPS_INS_MULEU_S     = C.MIPS_INS_MULEU_S
	MIPS_INS_MULQ_RS     = C.MIPS_INS_MULQ_RS
	MIPS_INS_MULQ_S      = C.MIPS_INS_MULQ_S
	MIPS_INS_MULR_Q      = C.MIPS_INS_MULR_Q
	MIPS_INS_MULSAQ_S    = C.MIPS_INS_MULSAQ_S
	MIPS_INS_MULSA       = C.MIPS_INS_MULSA
	MIPS_INS_MULT        = C.MIPS_INS_MULT
	MIPS_INS_MULTU       = C.MIPS_INS_MULTU
	MIPS_INS_MULU        = C.MIPS_INS_MULU
	MIPS_INS_MULV        = C.MIPS_INS_MULV
	MIPS_INS_MUL_Q       = C.MIPS_INS_MUL_Q
	MIPS_INS_MUL_S       = C.MIPS_INS_MUL_S
	MIPS_INS_NLOC        = C.MIPS_INS_NLOC
	MIPS_INS_NLZC        = C.MIPS_INS_NLZC
	MIPS_INS_NMADD       = C.MIPS_INS_NMADD
	MIPS_INS_NMSUB       = C.MIPS_INS_NMSUB
	MIPS_INS_NOR         = C.MIPS_INS_NOR
	MIPS_INS_NORI        = C.MIPS_INS_NORI
	MIPS_INS_NOT         = C.MIPS_INS_NOT
	MIPS_INS_OR          = C.MIPS_INS_OR
	MIPS_INS_ORI         = C.MIPS_INS_ORI
	MIPS_INS_PACKRL      = C.MIPS_INS_PACKRL
	MIPS_INS_PAUSE       = C.MIPS_INS_PAUSE
	MIPS_INS_PCKEV       = C.MIPS_INS_PCKEV
	MIPS_INS_PCKOD       = C.MIPS_INS_PCKOD
	MIPS_INS_PCNT        = C.MIPS_INS_PCNT
	MIPS_INS_PICK        = C.MIPS_INS_PICK
	MIPS_INS_POP         = C.MIPS_INS_POP
	MIPS_INS_PRECEQU     = C.MIPS_INS_PRECEQU
	MIPS_INS_PRECEQ      = C.MIPS_INS_PRECEQ
	MIPS_INS_PRECEU      = C.MIPS_INS_PRECEU
	MIPS_INS_PRECRQU_S   = C.MIPS_INS_PRECRQU_S
	MIPS_INS_PRECRQ      = C.MIPS_INS_PRECRQ
	MIPS_INS_PRECRQ_RS   = C.MIPS_INS_PRECRQ_RS
	MIPS_INS_PRECR       = C.MIPS_INS_PRECR
	MIPS_INS_PRECR_SRA   = C.MIPS_INS_PRECR_SRA
	MIPS_INS_PRECR_SRA_R = C.MIPS_INS_PRECR_SRA_R
	MIPS_INS_PREF        = C.MIPS_INS_PREF
	MIPS_INS_PREPEND     = C.MIPS_INS_PREPEND
	MIPS_INS_RADDU       = C.MIPS_INS_RADDU
	MIPS_INS_RDDSP       = C.MIPS_INS_RDDSP
	MIPS_INS_RDHWR       = C.MIPS_INS_RDHWR
	MIPS_INS_REPLV       = C.MIPS_INS_REPLV
	MIPS_INS_REPL        = C.MIPS_INS_REPL
	MIPS_INS_RINT        = C.MIPS_INS_RINT
	MIPS_INS_ROTR        = C.MIPS_INS_ROTR
	MIPS_INS_ROTRV       = C.MIPS_INS_ROTRV
	MIPS_INS_ROUND       = C.MIPS_INS_ROUND
	MIPS_INS_SAT_S       = C.MIPS_INS_SAT_S
	MIPS_INS_SAT_U       = C.MIPS_INS_SAT_U
	MIPS_INS_SB          = C.MIPS_INS_SB
	MIPS_INS_SC          = C.MIPS_INS_SC
	MIPS_INS_SCD         = C.MIPS_INS_SCD
	MIPS_INS_SD          = C.MIPS_INS_SD
	MIPS_INS_SDBBP       = C.MIPS_INS_SDBBP
	MIPS_INS_SDC1        = C.MIPS_INS_SDC1
	MIPS_INS_SDC2        = C.MIPS_INS_SDC2
	MIPS_INS_SDC3        = C.MIPS_INS_SDC3
	MIPS_INS_SDL         = C.MIPS_INS_SDL
	MIPS_INS_SDR         = C.MIPS_INS_SDR
	MIPS_INS_SDXC1       = C.MIPS_INS_SDXC1
	MIPS_INS_SEB         = C.MIPS_INS_SEB
	MIPS_INS_SEH         = C.MIPS_INS_SEH
	MIPS_INS_SELEQZ      = C.MIPS_INS_SELEQZ
	MIPS_INS_SELNEZ      = C.MIPS_INS_SELNEZ
	MIPS_INS_SEL         = C.MIPS_INS_SEL
	MIPS_INS_SEQ         = C.MIPS_INS_SEQ
	MIPS_INS_SEQI        = C.MIPS_INS_SEQI
	MIPS_INS_SH          = C.MIPS_INS_SH
	MIPS_INS_SHF         = C.MIPS_INS_SHF
	MIPS_INS_SHILO       = C.MIPS_INS_SHILO
	MIPS_INS_SHILOV      = C.MIPS_INS_SHILOV
	MIPS_INS_SHLLV       = C.MIPS_INS_SHLLV
	MIPS_INS_SHLLV_S     = C.MIPS_INS_SHLLV_S
	MIPS_INS_SHLL        = C.MIPS_INS_SHLL
	MIPS_INS_SHLL_S      = C.MIPS_INS_SHLL_S
	MIPS_INS_SHRAV       = C.MIPS_INS_SHRAV
	MIPS_INS_SHRAV_R     = C.MIPS_INS_SHRAV_R
	MIPS_INS_SHRA        = C.MIPS_INS_SHRA
	MIPS_INS_SHRA_R      = C.MIPS_INS_SHRA_R
	MIPS_INS_SHRLV       = C.MIPS_INS_SHRLV
	MIPS_INS_SHRL        = C.MIPS_INS_SHRL
	MIPS_INS_SLDI        = C.MIPS_INS_SLDI
	MIPS_INS_SLD         = C.MIPS_INS_SLD
	MIPS_INS_SLL         = C.MIPS_INS_SLL
	MIPS_INS_SLLI        = C.MIPS_INS_SLLI
	MIPS_INS_SLLV        = C.MIPS_INS_SLLV
	MIPS_INS_SLT         = C.MIPS_INS_SLT
	MIPS_INS_SLTI        = C.MIPS_INS_SLTI
	MIPS_INS_SLTIU       = C.MIPS_INS_SLTIU
	MIPS_INS_SLTU        = C.MIPS_INS_SLTU
	MIPS_INS_SNE         = C.MIPS_INS_SNE
	MIPS_INS_SNEI        = C.MIPS_INS_SNEI
	MIPS_INS_SPLATI      = C.MIPS_INS_SPLATI
	MIPS_INS_SPLAT       = C.MIPS_INS_SPLAT
	MIPS_INS_SRA         = C.MIPS_INS_SRA
	MIPS_INS_SRAI        = C.MIPS_INS_SRAI
	MIPS_INS_SRARI       = C.MIPS_INS_SRARI
	MIPS_INS_SRAR        = C.MIPS_INS_SRAR
	MIPS_INS_SRAV        = C.MIPS_INS_SRAV
	MIPS_INS_SRL         = C.MIPS_INS_SRL
	MIPS_INS_SRLI        = C.MIPS_INS_SRLI
	MIPS_INS_SRLRI       = C.MIPS_INS_SRLRI
	MIPS_INS_SRLR        = C.MIPS_INS_SRLR
	MIPS_INS_SRLV        = C.MIPS_INS_SRLV
	MIPS_INS_SSNOP       = C.MIPS_INS_SSNOP
	MIPS_INS_ST          = C.MIPS_INS_ST
	MIPS_INS_SUBQH       = C.MIPS_INS_SUBQH
	MIPS_INS_SUBQH_R     = C.MIPS_INS_SUBQH_R
	MIPS_INS_SUBQ        = C.MIPS_INS_SUBQ
	MIPS_INS_SUBQ_S      = C.MIPS_INS_SUBQ_S
	MIPS_INS_SUBSUS_U    = C.MIPS_INS_SUBSUS_U
	MIPS_INS_SUBSUU_S    = C.MIPS_INS_SUBSUU_S
	MIPS_INS_SUBS_S      = C.MIPS_INS_SUBS_S
	MIPS_INS_SUBS_U      = C.MIPS_INS_SUBS_U
	MIPS_INS_SUBUH       = C.MIPS_INS_SUBUH
	MIPS_INS_SUBUH_R     = C.MIPS_INS_SUBUH_R
	MIPS_INS_SUBU        = C.MIPS_INS_SUBU
	MIPS_INS_SUBU_S      = C.MIPS_INS_SUBU_S
	MIPS_INS_SUBVI       = C.MIPS_INS_SUBVI
	MIPS_INS_SUBV        = C.MIPS_INS_SUBV
	MIPS_INS_SUXC1       = C.MIPS_INS_SUXC1
	MIPS_INS_SW          = C.MIPS_INS_SW
	MIPS_INS_SWC1        = C.MIPS_INS_SWC1
	MIPS_INS_SWC2        = C.MIPS_INS_SWC2
	MIPS_INS_SWC3        = C.MIPS_INS_SWC3
	MIPS_INS_SWL         = C.MIPS_INS_SWL
	MIPS_INS_SWR         = C.MIPS_INS_SWR
	MIPS_INS_SWXC1       = C.MIPS_INS_SWXC1
	MIPS_INS_SYNC        = C.MIPS_INS_SYNC
	MIPS_INS_SYSCALL     = C.MIPS_INS_SYSCALL
	MIPS_INS_TEQ         = C.MIPS_INS_TEQ
	MIPS_INS_TEQI        = C.MIPS_INS_TEQI
	MIPS_INS_TGE         = C.MIPS_INS_TGE
	MIPS_INS_TGEI        = C.MIPS_INS_TGEI
	MIPS_INS_TGEIU       = C.MIPS_INS_TGEIU
	MIPS_INS_TGEU        = C.MIPS_INS_TGEU
	MIPS_INS_TLBP        = C.MIPS_INS_TLBP
	MIPS_INS_TLBR        = C.MIPS_INS_TLBR
	MIPS_INS_TLBWI       = C.MIPS_INS_TLBWI
	MIPS_INS_TLBWR       = C.MIPS_INS_TLBWR
	MIPS_INS_TLT         = C.MIPS_INS_TLT
	MIPS_INS_TLTI        = C.MIPS_INS_TLTI
	MIPS_INS_TLTIU       = C.MIPS_INS_TLTIU
	MIPS_INS_TLTU        = C.MIPS_INS_TLTU
	MIPS_INS_TNE         = C.MIPS_INS_TNE
	MIPS_INS_TNEI        = C.MIPS_INS_TNEI
	MIPS_INS_TRUNC       = C.MIPS_INS_TRUNC
	MIPS_INS_V3MULU      = C.MIPS_INS_V3MULU
	MIPS_INS_VMM0        = C.MIPS_INS_VMM0
	MIPS_INS_VMULU       = C.MIPS_INS_VMULU
	MIPS_INS_VSHF        = C.MIPS_INS_VSHF
	MIPS_INS_WAIT        = C.MIPS_INS_WAIT
	MIPS_INS_WRDSP       = C.MIPS_INS_WRDSP
	MIPS_INS_WSBH        = C.MIPS_INS_WSBH
	MIPS_INS_XOR         = C.MIPS_INS_XOR
	MIPS_INS_XORI        = C.MIPS_INS_XORI
)

// some alias instructions
const (
	MIPS_INS_NOP  = C.MIPS_INS_NOP
	MIPS_INS_NEGU = C.MIPS_INS_NEGU
)

// special instructions
const (
	MIPS_INS_JALR_HB = C.MIPS_INS_JALR_HB
	MIPS_INS_JR_HB   = C.MIPS_INS_JR_HB
	MIPS_INS_ENDING  = C.MIPS_INS_ENDING
)

// Group of MIPS instructions
const (
	MIPS_GRP_INVALID = C.MIPS_GRP_INVALID
)

// Generic groups
const (
	MIPS_GRP_JUMP = C.MIPS_GRP_JUMP
)

// Architecture-specific groups
const (
	MIPS_GRP_BITCOUNT       = C.MIPS_GRP_BITCOUNT
	MIPS_GRP_DSP            = C.MIPS_GRP_DSP
	MIPS_GRP_DSPR2          = C.MIPS_GRP_DSPR2
	MIPS_GRP_FPIDX          = C.MIPS_GRP_FPIDX
	MIPS_GRP_MSA            = C.MIPS_GRP_MSA
	MIPS_GRP_MIPS32R2       = C.MIPS_GRP_MIPS32R2
	MIPS_GRP_MIPS64         = C.MIPS_GRP_MIPS64
	MIPS_GRP_MIPS64R2       = C.MIPS_GRP_MIPS64R2
	MIPS_GRP_SEINREG        = C.MIPS_GRP_SEINREG
	MIPS_GRP_STDENC         = C.MIPS_GRP_STDENC
	MIPS_GRP_SWAP           = C.MIPS_GRP_SWAP
	MIPS_GRP_MICROMIPS      = C.MIPS_GRP_MICROMIPS
	MIPS_GRP_MIPS16MODE     = C.MIPS_GRP_MIPS16MODE
	MIPS_GRP_FP64BIT        = C.MIPS_GRP_FP64BIT
	MIPS_GRP_NONANSFPMATH   = C.MIPS_GRP_NONANSFPMATH
	MIPS_GRP_NOTFP64BIT     = C.MIPS_GRP_NOTFP64BIT
	MIPS_GRP_NOTINMICROMIPS = C.MIPS_GRP_NOTINMICROMIPS
	MIPS_GRP_NOTNACL        = C.MIPS_GRP_NOTNACL
	MIPS_GRP_NOTMIPS32R6    = C.MIPS_GRP_NOTMIPS32R6
	MIPS_GRP_NOTMIPS64R6    = C.MIPS_GRP_NOTMIPS64R6
	MIPS_GRP_CNMIPS         = C.MIPS_GRP_CNMIPS
	MIPS_GRP_MIPS32         = C.MIPS_GRP_MIPS32
	MIPS_GRP_MIPS32R6       = C.MIPS_GRP_MIPS32R6
	MIPS_GRP_MIPS64R6       = C.MIPS_GRP_MIPS64R6
	MIPS_GRP_MIPS2          = C.MIPS_GRP_MIPS2
	MIPS_GRP_MIPS3          = C.MIPS_GRP_MIPS3
	MIPS_GRP_MIPS3_32       = C.MIPS_GRP_MIPS3_32
	MIPS_GRP_MIPS3_32R2     = C.MIPS_GRP_MIPS3_32R2
	MIPS_GRP_MIPS4_32       = C.MIPS_GRP_MIPS4_32
	MIPS_GRP_MIPS4_32R2     = C.MIPS_GRP_MIPS4_32R2
	MIPS_GRP_MIPS5_32R2     = C.MIPS_GRP_MIPS5_32R2
	MIPS_GRP_GP32BIT        = C.MIPS_GRP_GP32BIT
	MIPS_GRP_GP64BIT        = C.MIPS_GRP_GP64BIT
	MIPS_GRP_ENDING         = C.MIPS_GRP_ENDING
)
