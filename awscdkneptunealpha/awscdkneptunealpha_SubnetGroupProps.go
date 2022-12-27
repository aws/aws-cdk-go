// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for creating a SubnetGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import neptune_alpha "github.com/aws/aws-cdk-go/awscdkneptunealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   subnetGroupProps := &subnetGroupProps{
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	subnetGroupName: jsii.String("subnetGroupName"),
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

