//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

func (l *jsiiProxy_LoadBalancerTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	if _record == nil {
		return fmt.Errorf("parameter _record is required, but nil was provided")
	}

	return nil
}

func validateNewLoadBalancerTargetParameters(loadBalancer awselasticloadbalancingv2.ILoadBalancerV2) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	return nil
}

