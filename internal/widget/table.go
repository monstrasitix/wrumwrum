package widget

type Table struct {
	Headings []TableHeading
	Rows     [][]string
}

type TableHeading struct {
	Text  string
	Width int
	Type  string
}

type TableRow interface {
	ToHTMLTable() []string
}

func ToRows(rows []TableRow) [][]string {
	output := make([][]string, len(rows))

	for i, row := range rows {
		output[i] = row.ToHTMLTable()
	}

	return output
}
