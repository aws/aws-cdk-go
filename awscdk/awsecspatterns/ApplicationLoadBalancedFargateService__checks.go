//go:build !no_runtime_type_checking

package awsecspatterns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_ApplicationLoadBalancedFargateService) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedFargateService) validateCreateAWSLogDriverParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateApplicationLoadBalancedFargateService_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationLoadBalancedFargateServiceParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedFargateServiceProps) error {
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

