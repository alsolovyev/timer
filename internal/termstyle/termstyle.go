package termstyle

import "github.com/muesli/termenv"

var (
	profile = termenv.ColorProfile()

	// Palette: https://catppuccin.ryanccn.dev/palette
	primary   = profile.Color("#cdd6f4")
	secondary = profile.Color("#6c7086")
)

func SetPrimaryColor(c string) {
	if c == "" {
		return
	}
	secondary = profile.Color(c)
}

func SetSecondaryColor(c string) {
	if c == "" {
		return
	}
	secondary = profile.Color(c)
}

func ToStyle(s string) termenv.Style {
	return profile.String(s)
}

func ToColor(s string, cl string) string {
	return profile.String(s).Foreground(profile.Color(cl)).String()
}

func ToColorBold(s string, cl string) string {
	return profile.String(s).Bold().Foreground(profile.Color(cl)).String()
}

func ToPrimary(s string) string {
	return profile.String(s).Foreground(primary).String()
}

func ToPrimaryBold(s string) string {
	return profile.String(s).Bold().Foreground(primary).String()
}

func ToSecondary(s string) string {
	return profile.String(s).Foreground(secondary).String()
}

func ToSecondaryBold(s string) string {
	return profile.String(s).Bold().Foreground(secondary).String()
}

func ToBold(s string) string {
	return profile.String(s).Bold().String()
}
