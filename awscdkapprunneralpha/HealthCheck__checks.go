//go:build !no_runtime_type_checking

package awscdkapprunneralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateHealthCheck_HttpParameters(options *HttpHealthCheckOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHealthCheck_TcpParameters(options *TcpHealthCheckOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

