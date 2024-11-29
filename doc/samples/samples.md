package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/lucasb-eyer/go-colorful"
)

type Document struct {
	Name string
	Time string
}

var docs = []Document{
	{"README.md", "2 minutes ago"},
	{"Example.md", "1 hour ago"},
	{"secrets.md", "1 week ago"},
}

const selected = 1

var purchased = []string{
	"Bananas",
	"Barley",
	"Cashews",
	"Coconut Milk",
	"Dill",
	"Eggs",
	"Fish Cake",
	"Leeks",
	"Papaya",
}

var dimEnumStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240")).
	MarginRight(1)

var highlightedEnumStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10")).
	MarginRight(1)

var faint = lipgloss.NewStyle().Faint(true)

// Tablas
const rowLength = 12

func makeRow(start, end int) []string {
	var row []string
	for i := start; i <= end; i++ {
		row = append(row, fmt.Sprint(i))
		row = append(row, "")
	}
	for i := len(row); i < rowLength; i++ {
		row = append(row, "")
	}
	return row
}

func makeEmptyRow() []string {
	return makeRow(0, -1)
}

// Tree (VIII)

type styles struct {
	base,
	block,
	enumerator,
	dir,
	toggle,
	file lipgloss.Style
}

func defaultStyles() styles {
	var s styles
	s.base = lipgloss.NewStyle().Background(lipgloss.Color("57")).Foreground(lipgloss.Color("225"))
	s.block = s.base.Padding(1, 3).Margin(1, 3).Width(40)
	s.enumerator = s.base.Foreground(lipgloss.Color("212")).PaddingRight(1)
	s.dir = s.base.Inline(true)
	s.toggle = s.base.Foreground(lipgloss.Color("207")).PaddingRight(1)
	s.file = s.base
	return s
}

type dir struct {
	name   string
	open   bool
	styles styles
}

func (d dir) String() string {
	t := d.styles.toggle.Render
	n := d.styles.dir.Render
	if d.open {
		return t("▼") + n(d.name)
	}
	return t("▶") + n(d.name)
}

type file struct {
	name   string
	styles styles
}

func (s file) String() string {
	return s.styles.file.Render(s.name)
}

// Lógica para autocompletar los meses
func autocompleteMonths(input string) []string {
	months := []string{"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December"}
	var matches []string
	for _, month := range months {
		if strings.HasPrefix(strings.ToLower(month), strings.ToLower(input)) {
			matches = append(matches, month)
		}
	}
	return matches
}

// Obtener el índice del mes en la lista, o -1 si no es válido
func getMonthIndex(input string, months []string) int {
	inputLower := strings.ToLower(input)
	for idx, month := range months {
		if strings.ToLower(month) == inputLower {
			return idx
		}
	}
	return -1
}
func getBoardIndex(input string, boards []string) int {
	inputLower := strings.ToLower(input)
	for idx, board := range boards {
		if strings.ToLower(board) == inputLower {
			return idx
		}
	}
	return -1
}

// Verificar si el mes ingresado es válido
func isValidMonth(input string) bool {
	months := []string{"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December"}
	inputLower := strings.ToLower(input)
	for _, month := range months {
		if strings.ToLower(month) == inputLower {
			return true
		}
	}
	return false
}

// Lógica para imprimir el dashboard
// Función para imprimir el dashboard
func printEntryMessage() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("236")).
		Padding(1, 2).
		Margin(1).
		Align(lipgloss.Center)

	// Mensaje principal
	return style.Render("¡Hola, esta es tu app con Lipgloss y Bubble Tea!")
}
func printExitMessage() string {

	// Estilo de despedida
	quitStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Background(lipgloss.Color("0")).
		Padding(1, 2).
		Margin(1).
		Align(lipgloss.Center)

	// Mensaje de despedida
	return quitStyle.Render("¡Adiós, gracias por usar la app!")
}

