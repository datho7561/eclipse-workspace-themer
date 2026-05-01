package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

var cursor_style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Green)

var selected_style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.BrightBlack).
	Background(lipgloss.Cyan)

var theme_map = map[string]theme{
	"gruvbox-light": {
		foreground:              theme_color{r: 60, g: 56, b: 54},
		background:              theme_color{r: 251, g: 241, b: 199},
		selection_foreground:    theme_color{r: 251, g: 241, b: 199},
		selection_background:    theme_color{r: 146, g: 131, b: 116},
		current_line_background: theme_color{r: 235, g: 219, b: 178},
		line_number_color:       theme_color{r: 168, g: 153, b: 132},
		comment_color:           theme_color{r: 146, g: 131, b: 116},
		field_color:             theme_color{r: 7, g: 102, b: 120},
		keyword_color:           theme_color{r: 143, g: 62, b: 113},
		string_color:            theme_color{r: 66, g: 123, b: 87},
		local_variable_color:    theme_color{r: 121, g: 116, b: 14},
		method_color:            theme_color{r: 175, g: 58, b: 2},
		number_color:            theme_color{r: 60, g: 56, b: 54},
	},
	"monokai": {
		foreground:              theme_color{r: 248, g: 248, b: 242},
		background:              theme_color{r: 39, g: 40, b: 34},
		selection_foreground:    theme_color{r: 166, g: 169, b: 170},
		selection_background:    theme_color{r: 73, g: 72, b: 62},
		current_line_background: theme_color{r: 62, g: 61, b: 50},
		line_number_color:       theme_color{r: 117, g: 113, b: 94},
		comment_color:           theme_color{r: 117, g: 113, b: 94},
		field_color:             theme_color{r: 102, g: 217, b: 239},
		keyword_color:           theme_color{r: 249, g: 38, b: 114},
		string_color:            theme_color{r: 230, g: 219, b: 116},
		local_variable_color:    theme_color{r: 252, g: 152, b: 103},
		method_color:            theme_color{r: 169, g: 220, b: 118},
		number_color:            theme_color{r: 174, g: 129, b: 255},
	},
	"one-dark": {
		foreground:              theme_color{r: 171, g: 178, b: 191},
		background:              theme_color{r: 40, g: 44, b: 52},
		selection_foreground:    theme_color{r: 255, g: 255, b: 255},
		selection_background:    theme_color{r: 62, g: 68, b: 81},
		current_line_background: theme_color{r: 44, g: 49, b: 60},
		line_number_color:       theme_color{r: 73, g: 81, b: 98},
		comment_color:           theme_color{r: 92, g: 99, b: 112},
		field_color:             theme_color{r: 86, g: 182, b: 194},
		keyword_color:           theme_color{r: 224, g: 108, b: 117},
		string_color:            theme_color{r: 229, g: 192, b: 123},
		local_variable_color:    theme_color{r: 209, g: 154, b: 102},
		method_color:            theme_color{r: 152, g: 195, b: 121},
		number_color:            theme_color{r: 198, g: 120, b: 221},
	},
	"solarized-light": {
		foreground:              theme_color{r: 88, g: 110, b: 117},
		background:              theme_color{r: 253, g: 246, b: 227},
		selection_foreground:    theme_color{r: 7, g: 54, b: 66},
		selection_background:    theme_color{r: 147, g: 161, b: 161},
		current_line_background: theme_color{r: 238, g: 232, b: 213},
		line_number_color:       theme_color{r: 101, g: 123, b: 131},
		comment_color:           theme_color{r: 147, g: 161, b: 161},
		field_color:             theme_color{r: 133, g: 153, b: 0},
		keyword_color:           theme_color{r: 7, g: 54, b: 66},
		string_color:            theme_color{r: 42, g: 161, b: 152},
		local_variable_color:    theme_color{r: 181, g: 137, b: 0},
		method_color:            theme_color{r: 38, g: 139, b: 210},
		number_color:            theme_color{r: 108, g: 113, b: 196},
		// TODO: adding seperate "type" themeing that's orange would really help the solarized port
	},
	"rose-pine-dawn": {
		foreground:              theme_color{r: 87, g: 82, b: 121},
		background:              theme_color{r: 250, g: 244, b: 237},
		selection_foreground:    theme_color{r: 87, g: 82, b: 121},
		selection_background:    theme_color{r: 242, g: 233, b: 225},
		current_line_background: theme_color{r: 242, g: 233, b: 225},
		line_number_color:       theme_color{r: 152, g: 147, b: 165},
		comment_color:           theme_color{r: 152, g: 147, b: 165},
		field_color:             theme_color{r: 86, g: 148, b: 159},
		keyword_color:           theme_color{r: 40, g: 105, b: 131},
		string_color:            theme_color{r: 234, g: 157, b: 52},
		local_variable_color:    theme_color{r: 144, g: 122, b: 169},
		method_color:            theme_color{r: 215, g: 130, b: 126},
		number_color:            theme_color{r: 144, g: 122, b: 169},
	},
	"tokyo-night": {
		foreground:              theme_color{r: 169, g: 177, b: 214},
		background:              theme_color{r: 26, g: 27, b: 38},
		selection_foreground:    theme_color{r: 192, g: 202, b: 245},
		selection_background:    theme_color{r: 65, g: 72, b: 104},
		current_line_background: theme_color{r: 36, g: 40, b: 59},
		line_number_color:       theme_color{r: 86, g: 95, b: 137},
		comment_color:           theme_color{r: 86, g: 95, b: 137},
		field_color:             theme_color{r: 125, g: 207, b: 255},
		keyword_color:           theme_color{r: 187, g: 154, b: 247},
		string_color:            theme_color{r: 158, g: 206, b: 106},
		local_variable_color:    theme_color{r: 115, g: 218, b: 202},
		method_color:            theme_color{r: 122, g: 162, b: 247},
		number_color:            theme_color{r: 255, g: 158, b: 100},
	},
}

