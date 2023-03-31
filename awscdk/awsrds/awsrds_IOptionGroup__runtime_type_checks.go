//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IOptionGroup) validateAddConfigurationParameters(configuration *OptionConfiguration) error {
	if configuration == nil {
		return fmt.Errorf("parameter configuration is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(configuration, func() string { return "parameter configuration" }); err != nil {
		return err
	}

	return nil
}

