package models

type Panel int

const (
	TabsPanelBoards Panel = iota
	TabsPanelMonths
	TabsPanelYears
	LeftSidebarPanel
	CentralPanel
	RightSidebarPanel
)

func (p Panel) String() string {
	switch p {
	case TabsPanelMonths:
		return "Tabs Months"
	case TabsPanelYears:
		return "Tabs Years"
	case TabsPanelBoards:
		return "Tabs Boards"
	case LeftSidebarPanel:
		return "Left Sidebar"
	case CentralPanel:
		return "Central"
	case RightSidebarPanel:
		return "Right Sidebar"
	default:
		return "?"
	}
}

func (p Panel) Int() int {
	return int(p)
}

func (p *Panel) Next() {
	*p = (*p + 1) % 4

	if *p == 0 {
		*p = 4
	}
}

func (p *Panel) Prev() {
	*p = (*p - 1) % 4
}

func (p *Panel) Reset() {
	*p = TabsPanelMonths
}