func duckDuckGooseEnumerator(items list.Items, i int) string {
	if items.At(i).Value() == "Goose" {
		return "Honk →"
	}
	return " "
}

func showListIExample() {
	enumStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00d787")).MarginRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))

	l := list.New("Duck", "Duck", "Duck", "Goose", "Duck").
		ItemStyle(itemStyle).
		EnumeratorStyle(enumStyle).
		Enumerator(duckDuckGooseEnumerator)
	fmt.Println(l)
}
func showListIIExample() {
	selected := 1
	baseStyle := lipgloss.NewStyle().
		MarginBottom(1).
		MarginLeft(1)
	dimColor := lipgloss.Color("250")
	hightlightColor := lipgloss.Color("#EE6FF8")

	l := list.New().
		Enumerator(func(_ list.Items, i int) string {
			if i == selected {
				return "│\n│"
			}
			return " "
		}).
		ItemStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			st := baseStyle
			if selected == i {
				return st.Foreground(hightlightColor)
			}
			return st.Foreground(dimColor)
		}).
		EnumeratorStyleFunc(func(_ list.Items, i int) lipgloss.Style {
			if selected == i {
				return lipgloss.NewStyle().Foreground(hightlightColor)
			}
			return lipgloss.NewStyle().Foreground(dimColor)
		})

	for _, d := range docs {
		l.Item(d.Name)
	}

	fmt.Println()
	fmt.Println(l)
}

func groceryEnumerator(items list.Items, i int) string {
	for _, p := range purchased {
		if items.At(i).Value() == p {
			return "✓"
		}
	}
	return "•"
}

func enumStyleFunc(items list.Items, i int) lipgloss.Style {
	for _, p := range purchased {
		if items.At(i).Value() == p {
			return highlightedEnumStyle
		}
	}
	return dimEnumStyle
}

func itemStyleFunc(items list.Items, i int) lipgloss.Style {
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	for _, p := range purchased {
		if items.At(i).Value() == p {
			return itemStyle.Strikethrough(true)
		}
	}
	return itemStyle
}
func showListIIIExample() {
	l := list.New(
		"Artichoke",
		"Baking Flour", "Bananas", "Barley", "Bean Sprouts",
		"Cashew Apple", "Cashews", "Coconut Milk", "Curry Paste", "Currywurst",
		"Dill", "Dragonfruit", "Dried Shrimp",
		"Eggs",
		"Fish Cake", "Furikake",
		"Jicama",
		"Kohlrabi",
		"Leeks", "Lentils", "Licorice Root",
	).
		Enumerator(groceryEnumerator).
		EnumeratorStyleFunc(enumStyleFunc).
		ItemStyleFunc(itemStyleFunc)

	fmt.Println(l)
}

func showListIVExample() {
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255")).MarginRight(1)

	l := list.New(
		"Glossier",
		"Claire’s Boutique",
		"Nyx",
		"Mac",
		"Milk",
	).
		Enumerator(list.Roman).
		EnumeratorStyle(enumeratorStyle).
		ItemStyle(itemStyle)

	fmt.Println(l)
}

func showListVExample() {
	l := list.New(
		"A",
		"B",
		"C",
		list.New(
			"D",
			"E",
			"F",
		).Enumerator(list.Roman),
		"G",
	)
	fmt.Println(l)
}

