//go:build !no_runtime_type_checking

package previewawsecsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateECSTaskStateChange_EcsTaskStateChangePatternParameters(options *ECSTaskStateChange_ECSTaskStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

