package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Specifies the VPC that you want your Amazon SageMaker training job to connect to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
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
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		SubnetName: jsii.String("subnetName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   	},
//   }
//
// Experimental.
type VpcConfig struct {
	// VPC.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// VPC subnets.
	// Experimental.
	Subnets *awsec2.SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

