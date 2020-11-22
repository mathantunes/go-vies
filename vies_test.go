package vies

import (
	"reflect"
	"testing"
)

func TestVIES_Validate(t *testing.T) {
	type args struct {
		vat string
	}
	tests := []struct {
		name    string
		v       Validator
		args    args
		want    *VATValidationResponse
		wantErr bool
	}{
		{
			name:    "Invalid VAT - Too short",
			v:       NewValidator(nil),
			args:    args{vat: "F"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid VAT - Wrong Number",
			v:    NewValidator(nil),
			args: args{vat: "FI123"},
			want: &VATValidationResponse{
				CountryCode: "FI",
				VATNumber:   "123",
				Name:        "---",
				Address:     "---",
				Valid:       false,
				Error:       &VATValidationResponse_VATValidationError{},
			},
			wantErr: false,
		},
		{
			name: "Valid VAT - From Finland",
			v:    NewValidator(nil),
			args: args{vat: "FI25160553"},
			want: &VATValidationResponse{
				CountryCode: "FI",
				VATNumber:   "25160553",
				Name:        "Comtower Finland Oy",
				Address:     "Sibeliuksenkatu 3\n08100 LOHJA",
				Valid:       true,
				Error:       &VATValidationResponse_VATValidationError{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Validate(tt.args.vat)
			if (err != nil) != tt.wantErr {
				t.Errorf("VIES.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VIES.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewValidator(t *testing.T) {
	type args struct {
		endpoint *string
	}
	var customEndpoint string = "CustomEndpoint"
	tests := []struct {
		name string
		args args
		want Validator
	}{
		{
			name: "Without custom endpoint",
			args: args{endpoint: nil},
			want: NewValidator(nil),
		},
		{
			name: "With custom endpoint",
			args: args{endpoint: &customEndpoint},
			want: NewValidator(&customEndpoint),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidator(tt.args.endpoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
