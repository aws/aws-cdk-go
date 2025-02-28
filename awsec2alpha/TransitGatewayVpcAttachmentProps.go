package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Common properties for creating a Transit Gateway VPC Attachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var transitGateway transitGateway
//   var transitGatewayVpcAttachmentOptions iTransitGatewayVpcAttachmentOptions
//   var vpc vpc
//
//   transitGatewayVpcAttachmentProps := &TransitGatewayVpcAttachmentProps{
//   	Subnets: []iSubnet{
//   		subnet,
//   	},
//   	TransitGateway: transitGateway,
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	TransitGatewayAttachmentName: jsii.String("transitGatewayAttachmentName"),
//   	VpcAttachmentOptions: transitGatewayVpcAttachmentOptions,
//   }
//
// Experimental.
type TransitGatewayVpcAttachmentProps struct {
	// A list of one or more subnets to place the attachment in.
	//
	// It is recommended to specify more subnets for better availability.
	// Experimental.
	Subnets *[]awsec2.ISubnet `field:"required" json:"subnets" yaml:"subnets"`
	// The transit gateway this attachment gets assigned to.
	// Experimental.
	TransitGateway ITransitGateway `field:"required" json:"transitGateway" yaml:"transitGateway"`
	// A VPC attachment(s) will get assigned to.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Physical name of this Transit Gateway VPC Attachment.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayAttachmentName *string `field:"optional" json:"transitGatewayAttachmentName" yaml:"transitGatewayAttachmentName"`
	// The VPC attachment options.
	// Default: - All options are disabled.
	//
	// Experimental.
	VpcAttachmentOptions ITransitGatewayVpcAttachmentOptions `field:"optional" json:"vpcAttachmentOptions" yaml:"vpcAttachmentOptions"`
}

