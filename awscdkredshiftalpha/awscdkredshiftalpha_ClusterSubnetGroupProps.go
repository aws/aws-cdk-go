// The CDK Construct Library for AWS::Redshift
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
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   clusterSubnetGroupProps := &clusterSubnetGroupProps{
//   	description: jsii.String("description"),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
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
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Which subnets within the VPC to associate with this group.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

