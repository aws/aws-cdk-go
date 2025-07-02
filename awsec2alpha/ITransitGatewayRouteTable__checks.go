//go:build !no_runtime_type_checking

package awsec2alpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (i *jsiiProxy_ITransitGatewayRouteTable) validateAddAssociationParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if transitGatewayAttachment == nil {
		return fmt.Errorf("parameter transitGatewayAttachment is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateAddBlackholeRouteParameters(id *string, destinationCidr *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if destinationCidr == nil {
		return fmt.Errorf("parameter destinationCidr is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateAddRouteParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment, destinationCidr *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if transitGatewayAttachment == nil {
		return fmt.Errorf("parameter transitGatewayAttachment is required, but nil was provided")
	}

	if destinationCidr == nil {
		return fmt.Errorf("parameter destinationCidr is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateEnablePropagationParameters(id *string, transitGatewayAttachment ITransitGatewayAttachment) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if transitGatewayAttachment == nil {
		return fmt.Errorf("parameter transitGatewayAttachment is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_ITransitGatewayRouteTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

