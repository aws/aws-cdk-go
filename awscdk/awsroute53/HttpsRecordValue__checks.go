//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateHttpsRecordValue_AliasParameters(targetName *string) error {
	if targetName == nil {
		return fmt.Errorf("parameter targetName is required, but nil was provided")
	}

	return nil
}

func validateHttpsRecordValue_ServiceParameters(props *HttpsRecordServiceModeProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

