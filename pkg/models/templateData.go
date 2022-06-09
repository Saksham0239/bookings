package models

//Template Data holds data to be sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //any type
	csrfToken string
	Flash     string
	Warning   string
	Error     string
}
