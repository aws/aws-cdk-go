package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Multi user hosted rotation options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   multiUserHostedRotationOptions := &MultiUserHostedRotationOptions{
//   	MasterSecret: secret,
//
//   	// the properties below are optional
//   	ExcludeCharacters: jsii.String("excludeCharacters"),
//   	FunctionName: jsii.String("functionName"),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	Vpc: vpc,
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
type MultiUserHostedRotationOptions struct {
	// A string of the characters that you don't want in the password.
	// Default: the same exclude characters as the ones used for the
	// secret or " %+~`#$&*()|[]{}:;<>?!'/@\"\\"
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// A name for the Lambda created to rotate the secret.
	// Default: - a CloudFormation generated name.
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// A list of security groups for the Lambda created to rotate the secret.
	// Default: - a new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC where the Lambda rotation function will run.
	// Default: - the Lambda is not deployed in a VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The master secret for a multi user rotation scheme.
	MasterSecret ISecret `field:"required" json:"masterSecret" yaml:"masterSecret"`
}

