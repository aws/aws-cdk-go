// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties of the AppRunner VPC Connector.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   	ipAddresses: ec2.ipAddresses.cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &vpcConnectorProps{
//   	vpc: vpc,
//   	vpcSubnets: vpc.selectSubnets(&subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	}),
//   	vpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &serviceProps{
//   	source: apprunner.source.fromEcrPublic(&ecrPublicProps{
//   		imageConfiguration: &imageConfiguration{
//   			port: jsii.Number(8000),
//   		},
//   		imageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	vpcConnector: vpcConnector,
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