func showListVIExample() {
	purple := lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		MarginRight(1)

	pink := lipgloss.NewStyle().
		Foreground(lipgloss.Color("212")).
		MarginRight(1)

	base := lipgloss.NewStyle().
		MarginBottom(1).
		MarginLeft(1)

	faint := lipgloss.NewStyle().Faint(true)

	dim := lipgloss.Color("250")
	highlight := lipgloss.Color("#EE6FF8")

	special := lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	checklistEnumStyle := func(items list.Items, index int) lipgloss.Style {
		switch index {
		case 1, 2, 4:
			return lipgloss.NewStyle().
				Foreground(special).
				PaddingRight(1)
		default:
			return lipgloss.NewStyle().PaddingRight(1)
		}
	}

	checklistEnum := func(items list.Items, index int) string {
		switch index {
		case 1, 2, 4:
			return "✓"
		default:
			return "•"
		}
	}

	checklistStyle := func(items list.Items, index int) lipgloss.Style {
		switch index {
		case 1, 2, 4:
			return lipgloss.NewStyle().
				Strikethrough(true).
				Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"})
		default:
			return lipgloss.NewStyle()
		}
	}

	colors := colorGrid(1, 5)

	titleStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#FFF7DB"))

	lipglossStyleFunc := func(items list.Items, index int) lipgloss.Style {
		if index == items.Length()-1 {
			return titleStyle.Padding(1, 2).Margin(0, 0, 1, 0).MaxWidth(20).Background(lipgloss.Color(colors[index][0]))
		}
		return titleStyle.Padding(0, 5-index, 0, index+2).MaxWidth(20).Background(lipgloss.Color(colors[index][0]))
	}

	history := "Medieval quince preserves, which went by the French name cotignac, produced in a clear version and a fruit pulp version, began to lose their medieval seasoning of spices in the 16th century. In the 17th century, La Varenne provided recipes for both thick and clear cotignac."

	l := list.New().
		EnumeratorStyle(purple).
		Item("Lip Gloss").
		Item("Blush").
		Item("Eye Shadow").
		Item("Mascara").
		Item("Foundation").
		Item(
			list.New().
				EnumeratorStyle(pink).
				Item("Citrus Fruits to Try").
				Item(
					list.New().
						ItemStyleFunc(checklistStyle).
						EnumeratorStyleFunc(checklistEnumStyle).
						Enumerator(checklistEnum).
						Item("Grapefruit").
						Item("Yuzu").
						Item("Citron").
						Item("Kumquat").
						Item("Pomelo"),
				).
				Item("Actual Lip Gloss Vendors").
				Item(
					list.New().
						ItemStyleFunc(checklistStyle).
						EnumeratorStyleFunc(checklistEnumStyle).
						Enumerator(checklistEnum).
						Item("Glossier").
						Item("Claire‘s Boutique").
						Item("Nyx").
						Item("Mac").
						Item("Milk").
						Item(
							list.New().
								EnumeratorStyle(purple).
								Enumerator(list.Dash).
								ItemStyleFunc(lipglossStyleFunc).
								Item("Lip Gloss").
								Item("Lip Gloss").
								Item("Lip Gloss").
								Item("Lip Gloss").
								Item(
									list.New().
										EnumeratorStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(colors[4][0])).MarginRight(1)).
										Item("\nStyle Definitions for Nice Terminal Layouts\n─────").
										Item("From Charm").
										Item("https://github.com/charmbracelet/lipgloss").
										Item(
											list.New().
												EnumeratorStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(colors[3][0])).MarginRight(1)).
												Item("Emperors: Julio-Claudian dynasty").
												Item(
													lipgloss.NewStyle().Padding(1).Render(
														list.New(
															"Augustus",
															"Tiberius",
															"Caligula",
															"Claudius",
															"Nero",
														).Enumerator(list.Roman).String(),
													),
												).
												Item(
													lipgloss.NewStyle().
														Bold(true).
														Foreground(lipgloss.Color("#FAFAFA")).
														Background(lipgloss.Color("#7D56F4")).
														AlignHorizontal(lipgloss.Center).
														AlignVertical(lipgloss.Center).
														Padding(1, 3).
														Margin(0, 1, 1, 1).
														Width(40).
														Render(history),
												).
												Item(
													table.New().
														Width(30).
														BorderStyle(purple.MarginRight(0)).
														StyleFunc(func(row, col int) lipgloss.Style {
															style := lipgloss.NewStyle()

															if col == 0 {
																style = style.Align(lipgloss.Center)
															} else {
																style = style.Align(lipgloss.Right).PaddingRight(2)
															}
															if row == 0 {
																return style.Bold(true).Align(lipgloss.Center).PaddingRight(0)
															}
															return style.Faint(true)
														}).
														Headers("ITEM", "QUANTITY").
														Row("Apple", "6").
														Row("Banana", "10").
														Row("Orange", "2").
														Row("Strawberry", "12"),
												).
												Item("Documents").
												Item(
													list.New().
														Enumerator(func(_ list.Items, i int) string {
															if i == 1 {
																return "│\n│"
															}
															return " "
														}).
														ItemStyleFunc(func(_ list.Items, i int) lipgloss.Style {
															if i == 1 {
																return base.Foreground(highlight)
															}
															return base.Foreground(dim)
														}).
														EnumeratorStyleFunc(func(_ list.Items, i int) lipgloss.Style {
															if i == 1 {
																return lipgloss.NewStyle().Foreground(highlight)
															}
															return lipgloss.NewStyle().Foreground(dim)
														}).
														Item("Foo Document\n" + faint.Render("1 day ago")).
														Item("Bar Document\n" + faint.Render("2 days ago")).
														Item("Baz Document\n" + faint.Render("10 minutes ago")).
														Item("Qux Document\n" + faint.Render("1 month ago")),
												).
												Item("EOF"),
										).
										Item("go get github.com/charmbracelet/lipgloss/list\n"),
								).
								Item("See ya later"),
						),
				).
				Item("List"),
		).
		Item("xoxo, Charm_™")

	fmt.Println(l)
}

