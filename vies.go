package vies

const (
	// VIESEndpoint Current default vies endpoint
	VIESEndpoint = "http://ec.europa.eu/taxation_customs/vies/services/checkVatService"
)

// NewValidator Initializes the VIES API. An endpoint can be passed in or the default VIESEndpoint will be used
func NewValidator(endpoint *string) Validator {
	if endpoint == nil {
		return &VIES{
			endpoint: VIESEndpoint,
			soap:     newSoap(),
		}
	}
	return &VIES{
		endpoint: *endpoint,
		soap:     newSoap(),
	}
}

// Validator Exposed API for validating VAT
type Validator interface {
	Validate(vat string) (*VATValidationResponse, error)
}

// VIES Holds the API implementation for VIES
type VIES struct {
	endpoint string
	soap     soapRequester
}

// Validate Implementation of validation
func (v *VIES) Validate(vat string) (*VATValidationResponse, error) {
	reqPayload, err := getCheckVatTemplate(vat)
	if err != nil {
		return nil, err
	}
	bytesPayload := []byte(reqPayload)
	resp, err := v.soap.MakeRequest(v.endpoint, "checkVatService", bytesPayload)
	if err != nil {
		return nil, err
	}
	vatResponse, err := extractVATResponse(resp)
	if err != nil {
		return nil, err
	}
	return &VATValidationResponse{
		Address:     vatResponse.Body.CheckVat.Address,
		CountryCode: vatResponse.Body.CheckVat.CountryCode,
		Name:        vatResponse.Body.CheckVat.Name,
		// RequestDate: time.Parse(time.RFC3339, vatResponse.Body.CheckVat.RequestDate),
		VATNumber: vatResponse.Body.CheckVat.VatNumber,
		Valid:     vatResponse.Body.CheckVat.Valid,
		Error: &VATValidationResponse_VATValidationError{
			FaultCode:   vatResponse.Body.Fault.FaultCode,
			FaultString: vatResponse.Body.Fault.FaultString,
		},
	}, nil
}
