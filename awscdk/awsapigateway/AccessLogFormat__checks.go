//go:build !no_runtime_type_checking

package awsapigateway

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAccessLogFormat_CustomParameters(format *string) error {
	if format == nil {
		return fmt.Errorf("parameter format is required, but nil was provided")
	}

	return nil
}

func validateAccessLogFormat_JsonWithStandardFieldsParameters(fields *JsonWithStandardFieldProps) error {
	if err := _jsii_.ValidateStruct(fields, func() string { return "parameter fields" }); err != nil {
		return err
	}

	return nil
}

