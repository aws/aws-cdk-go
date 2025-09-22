//go:build !no_runtime_type_checking

package awscdkmskalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateClientAuthentication_SaslParameters(props *SaslAuthProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateClientAuthentication_SaslTlsParameters(saslTlsProps *SaslTlsAuthProps) error {
	if saslTlsProps == nil {
		return fmt.Errorf("parameter saslTlsProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(saslTlsProps, func() string { return "parameter saslTlsProps" }); err != nil {
		return err
	}

	return nil
}

func validateClientAuthentication_TlsParameters(props *TlsAuthProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

