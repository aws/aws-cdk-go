//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"
)

func validateAppProtocol_SetGrpcParameters(val AppProtocol) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateAppProtocol_SetHttpParameters(val AppProtocol) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateAppProtocol_SetHttp2Parameters(val AppProtocol) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAppProtocolParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

