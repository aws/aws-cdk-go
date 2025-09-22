//go:build !no_runtime_type_checking

package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_ICfnResourceOptions) validateSetCreationPolicyParameters(val *CfnCreationPolicy) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_ICfnResourceOptions) validateSetUpdatePolicyParameters(val *CfnUpdatePolicy) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

