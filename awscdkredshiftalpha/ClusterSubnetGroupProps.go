package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for creating a ClusterSubnetGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import redshift_alpha "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var vpc Vpc
//
//   clusterSubnetGroupProps := &ClusterSubnetGroupProps{
//   	Description: jsii.String("description"),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	VpcSubnets: &SubnetSelection{
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
// Experimental.
type ClusterSubnetGroupProps struct {
	// Description of the subnet group.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The VPC to place the subnet group in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The removal policy to apply when the subnet group are removed from the stack or replaced during an update.
	// Default: RemovalPolicy.RETAIN
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Which subnets within the VPC to associate with this group.
	// Default: - private subnets.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