func colorGrid(xSteps, ySteps int) [][]string {
	x0y0, _ := colorful.Hex("#F25D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, ySteps)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(ySteps))
	}

	x1 := make([]colorful.Color, ySteps)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(ySteps))
	}

	grid := make([][]string, ySteps)
	for x := 0; x < ySteps; x++ {
		y0 := x0[x]
		grid[x] = make([]string, xSteps)
		for y := 0; y < xSteps; y++ {
			grid[x][y] = y0.BlendLuv(x1[x], float64(y)/float64(xSteps)).Hex()
		}
	}

	return grid
}

func showListExamples() {
	// Show List one example
	showListIExample()

	// Show List two example
	showListIIExample()

	// Show List three example
	showListIIIExample()

	// Show List four example
	showListIVExample()

	// Show List five example
	showListVExample()

	// Show List six example
	showListVIExample()
}

func showTableExamples() {

	// Show Table one example
	showTableIExample()

	// Show Table two example
	showTableIIExample()

	// Show Table three example
	showTableIIIExample()

	// Show Table four example
	showTableIVExample()

	// Show Table five example
	showTableVExample()
}

func showTreeExamples() {
	// Show Tree one example
	showTreeIExampleI()

	// Show Tree two example
	showTreeIExampleII()

	// Show Tree three example
	showTreeIExampleIII()

	// Show Tree four example
	showTreeIExampleIV()

	// Show Tree five example
	showTreeIExampleV()

	// Show Tree six example
	showTreeIExampleVI()

	// Show Tree seven example
	showTreeIExampleVII()

	// Show Tree eight example
	showTreeIExampleVIII()

	// Show Tree nine example
	showTreeIExampleIX()

	// Show Tree ten example
	showTreeIExampleX()
}

func showTableIExample() {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render

	t := table.New()
	t.Row("Bubble Tea", s("Milky"))
	t.Row("Milk Tea", s("Also milky"))
	t.Row("Actual milk", s("Milky as well"))
	fmt.Println(t.Render())
}

