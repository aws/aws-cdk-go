//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TaskDefinition) validateAddContainerParameters(id *string, props *ContainerDefinitionOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddExtensionParameters(extension ITaskDefinitionExtension) error {
	if extension == nil {
		return fmt.Errorf("parameter extension is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddFirelensLogRouterParameters(id *string, props *FirelensLogRouterDefinitionOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddInferenceAcceleratorParameters(inferenceAccelerator *InferenceAccelerator) error {
	if inferenceAccelerator == nil {
		return fmt.Errorf("parameter inferenceAccelerator is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(inferenceAccelerator, func() string { return "parameter inferenceAccelerator" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddPlacementConstraintParameters(constraint PlacementConstraint) error {
	if constraint == nil {
		return fmt.Errorf("parameter constraint is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddToExecutionRolePolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddToTaskRolePolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateAddVolumeParameters(volume *Volume) error {
	if volume == nil {
		return fmt.Errorf("parameter volume is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(volume, func() string { return "parameter volume" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateFindContainerParameters(containerName *string) error {
	if containerName == nil {
		return fmt.Errorf("parameter containerName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateFindPortMappingByNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (t *jsiiProxy_TaskDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TaskDefinition) validateGrantRunParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateTaskDefinition_FromTaskDefinitionArnParameters(scope constructs.Construct, id *string, taskDefinitionArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if taskDefinitionArn == nil {
		return fmt.Errorf("parameter taskDefinitionArn is required, but nil was provided")
	}

	return nil
}

func validateTaskDefinition_FromTaskDefinitionAttributesParameters(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateTaskDefinition_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTaskDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateTaskDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewTaskDefinitionParameters(scope constructs.Construct, id *string, props *TaskDefinitionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

