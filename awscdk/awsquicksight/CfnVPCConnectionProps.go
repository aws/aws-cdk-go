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
	// The availability status of the VPC connection.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
	// The AWS account ID of the account where you want to create a new VPC connection.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// A list of IP addresses of DNS resolver endpoints for the VPC connection.
	DnsResolvers *[]*string `field:"optional" json:"dnsResolvers" yaml:"dnsResolvers"`
	// The display name for the VPC connection.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the IAM role associated with the VPC connection.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon EC2 security group IDs associated with the VPC connection.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs for the VPC connection.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// A map of the key-value pairs for the resource tag or tags assigned to the VPC connection.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC connection that you're creating.
	//
	// This ID is a unique identifier for each AWS Region in an AWS account.
	VpcConnectionId *string `field:"optional" json:"vpcConnectionId" yaml:"vpcConnectionId"`
}

