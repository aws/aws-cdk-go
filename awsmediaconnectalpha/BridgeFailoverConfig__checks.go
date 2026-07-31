//go:build !no_runtime_type_checking

package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateBridgeFailoverConfig_FailoverParameters(options *BridgeFailoverOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

