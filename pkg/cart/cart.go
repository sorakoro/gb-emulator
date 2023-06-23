package cart

const (
	ROM = iota
	MBC1
	MBC2
	MBC3
	MBC5
)

type Cartridge struct {
	IsCGB                  bool
	Type, ROMSize, RAMSize byte
	MBC                    int
}

func New(rom []byte) *Cartridge {
	return &Cartridge{
		IsCGB:   rom[0x143] == 0x80 || rom[0x143] == 0xc0,
		Type:    rom[0x0147],
		ROMSize: rom[0x0148],
		RAMSize: rom[0x0149],
	}
}
