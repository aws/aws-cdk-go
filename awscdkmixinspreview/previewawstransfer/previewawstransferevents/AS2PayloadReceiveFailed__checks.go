//go:build !no_runtime_type_checking

package previewawstransferevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAS2PayloadReceiveFailed_As2PayloadReceiveFailedPatternParameters(options *AS2PayloadReceiveFailed_AS2PayloadReceiveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