func showTableIIExample() {
	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Foreground(lipgloss.Color("241"))

	board := [][]string{
		{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"},
		{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
		{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"},
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		Rows(board...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		})

	ranks := labelStyle.Render(strings.Join([]string{" A", "B", "C", "D", "E", "F", "G", "H  "}, "   "))
	files := labelStyle.Render(strings.Join([]string{" 1", "2", "3", "4", "5", "6", "7", "8 "}, "\n\n "))

	fmt.Println(lipgloss.JoinVertical(lipgloss.Right, lipgloss.JoinHorizontal(lipgloss.Center, files, t.Render()), ranks) + "\n")
}

func showTableIIIExample() {
	re := lipgloss.NewRenderer(os.Stdout)
	// Colors
	const (
		purple    = lipgloss.Color("99")
		gray      = lipgloss.Color("245")
		lightGray = lipgloss.Color("241")
	)
	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = re.NewStyle().Padding(0, 1).Width(14)
		// OddRowStyle is the lipgloss style used for odd-numbered table rows.
		OddRowStyle = CellStyle.Foreground(gray)
		// EvenRowStyle is the lipgloss style used for even-numbered table rows.
		EvenRowStyle = CellStyle.Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle = lipgloss.NewStyle().Foreground(purple)
	)

	rows := [][]string{
		{"Chinese", "您好", "你好"},
		{"Japanese", "こんにちは", "やあ"},
		{"Arabic", "أهلين", "أهلا"},
		{"Russian", "Здравствуйте", "Привет"},
		{"Spanish", "Hola", "¿Qué tal?"},
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			// Make the second column a little wider.
			if col == 1 {
				style = style.Width(22)
			}

			// Arabic is a right-to-left language, so right align the text.
			if row < len(rows) && rows[row-1][0] == "Arabic" && col != 0 {
				style = style.Align(lipgloss.Right)
			}

			return style
		}).
		Headers("LANGUAGE", "FORMAL", "INFORMAL").
		Rows(rows...)

	t.Row("English", "You look absolutely fabulous.", "How's it going?")

	fmt.Println(t)
}

func showTableIVExample() {
	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Width(3).Align(lipgloss.Right)
	swatchStyle := re.NewStyle().Width(6)

	data := [][]string{}
	for i := 0; i < 13; i += 8 {
		data = append(data, makeRow(i, i+5))
	}
	data = append(data, makeEmptyRow())
	for i := 6; i < 15; i += 8 {
		data = append(data, makeRow(i, i+1))
	}
	data = append(data, makeEmptyRow())
	for i := 16; i < 231; i += 6 {
		data = append(data, makeRow(i, i+5))
	}
	data = append(data, makeEmptyRow())
	for i := 232; i < 256; i += 6 {
		data = append(data, makeRow(i, i+5))
	}

	t := table.New().
		Border(lipgloss.HiddenBorder()).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {
			color := lipgloss.Color(fmt.Sprint(data[row-1][col-col%2]))
			switch {
			case col%2 == 0:
				return labelStyle.Foreground(color)
			default:
				return swatchStyle.Background(color)
			}
		})

	fmt.Println(t)
}

