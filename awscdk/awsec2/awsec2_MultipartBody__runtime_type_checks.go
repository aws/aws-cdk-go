//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMultipartBody_FromRawBodyParameters(opts *MultipartBodyOptions) error {
	if opts == nil {
		return fmt.Errorf("parameter opts is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func validateMultipartBody_FromUserDataParameters(userData UserData) error {
	if userData == nil {
		return fmt.Errorf("parameter userData is required, but nil was provided")
	}

	return nil
}

