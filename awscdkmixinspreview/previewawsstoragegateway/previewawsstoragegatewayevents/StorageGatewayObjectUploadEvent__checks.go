//go:build !no_runtime_type_checking

package previewawsstoragegatewayevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateStorageGatewayObjectUploadEvent_EventPatternParameters(options *StorageGatewayObjectUploadEvent_StorageGatewayObjectUploadEventProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

