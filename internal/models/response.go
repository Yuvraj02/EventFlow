package models

//Basic Response struct
type  Response struct{
	Status string
	Message string
}

func (w *Response) GetField (fieldName string) string {
	
	switch fieldName {
	case "Status":
		return w.Status
	case "Message":
		return w.Message
	default:
		return "Field Name does not match"
	}
	
}

