package metadata

type Grid struct {
	Columns []*Column `json:"column"`
}

type Column struct {
	Width           int           `json:"width"`
	Field           string        `json:"field"`
	AutoHeight      bool          `json:"autoHeight"`
	Editor          string        `json:"editor"`
	EditorOptions   any           `json:"editorOptions"`
	Template        string        `json:"template"`
	Align           string        `json:"align"`
	CellsFormat     string        `json:"cellsFormat"`
	CellsAlign      string        `json:"cellsAlign"`
	ColumnGroup     string        `json:"columnGroup"`
	Filter          string        `json:"filter"`
	AggType         AggType       `json:"aggType"`
	Freeze          Freeze        `json:"freeze"`
	Visible         bool          `json:"visible"`
	Use             bool          `json:"use"`
	Editable        bool          `json:"editable"`
	Sort            SortDirection `json:"sort"`
	AllowExport     bool          `json:"allowExport"`
	AllowGroup      bool          `json:"allowGroup"`
	AllowHide       bool          `json:"allowHide"`
	AllowSelect     bool          `json:"allowSelect"`
	AllowEdit       bool          `json:"allowEdit"`
	AllowSort       bool          `json:"allowSort"`
	AllowHeaderEdit bool          `json:"allowHeaderEdit"`
	AllowFilter     bool          `json:"allowFilter"`
	AllowReorder    bool          `json:"allowReorder"`
	AllowResize     bool          `json:"allowResize"`
	AllowNull       bool          `json:"allowNull"`
}

type Freeze string

const (
	FreezeLeft  Freeze = "left"
	FreezeRight Freeze = "right"
)

type AggType string

const (
	AggTypeSum AggType = "sum"
	AggTypeAvg AggType = "avg"
	AggTypeMax AggType = "max"
	AggTypeMin AggType = "min"
)

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)
