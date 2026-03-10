//go:build !no_runtime_type_checking

package previewawsrdsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateRDSDBProxyEvent_RdsDBProxyEventPatternParameters(options *RDSDBProxyEvent_RDSDBProxyEventProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