func showTableVExample() {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	selectedStyle := baseStyle.Foreground(lipgloss.Color("#01BE85")).Background(lipgloss.Color("#00432F"))
	typeColors := map[string]lipgloss.Color{
		"Bug":      lipgloss.Color("#D7FF87"),
		"Electric": lipgloss.Color("#FDFF90"),
		"Fire":     lipgloss.Color("#FF7698"),
		"Flying":   lipgloss.Color("#FF87D7"),
		"Grass":    lipgloss.Color("#75FBAB"),
		"Ground":   lipgloss.Color("#FF875F"),
		"Normal":   lipgloss.Color("#929292"),
		"Poison":   lipgloss.Color("#7D5AFC"),
		"Water":    lipgloss.Color("#00E2C7"),
	}
	dimTypeColors := map[string]lipgloss.Color{
		"Bug":      lipgloss.Color("#97AD64"),
		"Electric": lipgloss.Color("#FCFF5F"),
		"Fire":     lipgloss.Color("#BA5F75"),
		"Flying":   lipgloss.Color("#C97AB2"),
		"Grass":    lipgloss.Color("#59B980"),
		"Ground":   lipgloss.Color("#C77252"),
		"Normal":   lipgloss.Color("#727272"),
		"Poison":   lipgloss.Color("#634BD0"),
		"Water":    lipgloss.Color("#439F8E"),
	}

	headers := []string{"#", "Name", "Type 1", "Type 2", "Japanese", "Official Rom."}
	data := [][]string{
		{"1", "Bulbasaur", "Grass", "Poison", "フシギダネ", "Bulbasaur"},
		{"2", "Ivysaur", "Grass", "Poison", "フシギソウ", "Ivysaur"},
		{"3", "Venusaur", "Grass", "Poison", "フシギバナ", "Venusaur"},
		{"4", "Charmander", "Fire", "", "ヒトカゲ", "Hitokage"},
		{"5", "Charmeleon", "Fire", "", "リザード", "Lizardo"},
		{"6", "Charizard", "Fire", "Flying", "リザードン", "Lizardon"},
		{"7", "Squirtle", "Water", "", "ゼニガメ", "Zenigame"},
		{"8", "Wartortle", "Water", "", "カメール", "Kameil"},
		{"9", "Blastoise", "Water", "", "カメックス", "Kamex"},
		{"10", "Caterpie", "Bug", "", "キャタピー", "Caterpie"},
		{"11", "Metapod", "Bug", "", "トランセル", "Trancell"},
		{"12", "Butterfree", "Bug", "Flying", "バタフリー", "Butterfree"},
		{"13", "Weedle", "Bug", "Poison", "ビードル", "Beedle"},
		{"14", "Kakuna", "Bug", "Poison", "コクーン", "Cocoon"},
		{"15", "Beedrill", "Bug", "Poison", "スピアー", "Spear"},
		{"16", "Pidgey", "Normal", "Flying", "ポッポ", "Poppo"},
		{"17", "Pidgeotto", "Normal", "Flying", "ピジョン", "Pigeon"},
		{"18", "Pidgeot", "Normal", "Flying", "ピジョット", "Pigeot"},
		{"19", "Rattata", "Normal", "", "コラッタ", "Koratta"},
		{"20", "Raticate", "Normal", "", "ラッタ", "Ratta"},
		{"21", "Spearow", "Normal", "Flying", "オニスズメ", "Onisuzume"},
		{"22", "Fearow", "Normal", "Flying", "オニドリル", "Onidrill"},
		{"23", "Ekans", "Poison", "", "アーボ", "Arbo"},
		{"24", "Arbok", "Poison", "", "アーボック", "Arbok"},
		{"25", "Pikachu", "Electric", "", "ピカチュウ", "Pikachu"},
		{"26", "Raichu", "Electric", "", "ライチュウ", "Raichu"},
		{"27", "Sandshrew", "Ground", "", "サンド", "Sand"},
		{"28", "Sandslash", "Ground", "", "サンドパン", "Sandpan"},
	}

	CapitalizeHeaders := func(data []string) []string {
		for i := range data {
			data[i] = strings.ToUpper(data[i])
		}
		return data
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(CapitalizeHeaders(headers)...).
		Width(80).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}

			if data[row-1][1] == "Pikachu" {
				return selectedStyle
			}

			even := row%2 == 0

			switch col {
			case 2, 3: // Type 1 + 2
				c := typeColors
				if even {
					c = dimTypeColors
				}

				color := c[fmt.Sprint(data[row-1][col])]
				return baseStyle.Foreground(color)
			}

			if even {
				return baseStyle.Foreground(lipgloss.Color("245"))
			}
			return baseStyle.Foreground(lipgloss.Color("252"))
		})
	fmt.Println(t)
}

