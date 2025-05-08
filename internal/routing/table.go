package routing

import (
	"fmt"
	"net/http"

	"sim/internal/web"
	"sim/internal/widget"
)

type TableHandler struct{}

type TablePayload struct {
	Table widget.Table
}

type Asset struct {
	LicensePlate string
	Type         string
	AxleLoad     float32
	Odometer     float32
	TPMS         string
	BrakeLining  string
}

func (a Asset) ToHTMLTable() []string {
	return []string{
		a.LicensePlate,
		a.Type,
		fmt.Sprintf("%f", a.AxleLoad),
		fmt.Sprintf("%f", a.Odometer),
		a.TPMS,
		a.BrakeLining,
	}
}

func (h TableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getIndex(w)
	}
}

func (h TableHandler) getIndex(w http.ResponseWriter) {
	assets := []Asset{
		{LicensePlate: "Something", Type: "Trailer", AxleLoad: 3.3123, Odometer: 234, TPMS: "tpms", BrakeLining: "Lining"},
		{LicensePlate: "Something", Type: "Trailer", AxleLoad: 3.3123, Odometer: 234, TPMS: "tpms", BrakeLining: "Lining"},
		{LicensePlate: "Something", Type: "Trailer", AxleLoad: 3.3123, Odometer: 234, TPMS: "tpms", BrakeLining: "Lining"},
		{LicensePlate: "Something", Type: "Trailer", AxleLoad: 3.3123, Odometer: 234, TPMS: "tpms", BrakeLining: "Lining"},
		{LicensePlate: "Something", Type: "Trailer", AxleLoad: 3.3123, Odometer: 234, TPMS: "tpms", BrakeLining: "Lining"},
		{LicensePlate: "Something", Type: "Trailer", AxleLoad: 3.3123, Odometer: 234, TPMS: "tpms", BrakeLining: "Lining"},
	}

	rows := []widget.TableRow{}

	for _, asset := range assets {
		rows = append(rows, asset)
	}

	web.HTML["table"].ExecuteTemplate(w, "base", TablePayload{
		Table: widget.Table{
			Headings: []widget.TableHeading{
				{Width: 200, Text: "License plate", Type: "string"},
				{Width: 100, Text: "Type", Type: "string"},
				{Width: 200, Text: "Axle load", Type: "numeric"},
				{Width: 200, Text: "Odometer", Type: "numeric"},
				{Width: 200, Text: "TPMS", Type: "string"},
				{Width: 200, Text: "Brake lining", Type: "string"},
			},
			Rows: widget.ToRows(rows),
		},
	})
}
