//go:build !no_runtime_type_checking

package previewawsomicsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateVariantImportJobStatusChange_EventPatternParameters(options *VariantImportJobStatusChange_VariantImportJobStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

