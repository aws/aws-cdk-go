package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The configuration of the Amazon VPCs for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup SecurityGroup
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var vpc Vpc
//
//   vpcConfig := &VpcConfig{
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
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
type VpcConfig struct {
	// Defines the virtual networking environment for this cluster.
	//
	// Must have at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The security groups associated with the cluster.
	//
	// You can specify up to 5 security groups.
	// Default: - create new security group.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets associated with the cluster.
	// Default: - the Vpc default strategy if not specified.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

