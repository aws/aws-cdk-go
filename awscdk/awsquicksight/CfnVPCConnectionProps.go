package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPCConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCConnectionProps := &CfnVPCConnectionProps{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DnsResolvers: []*string{
//   		jsii.String("dnsResolvers"),
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConnectionId: jsii.String("vpcConnectionId"),
//   }
//
type CfnVPCConnectionProps struct {
	// `AWS::QuickSight::VPCConnection.AvailabilityStatus`.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
	// `AWS::QuickSight::VPCConnection.AwsAccountId`.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// `AWS::QuickSight::VPCConnection.DnsResolvers`.
	DnsResolvers *[]*string `field:"optional" json:"dnsResolvers" yaml:"dnsResolvers"`
	// `AWS::QuickSight::VPCConnection.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::QuickSight::VPCConnection.RoleArn`.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// `AWS::QuickSight::VPCConnection.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `AWS::QuickSight::VPCConnection.SubnetIds`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// `AWS::QuickSight::VPCConnection.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::QuickSight::VPCConnection.VPCConnectionId`.
	VpcConnectionId *string `field:"optional" json:"vpcConnectionId" yaml:"vpcConnectionId"`
}

