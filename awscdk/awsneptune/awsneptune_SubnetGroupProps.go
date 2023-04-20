package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties for creating a SubnetGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   subnetGroupProps := &SubnetGroupProps{
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RemovalPolicy: monocdk.RemovalPolicy_DESTROY,
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
//   		SubnetName: jsii.String("subnetName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   	},
//   }
//
// Experimental.
type SubnetGroupProps struct {
	// The VPC to place the subnet group in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Description of the subnet group.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The removal policy to apply when the subnet group are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The name of the subnet group.
	// Experimental.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Which subnets within the VPC to associate with this group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

