package cloudformation

type AWSServerlessFunction_Events = Value

/*

import (
	"encoding/json"

)

// AWSServerlessFunction_Events is a helper struct that can hold either a String or String value
type AWSServerlessFunction_Events struct {


	String **Value


	StringArray *[]*Value




}

func (r AWSServerlessFunction_Events) value() interface{} {


		if r.String != nil {
			return r.String
		}



		if r.StringArray != nil {
			return r.StringArray
		}








	return nil

}

func (r *AWSServerlessFunction_Events) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessFunction_Events) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {


			case *Value:
			r.String = &val



			case []*Value:
			r.StringArray = &val


		case map[string]interface{}:
		_ = val


		case []interface{}:
		_ = val

			json.Unmarshal(b, &r.StringArray)



	}

	return nil
}

*/
