package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for creating a SubnetGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   subnetGroupProps := &SubnetGroupProps{
//   	Description: jsii.String("description"),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	VpcSubnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type SubnetGroupProps struct {
	// Description of the subnet group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The VPC to place the subnet group in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The removal policy to apply when the subnet group are removed from the stack or replaced during an update.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The name of the subnet group.
	// Default: - a name is generated.
	//
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Which subnets within the VPC to associate with this group.
	// Default: - private subnets.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

