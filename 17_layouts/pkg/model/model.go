package model

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int64
	FloatMap      map[string]float64
	Data          map[string]interface{}
	RemoteAddress string
}
