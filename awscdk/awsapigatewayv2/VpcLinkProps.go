package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Properties for a VpcLink.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   alb := elb.NewApplicationLoadBalancer(this, jsii.String("AppLoadBalancer"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   })
//
//   vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &VpcLinkProps{
//   	Vpc: Vpc,
//   })
//
//   // Creating an HTTP ALB Integration:
//   albIntegration := awscdk.NewHttpAlbIntegration(jsii.String("ALBIntegration"), alb.Listeners[jsii.Number(0)], &HttpAlbIntegrationProps{
//   })
//
type VpcLinkProps struct {
	// The VPC in which the private resources reside.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// A list of security groups for the VPC link.
	// Default: - no security groups. Use `addSecurityGroups` to add security groups
	//
	SecurityGroups *[]interfacesawsec2.ISecurityGroupRef `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of subnets for the VPC link.
	// Default: - private subnets of the provided VPC. Use `addSubnets` to add more subnets
	//
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The name used to label and identify the VPC link.
	// Default: - automatically generated name.
	//
	VpcLinkName *string `field:"optional" json:"vpcLinkName" yaml:"vpcLinkName"`
}

