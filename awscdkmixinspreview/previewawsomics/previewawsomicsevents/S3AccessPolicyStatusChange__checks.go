//go:build !no_runtime_type_checking

package previewawsomicsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateS3AccessPolicyStatusChange_S3AccessPolicyStatusChangePatternParameters(options *S3AccessPolicyStatusChange_S3AccessPolicyStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

