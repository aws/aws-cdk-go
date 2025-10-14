package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options for a new DatabaseProxyEndpoint.
//
// Example:
//   var vpc vpc
//   var secrets []secret
//   var dbInstance databaseInstance
//
//
//   proxy := dbInstance.AddProxy(jsii.String("Proxy"), &DatabaseProxyOptions{
//   	Secrets: Secrets,
//   	Vpc: Vpc,
//   })
//
//   // Add a reader endpoint
//   proxy.AddEndpoint(jsii.String("ProxyEndpoint"), &DatabaseProxyEndpointOptions{
//   	Vpc: Vpc,
//   	TargetRole: rds.ProxyEndpointTargetRole_READ_ONLY,
//   })
//
type DatabaseProxyEndpointOptions struct {
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
}

