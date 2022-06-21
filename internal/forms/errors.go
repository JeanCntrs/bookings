package forms

type errors map[string][]string

// Add adds an error message for a give form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	er := e[field]

	if len(er) == 0 {
		return ""
	}

	return er[0]
}
