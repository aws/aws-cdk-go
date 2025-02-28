package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for Transit Gateway VPC Attachment.
// Experimental.
type ITransitGatewayVpcAttachmentOptions interface {
	// Enable or disable appliance mode support.
	// Default: - disable (false).
	//
	// Experimental.
	ApplianceModeSupport() *bool
	// Enable or disable DNS support.
	// Default: - disable (false).
	//
	// Experimental.
	DnsSupport() *bool
	// Enable or disable IPv6 support.
	// Default: - disable (false).
	//
	// Experimental.
	Ipv6Support() *bool
	// Enables you to reference a security group across VPCs attached to a transit gateway.
	// Default: - disable (false).
	//
	// Experimental.
	SecurityGroupReferencingSupport() *bool
}

// The jsii proxy for ITransitGatewayVpcAttachmentOptions
type jsiiProxy_ITransitGatewayVpcAttachmentOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentOptions) ApplianceModeSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"applianceModeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentOptions) DnsSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentOptions) Ipv6Support() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ipv6Support",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayVpcAttachmentOptions) SecurityGroupReferencingSupport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"securityGroupReferencingSupport",
		&returns,
	)
	return returns
}

