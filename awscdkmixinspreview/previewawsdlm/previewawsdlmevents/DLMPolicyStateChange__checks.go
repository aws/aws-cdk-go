//go:build !no_runtime_type_checking

package previewawsdlmevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDLMPolicyStateChange_DlmPolicyStateChangePatternParameters(options *DLMPolicyStateChange_DLMPolicyStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

