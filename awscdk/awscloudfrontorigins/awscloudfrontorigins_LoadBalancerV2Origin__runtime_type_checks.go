//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscloudfrontorigins

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LoadBalancerV2Origin) validateBindParameters(_scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewLoadBalancerV2OriginParameters(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2, props *LoadBalancerV2OriginProps) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

