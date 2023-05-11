//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IResolveContext) validateRegisterPostProcessorParameters(postProcessor IPostProcessor) error {
	if postProcessor == nil {
		return fmt.Errorf("parameter postProcessor is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IResolveContext) validateResolveParameters(x interface{}, options *ResolveChangeContextOptions) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

