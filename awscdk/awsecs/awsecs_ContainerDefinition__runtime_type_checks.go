//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ContainerDefinition) validateAddContainerDependenciesParameters(containerDependencies *[]*ContainerDependency) error {
	for idx_808b21, v := range *containerDependencies {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter containerDependencies[%#v]", idx_808b21) }); err != nil {
			return err
		}
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddEnvironmentParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddLinkParameters(container ContainerDefinition) error {
	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddMountPointsParameters(mountPoints *[]*MountPoint) error {
	for idx_ccec52, v := range *mountPoints {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter mountPoints[%#v]", idx_ccec52) }); err != nil {
			return err
		}
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddPortMappingsParameters(portMappings *[]*PortMapping) error {
	for idx_455c7f, v := range *portMappings {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter portMappings[%#v]", idx_455c7f) }); err != nil {
			return err
		}
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddScratchParameters(scratch *ScratchSpace) error {
	if scratch == nil {
		return fmt.Errorf("parameter scratch is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(scratch, func() string { return "parameter scratch" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddSecretParameters(name *string, secret Secret) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddToExecutionPolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddUlimitsParameters(ulimits *[]*Ulimit) error {
	for idx_444e41, v := range *ulimits {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter ulimits[%#v]", idx_444e41) }); err != nil {
			return err
		}
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddVolumesFromParameters(volumesFrom *[]*VolumeFrom) error {
	for idx_57d77a, v := range *volumesFrom {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter volumesFrom[%#v]", idx_57d77a) }); err != nil {
			return err
		}
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateFindPortMappingParameters(containerPort *float64, protocol Protocol) error {
	if containerPort == nil {
		return fmt.Errorf("parameter containerPort is required, but nil was provided")
	}

	if protocol == "" {
		return fmt.Errorf("parameter protocol is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateFindPortMappingByNameParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateContainerDefinition_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewContainerDefinitionParameters(scope constructs.Construct, id *string, props *ContainerDefinitionProps) error {
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

