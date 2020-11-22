package vies

import (
	"fmt"
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
