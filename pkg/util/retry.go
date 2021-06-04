package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RetryConstants struct {
	url string
}

// RetryConstants create for retry
func NewRetryConstants(url string) *RetryConstants {
	return &RetryConstants{url}
}

func (r *RetryConstants) SendHandlerError(data string, integracion string, sessionID string, email *EmailConstants) error {
	fmt.Printf("Enviar handler error \n")
	fmt.Println()

	//url := "http://localhost:8080/int-api-wms-handler-error/api/v1/error/upload-file"
	fmt.Println("URL:>", r.url)

	/*var jsonMail bytes.Buffer
	jsonMail.WriteString(`{"applicationCode": "I0469","company": "SP",`)
	jsonMail.WriteString(`"email": {"subject": "Error API WMS Inventory History",`)
	jsonMail.WriteString(`"keys": [{"name": "content","detail": "`)
	jsonMail.WriteString(mensaje)
	jsonMail.WriteString(`"}]}}`)*/

	/*var data = `[H1]P445|SPSA|CREATE|LTL/TL|OSP44500032858|RT-ECOMMERCE ARN-745	|||||P261||8330011817|2|198|1940598|0|20200618000000||
	[H2]1|1|||2|198|1940598|0|P261||AV JOSE CARLOS MARIATEGUI S/N|72169266||LIMA|LIMA|||+51963399799|||P261||AV JOSE CARLOS MARIATEGUI S/N|72169266||LIMA|LIMA|||+51963399799||261 - SPSA PVEA EL AGUSTINO|RT P261|SDR00000241036|20200616000000|20200617000000|20200617000000|||||||||||||2|P261|||||CTP4450000096061|00020173098|2200201730982|||||||0|0|0||||False|UNITS|1|||||||||||99|970299|0|ECO
	[H2]2|1|||2|198|1940598|0|P261||AV JOSE CARLOS MARIATEGUI S/N|72169266||LIMA|LIMA|||+51963399799|||P261||AV JOSE CARLOS MARIATEGUI S/N|72169266||LIMA|LIMA|||+51963399799||261 - SPSA PVEA EL AGUSTINO|RT P261|SDR00000241036|20200616000000|20200617000000|20200617000000|||||||||||||1|P261|||||CTP4450000096060|00020173129|2200201731293|||||||0|0|0||||False|UNITS|1|||||||||||99|970299|0|ECO
	`*/

	fmt.Println("jsonMail:>", data)
	var jsonStr = []byte(data)
	req, err := http.NewRequest("POST", r.url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	//q.Add("integracion", "abc")
	//q.Add("sessionID", "666")
	q.Add("integracion", integracion)
	q.Add("sessionID", sessionID)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		log.Println("Error SendHandlerError: " + err.Error())
		email.SendEmail(CorrectFormat(err.Error()))
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
