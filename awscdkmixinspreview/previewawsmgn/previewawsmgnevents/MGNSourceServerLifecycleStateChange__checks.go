//go:build !no_runtime_type_checking

package previewawsmgnevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMGNSourceServerLifecycleStateChange_EventPatternParameters(options *MGNSourceServerLifecycleStateChange_MGNSourceServerLifecycleStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

