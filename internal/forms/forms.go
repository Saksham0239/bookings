package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values        //form data is accessed here
	Errors     errors //keeps all the errors associated with all the fields
}

//creates a new form struct
func New(data url.Values) *Form {

	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//function to check for all required fields
func (f *Form) Required(fields ...string) {

	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}

}

//checks if a particular field is available in the form
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		return false
	} else {
		return true
	}
}

//Valid returns true if there are no errors otherwise false
func (f *Form) Valid() bool {
	fmt.Println("Valid working ", len(f.Errors) == 0)
	return len(f.Errors) == 0
}

//Validator for checking the length of the entered field in the form
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This string must be atleast %d characters long ", length))
		return false
	}
	return true
}

//IsEmail checks for valid email address using goValidator
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email Address")
	}
}
