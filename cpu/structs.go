package cpu

type SYMBOL struct {
	forward map[LONGWORD]LONGWORD // map location to displacement
	flags   LONGWORD
	value   LONGWORD
	name    string
}

type BREAKSTR struct {
	pc    LONGWORD
	kind  LONGWORD
	after LONGWORD
}

type OPCODE struct {
	extended    BYTE
	function    BYTE
	index       WORD
	access      [6]LONGWORD
	count       BYTE
	address     [6]LONGWORD
	VAXaddr     [6]LONGWORD
	size        [6]LONGWORD
	is_register [6]LONGWORD
	regnum      [6]LONGWORD
	name        string
}
