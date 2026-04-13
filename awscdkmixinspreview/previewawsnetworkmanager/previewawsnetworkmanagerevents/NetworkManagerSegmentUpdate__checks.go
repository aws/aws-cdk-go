//go:build !no_runtime_type_checking

package previewawsnetworkmanagerevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNetworkManagerSegmentUpdate_EventPatternParameters(options *NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

