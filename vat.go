package vies

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

const checkVatTemplate = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
<Body>
	<checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
		<countryCode>%s</countryCode>
		<vatNumber>%s</vatNumber>
	</checkVat>
</Body>
</Envelope>`

func getCheckVatTemplate(vat string) string {
	return fmt.Sprintf(checkVatTemplate, strings.ToUpper(vat[0:2]), vat[2:])
}

// ValidationVAT Response message for valid responses from VIES
type ValidationVAT struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	RequestDate string `xml:"requestDate"`
	Valid       bool   `xml:"valid"`
	Name        string `xml:"name"`
	Address     string `xml:"address"`
}

// ValidationFault Response message for invalid responses from VIES
type ValidationFault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}

// body XML Interpretation Body for SOAP
type body struct {
	XMLName  xml.Name
	CheckVat ValidationVAT   `xml:"checkVatResponse"`
	Fault    ValidationFault `xml:"Fault"`
}

// vatResponse XML Interpretation of SOAP CheckVAT Response
type vatResponse struct {
	XMLName xml.Name
	Body    body
}

func extractVATResponse(rc io.ReadCloser) (*vatResponse, error) {
	vatResp := new(vatResponse)
	err := xml.NewDecoder(rc).Decode(vatResp)
	return vatResp, err
}
