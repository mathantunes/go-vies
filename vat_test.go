package vies

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func Test_getCheckVatTemplate(t *testing.T) {
	type args struct {
		vat string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid VAT",
			args: args{vat: "FI25160553"},
			want: strings.TrimSpace(`<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
<Body>
	<checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
		<countryCode>FI</countryCode>
		<vatNumber>25160553</vatNumber>
	</checkVat>
</Body>
</Envelope>`),
		},
		{
			name:    "Invalid VAT",
			args:    args{vat: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCheckVatTemplate(tt.args.vat)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCheckVatTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCheckVatTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractVATResponse(t *testing.T) {
	type args struct {
		rc io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    *vatResponse
		wantErr bool
	}{
		{
			name:    "Invalid VAT Response",
			args:    args{rc: ioutil.NopCloser(bytes.NewReader([]byte("Invalid XML response")))},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVATResponse(tt.args.rc)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVATResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractVATResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
