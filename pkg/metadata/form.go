package metadata

type Form struct {
	Editors []*Editor `json:"widget"`
}

type Editor struct {
	Widget Widget `json:"widget"`
	Help   string `json:"help"`
}

type Widget string

const (
	None              Widget = ""
	AltDateTimeWidget Widget = "altDateTime"
	AltDateWidget     Widget = "altDate"
	CheckboxesWidget  Widget = "checkboxes"
	CheckboxWidget    Widget = "checkbox"
	ColorWidget       Widget = "color"
	DateTimeWidget    Widget = "dateTime"
	DateWidget        Widget = "date"
	EmailWidget       Widget = "email"
	FileWidget        Widget = "file"
	HiddenWidget      Widget = "hidden"
	PasswordWidget    Widget = "password"
	RadioWidget       Widget = "radio"
	RangeWidget       Widget = "range"
	SelectWidget      Widget = "select"
	TextareaWidget    Widget = "textarea"
	TextWidget        Widget = "text"
	TimeWidget        Widget = "time"
	UpDownWidget      Widget = "upDown"
	URLWidget         Widget = "URL"
)
