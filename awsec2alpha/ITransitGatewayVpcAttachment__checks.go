//go:build !no_runtime_type_checking

package awsec2alpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func (i *jsiiProxy_ITransitGatewayVpcAttachment) validateAddSubnetsParameters(subnets *[]awsec2.ISubnet) error {
	if subnets == nil {
		return fmt.Errorf("parameter subnets is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITransitGatewayVpcAttachment) validateRemoveSubnetsParameters(subnets *[]awsec2.ISubnet) error {
	if subnets == nil {
		return fmt.Errorf("parameter subnets is required, but nil was provided")
	}

	return nil
}

