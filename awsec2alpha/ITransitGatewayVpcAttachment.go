package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Represents a Transit Gateway VPC Attachment.
// Experimental.
type ITransitGatewayVpcAttachment interface {
	ITransitGatewayAttachment
	// Add additional subnets to this attachment.
	// Experimental.
	AddSubnets(subnets *[]awsec2.ISubnet)
	// Remove subnets from this attachment.
	// Experimental.
	RemoveSubnets(subnets *[]awsec2.ISubnet)
}

// The jsii proxy for ITransitGatewayVpcAttachment
type jsiiProxy_ITransitGatewayVpcAttachment struct {
	jsiiProxy_ITransitGatewayAttachment
}

func (i *jsiiProxy_ITransitGatewayVpcAttachment) AddSubnets(subnets *[]awsec2.ISubnet) {
	if err := i.validateAddSubnetsParameters(subnets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addSubnets",
		[]interface{}{subnets},
	)
}

func (i *jsiiProxy_ITransitGatewayVpcAttachment) RemoveSubnets(subnets *[]awsec2.ISubnet) {
	if err := i.validateRemoveSubnetsParameters(subnets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"removeSubnets",
		[]interface{}{subnets},
	)
}

