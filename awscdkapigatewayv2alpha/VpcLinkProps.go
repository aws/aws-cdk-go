package awscdkapigatewayv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for a VpcLink.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &VpcLinkProps{
//   	Vpc: Vpc,
//   })
//
// Experimental.
type VpcLinkProps struct {
	// The VPC in which the private resources reside.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// A list of security groups for the VPC link.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of subnets for the VPC link.
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The name used to label and identify the VPC link.
	// Experimental.
	VpcLinkName *string `field:"optional" json:"vpcLinkName" yaml:"vpcLinkName"`
}

