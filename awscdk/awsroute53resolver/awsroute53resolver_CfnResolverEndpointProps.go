package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnResolverEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResolverEndpointProps := &CfnResolverEndpointProps{
//   	Direction: jsii.String("direction"),
//   	IpAddresses: []interface{}{
//   		&IpAddressRequestProperty{
//   			SubnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			Ip: jsii.String("ip"),
//   			Ipv6: jsii.String("ipv6"),
//   		},
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	OutpostArn: jsii.String("outpostArn"),
//   	PreferredInstanceType: jsii.String("preferredInstanceType"),
//   	ResolverEndpointType: jsii.String("resolverEndpointType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnResolverEndpointProps struct {
	// Indicates whether the Resolver endpoint allows inbound or outbound DNS queries:.
	//
	// - `INBOUND` : allows DNS queries to your VPC from your network
	// - `OUTBOUND` : allows DNS queries from your VPC to your network.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	//
	// The subnet ID uniquely identifies a VPC.
	IpAddresses interface{} `field:"required" json:"ipAddresses" yaml:"ipAddresses"`
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Route53Resolver::ResolverEndpoint.OutpostArn`.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// `AWS::Route53Resolver::ResolverEndpoint.PreferredInstanceType`.
	PreferredInstanceType *string `field:"optional" json:"preferredInstanceType" yaml:"preferredInstanceType"`
	// The Resolver endpoint IP address type.
	ResolverEndpointType *string `field:"optional" json:"resolverEndpointType" yaml:"resolverEndpointType"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

