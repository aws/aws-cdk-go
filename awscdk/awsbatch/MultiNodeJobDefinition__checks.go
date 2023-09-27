//go:build !no_runtime_type_checking

package awsbatch

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_MultiNodeJobDefinition) validateAddContainerParameters(container *MultiNodeContainer) error {
	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(container, func() string { return "parameter container" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MultiNodeJobDefinition) validateAddRetryStrategyParameters(strategy RetryStrategy) error {
	if strategy == nil {
		return fmt.Errorf("parameter strategy is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MultiNodeJobDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MultiNodeJobDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MultiNodeJobDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func validateMultiNodeJobDefinition_FromJobDefinitionArnParameters(scope constructs.Construct, id *string, jobDefinitionArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if jobDefinitionArn == nil {
		return fmt.Errorf("parameter jobDefinitionArn is required, but nil was provided")
	}

	return nil
}

func validateMultiNodeJobDefinition_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateMultiNodeJobDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateMultiNodeJobDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewMultiNodeJobDefinitionParameters(scope constructs.Construct, id *string, props *MultiNodeJobDefinitionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

