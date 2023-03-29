// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties of the AppRunner VPC Connector.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &VpcConnectorProps{
//   	Vpc: Vpc,
//   	VpcSubnets: vpc.selectSubnets(&SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	}),
//   	VpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	VpcConnector: VpcConnector,
//   })
//
// Experimental.
type VpcConnectorProps struct {
	// The VPC for the VPC Connector.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The name for the VpcConnector.
	// Experimental.
	VpcConnectorName *string `field:"optional" json:"vpcConnectorName" yaml:"vpcConnectorName"`
	// Where to place the VPC Connector within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

