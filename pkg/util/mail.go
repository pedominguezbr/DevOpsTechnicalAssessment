package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// MailConstants handle control
type EmailConstants struct {
	id  string
	url string
}

// MailConstants create for email
func NewEmailConstants(id string, url string) *EmailConstants {
	return &EmailConstants{id, url}
}

func (m *EmailConstants) SendEmail(mensaje string) error {
	fmt.Printf("Enviar correo \n")
	fmt.Printf(`mensaje SendMail: "%s"\n`, mensaje)
	fmt.Println()

	//url := "https://ws-envio-correo-qa.spsa-api.lblapiqa.spsa.xyz/send-mail"
	fmt.Println("URL:>", m.url)

	var jsonMail bytes.Buffer
	//jsonMail.WriteString(`{"applicationCode": "I0469","company": "SP",`)
	var strApp = `{"applicationCode": "` + m.id + `","company": "SP",`
	jsonMail.WriteString(strApp)
	jsonMail.WriteString(`"email": {"subject": "Error API WMS Inventory History",`)
	jsonMail.WriteString(`"keys": [{"name": "content","detail": "`)
	jsonMail.WriteString(mensaje)
	jsonMail.WriteString(`"}]}}`)

	fmt.Println("jsonMail:>", jsonMail.String())
	var jsonStr = []byte(jsonMail.String())
	req, err := http.NewRequest("POST", m.url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		log.Println("Error SendEmail: " + err.Error())
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	// fmt.Printf("resp: %v\n", resp)
	fmt.Printf("Correo enviado\n")
	return nil
}
