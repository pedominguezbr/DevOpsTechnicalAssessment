package util

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gofrs/uuid"
)

func ValidateInputs(obj []interface{}) {
	for i := 0; i < len(obj); i++ {
		for _, v := range obj[i].([]interface{}) { // use type assertion to loop over []interface{}
			fmt.Printf("%v ", v)
		}
	}

}

func ResponseService(Isvalid bool, typeMsg int, code string, msg string) string {
	msgData := Message{Type: typeMsg, Code: code, Message: msg}
	msgDataArray := []Message{msgData}
	//MessageData = append(MessageData, msgData)
	response := Response{Messages: msgDataArray, IsValid: Isvalid}

	//fmt.Println("ResponseService.response ", response)
	var jsonData []byte
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return ""
	}
	//fmt.Println("ResponseService.Marshal ", string(jsonData))
	return fmt.Sprintf("%v", string(jsonData))
}

func ResponseMessageService(msg string) string {
	msgData := MessageResult{Message: msg}

	var jsonData []byte
	jsonData, err := json.Marshal(msgData)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("%v", string(jsonData))
}

func StringResponseToResponseObj(stringObj string) interface{} {
	//log.Println("stringObj: ", stringObj)
	byt := []byte(stringObj)
	var responseData map[string]interface{}

	if err := json.Unmarshal(byt, &responseData); err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		panic(err)
	}
	return responseData
}

func ConvertErrorDvToErrorDbStrcut(err error) ErrorDBStruct {
	res2B, _ := json.Marshal(err)
	var eDv ErrorDBStruct
	json.Unmarshal(res2B, &eDv)
	return eDv
}

func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

func CorrectFormat(mensaje string) string {
	formatText := strings.Replace(mensaje, "\"", "", -1)
	formatText = strings.Replace(formatText, "\n", " ", -1)
	return strings.TrimSpace(formatText)
}

// NewUUID creates a new unique universal identifier
func NewUUID() (string, error) {

	result, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return result.String(), nil
}
