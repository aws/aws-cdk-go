//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BaseLoadBalancer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (b *jsiiProxy_BaseLoadBalancer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateLogAccessLogsParameters(bucket awss3.IBucket) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateRemoveAttributeParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BaseLoadBalancer) validateSetAttributeParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateBaseLoadBalancer_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateBaseLoadBalancer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBaseLoadBalancer_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewBaseLoadBalancerParameters(scope constructs.Construct, id *string, baseProps *BaseLoadBalancerProps, additionalProps interface{}) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if baseProps == nil {
		return fmt.Errorf("parameter baseProps is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(baseProps, func() string { return "parameter baseProps" }); err != nil {
		return err
	}

	if additionalProps == nil {
		return fmt.Errorf("parameter additionalProps is required, but nil was provided")
	}

	return nil
}

