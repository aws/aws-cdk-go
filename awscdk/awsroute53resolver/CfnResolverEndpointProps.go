package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	ResolverEndpointType: jsii.String("resolverEndpointType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html
//
type CfnResolverEndpointProps struct {
	// Indicates whether the Resolver endpoint allows inbound or outbound DNS queries:.
	//
	// - `INBOUND` : allows DNS queries to your VPC from your network
	// - `OUTBOUND` : allows DNS queries from your VPC to your network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-direction
	//
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	//
	// The subnet ID uniquely identifies a VPC.
	//
	// > Even though the minimum is 1, Route 53 requires that you create at least two.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-ipaddresses
	//
	IpAddresses interface{} `field:"required" json:"ipAddresses" yaml:"ipAddresses"`
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN (Amazon Resource Name) for the Outpost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-outpostarn
	//
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The Amazon EC2 instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-preferredinstancetype
	//
	PreferredInstanceType *string `field:"optional" json:"preferredInstanceType" yaml:"preferredInstanceType"`
	// Protocols used for the endpoint. DoH-FIPS is applicable for inbound endpoints only.
	//
	// For an inbound endpoint you can apply the protocols as follows:
	//
	// - Do53 and DoH in combination.
	// - Do53 and DoH-FIPS in combination.
	// - Do53 alone.
	// - DoH alone.
	// - DoH-FIPS alone.
	// - None, which is treated as Do53.
	//
	// For an outbound endpoint you can apply the protocols as follows:
	//
	// - Do53 and DoH in combination.
	// - Do53 alone.
	// - DoH alone.
	// - None, which is treated as Do53.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-protocols
	//
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The Resolver endpoint IP address type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-resolverendpointtype
	//
	ResolverEndpointType *string `field:"optional" json:"resolverEndpointType" yaml:"resolverEndpointType"`
	// Route 53 Resolver doesn't support updating tags through CloudFormation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#cfn-route53resolver-resolverendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

