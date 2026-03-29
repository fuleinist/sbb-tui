package ui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/necrom4/sbb-tui/config"
)

const (
	// Layout dimensions
	borderSize          = 2
	headerHeight        = 3
	resultMargin        = 1
	simpleConnHeight    = 9
	simpleConnMargin    = 3
	helpBarHeight       = 1
	stopsLineFixedWidth = (borderSize * 2) + (simpleConnMargin * 2) + (2+5)*2 + 6
	stopsLineMinWidth   = 10
	detailPaddingH      = 3
	detailPaddingV      = 1
	minTermWidth        = 80
	minTermHeight       = 24
)

type styles struct {
	text           lipgloss.Style
	errorText      lipgloss.Style
	ghostText      lipgloss.Style
	active         lipgloss.Style
	inactive       lipgloss.Style
	detailedResult lipgloss.Style
	helpKey        lipgloss.Style
	helpDesc       lipgloss.Style
	warning        lipgloss.Style
	warningBold    lipgloss.Style
	vehicleIcon    lipgloss.Style
	vehicleModel   lipgloss.Style
	company        lipgloss.Style
	logo           lipgloss.Style
	bold           lipgloss.Style
}

func newStyles(theme config.Theme) styles {
	return styles{
		text: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Text)),
		errorText: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.ErrorText)),
		ghostText: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.GhostText)),
		active: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(theme.ActiveBorder)).
			Foreground(lipgloss.Color(theme.Text)).
			Padding(0, 1),
		inactive: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(theme.InactiveBorder)).
			Foreground(lipgloss.Color(theme.Text)).
			Padding(0, 1),
		detailedResult: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(theme.ActiveBorder)).
			Padding(detailPaddingV, detailPaddingH),
		helpKey: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(theme.KeysFg)).
			Background(lipgloss.Color(theme.KeysBg)).
			Padding(0, 1),
		helpDesc: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.GhostText)),
		warning: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.WarningFlag)),
		warningBold: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.WarningFlag)).
			Bold(true),
		vehicleIcon: lipgloss.NewStyle().
			Background(lipgloss.Color(theme.VehicleBg)).
			Foreground(lipgloss.Color(theme.VehicleFg)),
		vehicleModel: lipgloss.NewStyle().
			Background(lipgloss.Color(theme.ModelBg)).
			Foreground(lipgloss.Color(theme.ModelFg)).
			Bold(true),
		company: lipgloss.NewStyle().
			Background(lipgloss.Color(theme.CompanyBg)).
			Foreground(lipgloss.Color(theme.CompanyFg)),
		logo: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Logo)),
		bold: lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Text)).
			Bold(true),
	}
}
