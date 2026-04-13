//go:build !no_runtime_type_checking

package previewawss3events

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateObjectTagsDeleted_EventPatternParameters(options *ObjectTagsDeleted_ObjectTagsDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

