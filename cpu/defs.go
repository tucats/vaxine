package cpu

const (
	AsmDialectAny = 0
	AsmDialectM32 = 1
	AsmDialectGAS = 2

	K_NOFROWARD = -1
	K_ADDR_SIZE = 0x000F
	K_ADDR_MODE = 0x00F0

	K_ADDR_BASE = 0x0000
	K_ADDR_B    = 0x0001
	K_ADDR_W    = 0x0002
	K_ADDR_L    = 0x0004
	K_ADDR_Q    = 0x0008

	K_DISP_BASE = 0x0010

	K_DISP_B = 0x0011
	K_DISP_W = 0x0012
	K_DISP_L = 0x0014
	K_DISP_Q = 0x0018

	K_BRANCH_BASE = 0x0020

	K_BRANCH_B = 0x0021
	K_BRANCH_W = 0x0022
	K_BRANCH_L = 0x0024
	K_BRANCH_Q = 0x0028

	K_CASE_BASE = 0x0040
	K_CASE_W    = 0x0042

	//  Operand decoded locations, used to determine accessability for operands

	OP_MEMORY    = 0x00 //  Operand reference memory location (MBZ!)
	OP_REGISTER  = 0x01 //  Operand references VAX register
	OP_TEMPORARY = 0x03 //  Operand references temporary register
	//      which means readonly and non-addressable

	//  Machine exception names

	EXC_UNUSED     = 0x00
	EXC_CHECK      = 0x04 //  Machine check
	EXC_KSNV       = 0x08 //  Kernel stack not valid
	EXC_POWER      = 0x0C //  Power failure
	EXC_PRIV       = 0x10 //  Priveleged or reserved fault
	EXC_CUSTOMER   = 0x14 //  Customer reserved instruction
	EXC_RESOP      = 0x18 //  Reserved operand fault
	EXC_RESADDR    = 0x1C //  Reserved addressing mode fault
	EXC_ACCVIO     = 0x20 //  Access control violation
	EXC_TNV        = 0x24 //  Translation not valid
	EXC_TP         = 0x28 //  Trace pending fault
	EXC_BPT        = 0x2C //  Breakpoint fault
	EXC_COMPAT     = 0x30 //  Compatability mode fault (PDP-11)
	EXC_ARITH      = 0x34 //  Arithmetic fault
	EXC_CHMK       = 0x40 //  Change mode to kernel
	EXC_CHME       = 0x44 //  Change mode to executive
	EXC_CHMS       = 0x48 //  Change mode to supervisor
	EXC_CHMU       = 0x4C //  Change mode to user
	EXC_SBI        = 0x50 //  SBI SILO COMPARE
	EXC_CMRD       = 0x54 //  Corrected Memory read data
	EXC_SBIALERT   = 0x58 //  SBI Alert
	EXC_SBIFAULT   = 0x5C //  SBI fault
	EXC_MWT        = 0x60 //  Memory write timeout
	EXC_SOFTWARE1  = 0x84 //  Software level 1
	EXC_SOFTWARE2  = 0x88 //  Software level 2
	EXC_SOFTWARE3  = 0x8C //  Software level 3
	EXC_SOFTWARE4  = 0x90 //  Software level 4
	EXC_SOFTWARE5  = 0x94 //  Software level 5
	EXC_SOFTWARE6  = 0x98 //  Software level 6
	EXC_SOFTWARE7  = 0x9C //  Software level 7
	EXC_SOFTWARE8  = 0xA0 //  Software level 8
	EXC_SOFTWARE9  = 0xA4 //  Software level 9
	EXC_SOFTWARE10 = 0xA8 //  Software level 10
	EXC_SOFTWARE11 = 0xAC //  Software level 11
	EXC_SOFTWARE12 = 0xB0 //  Software level 12
	EXC_SOFTWARE13 = 0xB4 //  Software level 13
	EXC_SOFTWARE14 = 0xB8 //  Software level 14
	EXC_SOFTWARE15 = 0xBC //  Software level 15
	EXC_INTERVAL   = 0xC0 //  Interval timer
	EXC_CONREAD    = 0xF8 //  Console read
	EXC_CONWRITE   = 0xFC //  Console write

	//  Flags to describe misc console functions

	CONSOLE_EXPAND  = 0x0001 //  Expand symbol substitutions in command strings
	CONSOLE_VERBOSE = 0x0002 //  Prattle on about things as we do them?

	//  Flags to describe symbols

	SYM_NONE      = 0      // No flags set
	SYM_ENTRY     = 0x0001 // This symbol is an entry point
	SYM_FREFS     = 0x0002 // This symbol has forward refs
	SYM_LABEL     = 0x0004 // This symbols is a label
	SYM_LOCAL     = 0x0008 // Symbol has local scope to entry
	SYM_PERMANENT = 0x0010 // Symbol is permanent.
	SYM_SYSTEM    = 0x0020 // Symbol is a system symbol
	SYM_STRING    = 0x0040 // Symbol is a string, value is addr

	//  Characteristics of the virtual VAX machine state

	MAXREG = 63 // Number of registers, 0-15 architecturally defined, the rest are temporaries, etc.

	MAXPRIVREG = 128 // Number of privileged registers

	//  Addressing mode values, used in handling opcode operands and
	//  privileged register access modes.

	OP_NL = 0 // Operand not used
	OP_RD = 1 // Operand is readonly
	OP_WR = 2 // Operand is writeonly
	OP_MD = 3 // Operand is updated in place
	OP_AD = 4 // Operand is an address
	OP_VA = 5 // Operand is a variable bitfield address
	OP_BR = 6 // Operand is a displacement branch
	OP_IM = 7 // Operand is an implicit immediate operand

	//  Access mode definitions for PSL, page tables, etc.

	AM_KERNEL     = 0
	AM_EXECUTIVE  = 1
	AM_SUPERVISOR = 2
	AM_USER       = 3

	//  Misc debug flags

	DBG_SYMBOLS    = 0x00000001 // Symbol table debugging
	DBG_DEBUG      = 0x00000002 // Invoke the native debugger if possible
	DBG_MEMORY     = 0x00000004 // Debug memory
	DBG_INTERRUPTS = 0x00000008
	DBG_EXCEPTIONS = 0x00000010
	DBG_VM         = 0x00000020 // Debug virtual memory mapping
	DBG_TB         = 0x00000040 // Debug translation caching
	DBG_FUNCTIONS  = 0x00000080
	DBG_REGISTERS  = 0x00000100
	DBG_IMAGES     = 0x00000200 // Display image info on RUN command?
	DBG_USERHALT   = 0x00000400 // Will HALT halt or RESOP when not in kernel mode?
	DBG_KBD        = 0x00000800 // Debug console keyboard input
	DBG_SERVICES   = 0x00001000 // Debug system services
	DBG_DCL        = 0x00002000 // Debug DCL parsing
	DBG_UNIMP      = 0x00004000 // Debug unimplemented instructions?
	DBG_CHM        = 0x00008000 // Debug change mode operations
	DBG_EXPAND     = 0x00010000 // Command line expansion enabled?
	DBG_LOGICALS   = 0x00020000 // Debug logical name operations?
	DBG_DEVICES    = 0x00040000 // Debug device operations?
	DBG_LIBINIT    = 0x00080000 // Run LIB$INITIALIZE by default?
	DBG_PROCESS    = 0x00100000 // Debug process services?
	DBG_FULLDISASM = 0x00200000 // Disasm includes operands addresses?
	DBG_RMS        = 0x00400000 // Debug RMS calls?
	DBG_P1         = 0x01000000 // Misc debugging flag 1
	DBG_P2         = 0x02000000 // Misc debugging flag 1
	DBG_P3         = 0x04000000 // Misc debugging flag 1
	DBG_P4         = 0x08000000 // Misc debugging flag 1

	//  Define the register mnemonics.  These are used as "vax.NAME" where "vax" is
	//  the handle to the virtual machine and "NAME" is the name of the register.

	KSP    = 0
	ESP    = 1
	SSP    = 2
	USP    = 3
	ISP    = 4
	P0BR   = 8
	P0LR   = 9
	P1BR   = 10
	P1LR   = 11
	SBR    = 12
	SLR    = 13
	PCBB   = 16
	SCBB   = 17
	IPL    = 18
	ASTLVL = 19
	SIRR   = 20
	SISR   = 21
	ICCS   = 24
	NICR   = 25
	ICR    = 26
	TODR   = 27
	RXCS   = 32
	RXDB   = 33
	TXCS   = 34
	TXDB   = 35
	TBDR   = 36 // 11/730 specific
	MAPEN  = 56
	TBIA   = 57
	TBIS   = 58
	PMR    = 61
	SID    = 62
	TBCHK  = 63

	//  Mnemonics for the first 6 temporary registers, used for
	//  character instructions.

	T0 = 16
	T1 = 17
	T2 = 18
	T3 = 19
	T4 = 20
	T5 = 21

	//  Mnemonics for the standard register file

	R15 = 15
	R14 = 14
	R13 = 13
	R12 = 12
	R11 = 11
	R10 = 10
	R9  = 9
	R8  = 8
	R7  = 7
	R6  = 6
	R5  = 5
	R4  = 4
	R3  = 3
	R2  = 2
	R1  = 1
	R0  = 0

	//  Overloaded mnemonics for the pre-defined registers

	PC = 15
	SP = 14
	FP = 13
	AP = 12

	//  Types of breakpoints

	BREAK_ADDRESS     = 0    //  Break on an address
	BREAK_FAULT       = 1    //  Break on exception/interrupt
	BREAK_INSTRUCTION = 2    //  Break on a specific opcode
	BREAK_TYPE        = 0x0F //  Mask for finding type code
	BREAK_TEMPORARY   = 0x80 //  Temporary breakpoint
	BREAK_STEP        = 0x40 //  STEP-based breakpoint

	//  Miscellaneous constants

	CONSOLE_COMMENT   = ';'  //  Default comment delimiter
	CONSOLE_SEPARATOR = 0xFF //  No command separator

	//  Types of STEP modes

	STEP_NONE        = 0
	STEP_INSTRUCTION = 1
	STEP_OVER        = 2
	STEP_RETURN      = 3
)

