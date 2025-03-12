//go:build !no_runtime_type_checking

package awsecspatterns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateAddPortMappingForTargetsParameters(container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	if targets == nil {
		return fmt.Errorf("parameter targets is required, but nil was provided")
	}
	for idx_26cafb, v := range *targets {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter targets[%#v]", idx_26cafb) }); err != nil {
			return err
		}
	}

	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkMultipleTargetGroupsFargateService) validateRegisterECSTargetsParameters(service awsecs.BaseService, container awsecs.ContainerDefinition, targets *[]*NetworkTargetProps) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if container == nil {
		return fmt.Errorf("parameter container is required, but nil was provided")
	}

	if targets == nil {
		return fmt.Errorf("parameter targets is required, but nil was provided")
	}
	for idx_26cafb, v := range *targets {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter targets[%#v]", idx_26cafb) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNetworkMultipleTargetGroupsFargateService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkMultipleTargetGroupsFargateServiceParameters(scope constructs.Construct, id *string, props *NetworkMultipleTargetGroupsFargateServiceProps) error {
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

