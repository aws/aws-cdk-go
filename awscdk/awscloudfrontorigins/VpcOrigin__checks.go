//go:build !no_runtime_type_checking

package awscloudfrontorigins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (v *jsiiProxy_VpcOrigin) validateBindParameters(scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateVpcOrigin_WithApplicationLoadBalancerParameters(alb awselasticloadbalancingv2.IApplicationLoadBalancer, props *VpcOriginWithEndpointProps) error {
	if alb == nil {
		return fmt.Errorf("parameter alb is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateVpcOrigin_WithEc2InstanceParameters(instance awsec2.IInstance, props *VpcOriginWithEndpointProps) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateVpcOrigin_WithNetworkLoadBalancerParameters(nlb awselasticloadbalancingv2.INetworkLoadBalancer, props *VpcOriginWithEndpointProps) error {
	if nlb == nil {
		return fmt.Errorf("parameter nlb is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateVpcOrigin_WithVpcOriginParameters(origin awscloudfront.IVpcOrigin, props *VpcOriginProps) error {
	if origin == nil {
		return fmt.Errorf("parameter origin is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewVpcOriginParameters(domainName *string, props *VpcOriginProps) error {
	if domainName == nil {
		return fmt.Errorf("parameter domainName is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