var eclipse_ui_to_modify = map[string]bool{
	"AbstractTextEditor.Color.Background":                        true,
	"AbstractTextEditor.Color.Background.SystemDefault":          true,
	"AbstractTextEditor.Color.Foreground":                        true,
	"AbstractTextEditor.Color.Foreground.SystemDefault":          true,
	"AbstractTextEditor.Color.SelectionBackground":               true,
	"AbstractTextEditor.Color.SelectionBackground.SystemDefault": true,
	"AbstractTextEditor.Color.SelectionForeground":               true,
	"AbstractTextEditor.Color.SelectionForeground.SystemDefault": true,
	"currentLineColor":            true,
	"eclipse.preferences.version": true,
	"lineNumberColor":             true,
	"printMarginColor":            true,
}

var eclipse_jdt_ui_to_modify = map[string]bool{
	"java_comment_task_tag":                             true,
	"java_doc_default":                                  true,
	"java_doc_keyword":                                  true,
	"java_doc_link":                                     true,
	"java_doc_tag":                                      true,
	"java_keyword":                                      true,
	"java_keyword_return":                               true,
	"java_multi_line_comment":                           true,
	"java_single_line_comment":                          true,
	"java_string":                                       true,
	"semanticHighlighting.annotation.color":             true,
	"semanticHighlighting.field.color":                  true,
	"semanticHighlighting.localVariable.color":          true,
	"semanticHighlighting.restrictedKeywords.color":     true,
	"semanticHighlighting.staticField.color":            true,
	"semanticHighlighting.staticFinalField.color":       true,
	"semanticHighlighting.method.color":                 true,
	"semanticHighlighting.method.enabled":               true,
	"semanticHighlighting.methodDeclarationName.color":  true,
	"semanticHighlighting.staticMethodInvocation.color": true,
	"sourceHoverBackgroundColor":                        true,
	"java_default":                                      true,
	"java_bracket":                                      true,
	"java_operator":                                     true,
	"semanticHighlighting.deprecatedMember.color":       true,
	"semanticHighlighting.number.enabled":               true,
	"semanticHighlighting.number.color":                 true,
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	keys := make([]string, 0, len(theme_map))
	keys = append(keys, "default")
	for k := range theme_map {
		keys = append(keys, k)
	}
	return model{
		choices: keys,
		cursor:  0,
	}
}

type model struct {
	choices []string
	cursor  int
}

type theme struct {
	foreground              theme_color
	background              theme_color
	selection_foreground    theme_color
	selection_background    theme_color
	current_line_background theme_color
	line_number_color       theme_color
	keyword_color           theme_color
	comment_color           theme_color
	string_color            theme_color
	field_color             theme_color
	local_variable_color    theme_color
	method_color            theme_color
	number_color            theme_color
}

