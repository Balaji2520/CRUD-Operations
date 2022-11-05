package student

type StudentDetails struct {
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	DOB     string `json:"dob"`
	Year    uint   `json:"year"`
	Section string `json:"section"`
}
