package enums

type message string

var (
	Faileddecode  message = "Failed to decode the input message please see properly"
	Statusok      message = "Request was succesfully executed "
	Validation    message = "Validation failed"
	ServerIssue   message = "Problem in the server side"
	Deletesuccess message = "task deleted succesfully"
)
