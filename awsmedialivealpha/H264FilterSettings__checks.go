//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateH264FilterSettings_BandwidthReductionFilterParameters(props *BandwidthReductionFilterProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateH264FilterSettings_TemporalFilterParameters(props *TemporalFilterProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

