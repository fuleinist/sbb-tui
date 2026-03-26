package main

import (
	"fmt"
	"os"

	"github.com/necrom4/sbb-tui/views"

	tea "github.com/charmbracelet/bubbletea"
	flag "github.com/spf13/pflag"
)

// version is set at build time via ldflags.
var version = "dev"

func main() {
	from := flag.String("from", "", "Pre-fill departure station")
	to := flag.String("to", "", "Pre-fill arrival station")
	date := flag.String("date", "", "Pre-fill date (DD.MM.YYYY)")
	timeStr := flag.String("time", "", "Pre-fill time (HH:MM)")
	arrival := flag.Bool("arrival", false, "Use arrival time instead of departure time")
	noNerdFont := flag.Bool("no-nerdfont", false, "Use Unicode fallback icons instead of Nerd Font icons")
	showVersion := flag.BoolP("version", "v", false, "Print version and exit")

	// --help
	flag.Usage = func() {
		fmt.Println("sbb-tui - Swiss SBB/CFF/FFS timetable app for the terminal")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  sbb-tui [flags]")
		fmt.Println()
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("sbb-tui v%s\n", version)
		os.Exit(0)
	}

	cfg := views.Config{
		From:          *from,
		To:            *to,
		Date:          *date,
		Time:          *timeStr,
		IsArrivalTime: *arrival,
		NoNerdFont:    *noNerdFont,
	}

	m := views.InitialModel(cfg)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "could not run program:", err)
		os.Exit(1)
	}
}