type theme_color struct {
	r int
	g int
	b int
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key selects the theme
		case "enter":
			if m.choices[m.cursor] == "default" {
				clearTheme()
			} else {
				setTheme(theme_map[m.choices[m.cursor]])
			}
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func clearTheme() {
	wd, err := os.Getwd()
	check(err)

	// first settings file; o.e.ui
	path1 := filepath.Join(wd, ".metadata", ".plugins", "org.eclipse.core.runtime", ".settings", "org.eclipse.ui.editors.prefs")
	dat, err := os.ReadFile(path1)
	check(err)
	entries := strings.Split(string(dat), "\n")
	new_entries := []string{}
	for i := range entries {
		parsed_line := strings.Split(entries[i], "=")
		if !eclipse_ui_to_modify[parsed_line[0]] {
			new_entries = append(new_entries, entries[i])
		}
	}
	err = os.WriteFile(path1, []byte(strings.Join(new_entries, "\n")), 0644)
	check(err)

	// second settings file; o.e.j.ui
	path2 := filepath.Join(wd, ".metadata", ".plugins", "org.eclipse.core.runtime", ".settings", "org.eclipse.jdt.ui.prefs")
	dat, err = os.ReadFile(path2)
	check(err)
	entries = strings.Split(string(dat), "\n")
	new_entries = []string{}
	for i := range entries {
		parsed_line := strings.Split(entries[i], "=")
		if !eclipse_jdt_ui_to_modify[parsed_line[0]] {
			new_entries = append(new_entries, entries[i])
		}
	}
	err = os.WriteFile(path2, []byte(strings.Join(new_entries, "\n")), 0644)
	check(err)
}

func setTheme(t theme) {
	wd, err := os.Getwd()
	check(err)

	// first settings file; o.e.ui
	path1 := filepath.Join(wd, ".metadata", ".plugins", "org.eclipse.core.runtime", ".settings", "org.eclipse.ui.editors.prefs")
	dat, err := os.ReadFile(path1)
	check(err)
	entries := strings.Split(string(dat), "\n")
	new_entries := []string{}
	for i := range entries {
		parsed_line := strings.Split(entries[i], "=")
		if !eclipse_ui_to_modify[parsed_line[0]] {
			new_entries = append(new_entries, entries[i])
		}
	}
	new_entries = append(new_entries, fmt.Sprintf("AbstractTextEditor.Color.Background=%s", t.background))
	new_entries = append(new_entries, "AbstractTextEditor.Color.Background.SystemDefault=false")
	new_entries = append(new_entries, fmt.Sprintf("AbstractTextEditor.Color.Foreground=%s", t.foreground))
	new_entries = append(new_entries, "AbstractTextEditor.Color.Foreground.SystemDefault=false")
	new_entries = append(new_entries, fmt.Sprintf("AbstractTextEditor.Color.SelectionBackground=%s", t.selection_background))
	new_entries = append(new_entries, "AbstractTextEditor.Color.SelectionBackground.SystemDefault=false")
	new_entries = append(new_entries, fmt.Sprintf("AbstractTextEditor.Color.SelectionForeground=%s", t.selection_foreground))
	new_entries = append(new_entries, "AbstractTextEditor.Color.SelectionForeground.SystemDefault=false")
	new_entries = append(new_entries, fmt.Sprintf("currentLineColor=%s", t.current_line_background))
	new_entries = append(new_entries, fmt.Sprintf("lineNumberColor=%s", t.line_number_color))
	new_entries = append(new_entries, fmt.Sprintf("printMarginColor=%s", t.background))
	err = os.WriteFile(path1, []byte(strings.Join(new_entries, "\n")), 0644)
	check(err)

	// second settings file; o.e.j.ui
	path2 := filepath.Join(wd, ".metadata", ".plugins", "org.eclipse.core.runtime", ".settings", "org.eclipse.jdt.ui.prefs")
	dat, err = os.ReadFile(path2)
	check(err)
	entries = strings.Split(string(dat), "\n")
	new_entries = []string{}
	for i := range entries {
		parsed_line := strings.Split(entries[i], "=")
		if !eclipse_jdt_ui_to_modify[parsed_line[0]] {
			new_entries = append(new_entries, entries[i])
		}
	}
	new_entries = append(new_entries, fmt.Sprintf("java_bracket=%s", t.foreground))
	new_entries = append(new_entries, fmt.Sprintf("java_operator=%s", t.foreground))
	new_entries = append(new_entries, fmt.Sprintf("java_default=%s", t.foreground))
	new_entries = append(new_entries, fmt.Sprintf("java_comment_task_tag=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("java_doc_default=%s", t.comment_color))
	new_entries = append(new_entries, fmt.Sprintf("java_doc_keyword=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("java_doc_link=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("java_doc_tag=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("java_keyword=%s", t.keyword_color))
	new_entries = append(new_entries, fmt.Sprintf("java_keyword_return=%s", t.keyword_color))
	new_entries = append(new_entries, fmt.Sprintf("java_multi_line_comment=%s", t.comment_color))
	new_entries = append(new_entries, fmt.Sprintf("java_single_line_comment=%s", t.comment_color))
	new_entries = append(new_entries, fmt.Sprintf("java_string=%s", t.string_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.annotation.color=%s", t.comment_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.field.color=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.localVariable.color=%s", t.local_variable_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.restrictedKeywords.color=%s", t.keyword_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.staticField.color=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.staticFinalField.color=%s", t.field_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.method.color=%s", t.method_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.deprecatedMember.color=%s", t.foreground))
	new_entries = append(new_entries, "semanticHighlighting.method.enabled=true")
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.methodDeclarationName.color=%s", t.method_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.staticMethodInvocation.color=%s", t.method_color))
	new_entries = append(new_entries, fmt.Sprintf("semanticHighlighting.number.color=%s", t.number_color))
	new_entries = append(new_entries, "semanticHighlighting.number.enabled=true")
	new_entries = append(new_entries, fmt.Sprintf("sourceHoverBackgroundColor=%s", t.background))
	err = os.WriteFile(path2, []byte(strings.Join(new_entries, "\n")), 0644)
	check(err)
}

func (c theme_color) String() string {
	return fmt.Sprintf("%d,%d,%d", c.r, c.g, c.b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (m model) View() tea.View {
	// The header
	s := "Which theme should be used?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		fancified_choice := choice
		if m.cursor == i {
			cursor = ">" // cursor!
			fancified_choice = selected_style.Render(choice)
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor_style.Render(cursor), fancified_choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
