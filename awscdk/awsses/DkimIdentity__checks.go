//go:build !no_runtime_type_checking

package awsses

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_DkimIdentity) validateBindParameters(emailIdentity EmailIdentity) error {
	if emailIdentity == nil {
		return fmt.Errorf("parameter emailIdentity is required, but nil was provided")
	}

	return nil
}

func validateDkimIdentity_ByoDkimParameters(options *ByoDkimOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

