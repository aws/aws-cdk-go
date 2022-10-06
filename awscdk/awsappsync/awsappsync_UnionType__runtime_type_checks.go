//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (u *jsiiProxy_UnionType) validateAddFieldParameters(options *AddFieldOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (u *jsiiProxy_UnionType) validateAttributeParameters(options *BaseTypeOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewUnionTypeParameters(name *string, options *UnionTypeOptions) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

