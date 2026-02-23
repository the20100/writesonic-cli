package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/mattn/go-isatty"
)

// IsJSON returns true when output should be JSON (piped or explicit flags).
func IsJSON(jsonFlag, prettyFlag bool) bool {
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		return true
	}
	return jsonFlag || prettyFlag
}

// PrintJSON encodes v as JSON to stdout.
func PrintJSON(v any, pretty bool) error {
	enc := json.NewEncoder(os.Stdout)
	if pretty {
		enc.SetIndent("", "  ")
	}
	return enc.Encode(v)
}

// PrintTable prints tab-aligned rows with a header row.
func PrintTable(headers []string, rows [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(headers, "\t"))
	for _, row := range rows {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}

// PrintKeyValue prints two-column key/value rows.
func PrintKeyValue(rows [][]string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	for _, row := range rows {
		if len(row) < 2 || row[1] == "" || row[1] == "-" {
			continue
		}
		fmt.Fprintf(w, "%s\t%s\n", row[0], row[1])
	}
	w.Flush()
}

// PrintText prints each result as a numbered block of text.
func PrintText(results []string) {
	for i, t := range results {
		if len(results) > 1 {
			fmt.Printf("--- Result %d ---\n", i+1)
		}
		fmt.Println(t)
		if i < len(results)-1 {
			fmt.Println()
		}
	}
}

// Truncate shortens s to maxLen with an ellipsis if needed.
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-1] + "…"
}