func showTreeIExampleI() {
	enumeratorStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("0")).
		Padding(0, 1)

	headerItemStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#ee6ff8")).
		Foreground(lipgloss.Color("#ecfe65")).
		Bold(true).
		Padding(0, 1)

	itemStyle := headerItemStyle.Background(lipgloss.Color("0"))

	t := tree.Root("# Table of Contents").
		RootStyle(itemStyle).
		ItemStyle(itemStyle).
		EnumeratorStyle(enumeratorStyle).
		Child(
			tree.Root("## Chapter 1").
				Child("Chapter 1.1").
				Child("Chapter 1.2"),
		).
		Child(
			tree.Root("## Chapter 2").
				Child("Chapter 2.1").
				Child("Chapter 2.2"),
		)

	fmt.Println(t)
}

func showTreeIExampleII() {
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).PaddingRight(1)
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true).PaddingRight(1)

	t := tree.Root(".").EnumeratorStyle(enumeratorStyle).ItemStyle(itemStyle)
	_ = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			t.Child(tree.Root(path))
		}
		return nil
	})

	fmt.Println(t)
}

func showTreeIExampleIII() {
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("63")).MarginRight(1)
	rootStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("35"))
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	t := tree.
		Root("⁜ Makeup").
		Child(
			"Glossier",
			"Fenty Beauty",
			tree.New().Child(
				"Gloss Bomb Universal Lip Luminizer",
				"Hot Cheeks Velour Blushlighter",
			),
			"Nyx",
			"Mac",
			"Milk",
		).
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(enumeratorStyle).
		RootStyle(rootStyle).
		ItemStyle(itemStyle)

	fmt.Println(t)
}

func showTreeIExampleIV() {
	itemStyle := lipgloss.NewStyle().MarginRight(1)
	enumeratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).MarginRight(1)

	t := tree.Root("Groceries").
		Child(
			tree.Root("Fruits").
				Child(
					"Blood Orange",
					"Papaya",
					"Dragonfruit",
					"Yuzu",
				),
			tree.Root("Items").
				Child(
					"Cat Food",
					"Nutella",
					"Powdered Sugar",
				),
			tree.Root("Veggies").
				Child(
					"Leek",
					"Artichoke",
				),
		).ItemStyle(itemStyle).EnumeratorStyle(enumeratorStyle).Enumerator(tree.RoundedEnumerator)

	fmt.Println(t)
}

func showTreeIExampleV() {
	t := tree.Root(".").
		Child("macOS").
		Child(
			tree.New().
				Root("Linux").
				Child("NixOS").
				Child("Arch Linux (btw)").
				Child("Void Linux"),
		).
		Child(
			tree.New().
				Root("BSD").
				Child("FreeBSD").
				Child("OpenBSD"),
		)

	fmt.Println(t)
}

func showTreeIExampleVI() {
	purple := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
	pink := lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)

	t := tree.New().
		Child(
			"Glossier",
			"Claire’s Boutique",
			tree.Root("Nyx").
				Child("Lip Gloss", "Foundation").
				EnumeratorStyle(pink),
			"Mac",
			"Milk",
		).
		EnumeratorStyle(purple)
	fmt.Println(t)
}

func showTreeIExampleVII() {
	s := defaultStyles()

	t := tree.Root(dir{"~/charm", true, s}).
		Enumerator(tree.RoundedEnumerator).
		EnumeratorStyle(s.enumerator).
		Child(
			dir{"ayman", false, s},
			tree.Root(dir{"bash", true, s}).
				Child(
					tree.Root(dir{"tools", true, s}).
						Child(
							file{"zsh", s},
							file{"doom-emacs", s},
						),
				),
			tree.Root(dir{"carlos", true, s}).
				Child(
					tree.Root(dir{"emotes", true, s}).
						Child(
							file{"chefkiss.png", s},
							file{"kekw.png", s},
						),
				),
			dir{"maas", false, s},
		)

	fmt.Println(s.block.Render(t.String()))
}

func showTreeIExampleVIII() {

}

func showTreeIExampleIX() {

}

func showTreeIExampleX() {

}

func showTreeIExampleXI() {

}

func showTreeIExampleXII() {

}
