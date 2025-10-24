package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Specifies the VPC that you want your Amazon SageMaker training job to connect to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var vpc Vpc
//
//   vpcConfig := &VpcConfig{
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	Subnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []SubnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []ISubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type VpcConfig struct {
	// VPC.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// VPC subnets.
	// Default: - Private Subnets are selected.
	//
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

