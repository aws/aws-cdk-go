package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Construction properties for a DatabaseProxyEndpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var databaseProxy databaseProxy
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   databaseProxyEndpointProps := &DatabaseProxyEndpointProps{
//   	DbProxy: databaseProxy,
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	DbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	TargetRole: awscdk.Aws_rds.ProxyEndpointTargetRole_READ_WRITE,
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
type DatabaseProxyEndpointProps struct {
	// The VPC of the DB proxy endpoint.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The name of the DB proxy endpoint.
	// Default: - a CDK generated name.
	//
	DbProxyEndpointName *string `field:"optional" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The VPC security groups to associate with the new proxy endpoint.
	// Default: - Default security group for the VPC.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	// Default: - ProxyEndpointTargetRole.READ_WRITE
	//
	TargetRole ProxyEndpointTargetRole `field:"optional" json:"targetRole" yaml:"targetRole"`
	// The subnets of DB proxy endpoint.
	// Default: - the VPC default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The DB proxy associated with the DB proxy endpoint.
	DbProxy IDatabaseProxy `field:"required" json:"dbProxy" yaml:"dbProxy"`
}

