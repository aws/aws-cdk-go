//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_MissingRemovalPolicies) validateApplyParameters(policy RemovalPolicy, props *RemovalPolicyProps) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateDestroyParameters(props *RemovalPolicyProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateRetainParameters(props *RemovalPolicyProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateRetainOnUpdateOrDeleteParameters(props *RemovalPolicyProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateSnapshotParameters(props *RemovalPolicyProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateMissingRemovalPolicies_OfParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

