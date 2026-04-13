//go:build !no_runtime_type_checking

package previewawstransferevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSFTPConnectorFileRetrieveFailed_EventPatternParameters(options *SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

