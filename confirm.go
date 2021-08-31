package confirmbox

//New creates a new confirm box and wait for the user input
//Returns true if user accepts
func New(title, content string) bool {
	return isConfirmed(title, content)
}
