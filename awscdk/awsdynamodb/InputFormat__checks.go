//go:build !no_runtime_type_checking

package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateInputFormat_CsvParameters(options *CsvOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

