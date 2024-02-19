package palette

// Palette: https://catppuccin.ryanccn.dev/palette
type Palette struct {
	Base      string
	Black     string
	Blue      string
	Crust     string
	Flamingo  string
	Green     string
	Lavender  string
	Mantle    string
	Maroon    string
	Mauve     string
	Overlay0  string
	Overlay1  string
	Overlay2  string
	Peach     string
	Pink      string
	Red       string
	Rosewater string
	Sapphire  string
	Sky       string
	Subtext0  string
	Subtext1  string
	Surface0  string
	Surface1  string
	Surface2  string
	Teal      string
	Text      string
	Yellow    string
}

var (
	Primary   = Mocha.Text
	Secondary = Mocha.Overlay0
)

var Mocha = Palette{
	Base:      "#1e1e2e",
	Black:     "#11111b",
	Blue:      "#89b4fa",
	Crust:     "#11111b",
	Flamingo:  "#f2cdcd",
	Green:     "#a6e3a1",
	Lavender:  "#b4befe",
	Mantle:    "#181825",
	Maroon:    "#eba0ac",
	Mauve:     "#cba6f7",
	Overlay0:  "#6c7086",
	Overlay1:  "#7f849c",
	Overlay2:  "#9399b2",
	Peach:     "#fab387",
	Pink:      "#f5c2e7",
	Red:       "#f38ba8",
	Sapphire:  "#74c7ec",
	Sky:       "#89dceb",
	Subtext0:  "#a6adc8",
	Subtext1:  "#bac2de",
	Surface0:  "#313244",
	Surface1:  "#45475a",
	Surface2:  "#585b70",
	Teal:      "#94e2d5",
	Text:      "#cdd6f4",
	Yellow:    "#f9e2af",
	Rosewater: "#f5e0dc",
}
