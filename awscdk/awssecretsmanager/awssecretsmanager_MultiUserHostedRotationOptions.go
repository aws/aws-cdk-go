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
//   multiUserHostedRotationOptions := &multiUserHostedRotationOptions{
//   	masterSecret: secret,
//
//   	// the properties below are optional
//   	excludeCharacters: jsii.String("excludeCharacters"),
//   	functionName: jsii.String("functionName"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	vpc: vpc,
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
type MultiUserHostedRotationOptions struct {
	// A string of the characters that you don't want in the password.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// A name for the Lambda created to rotate the secret.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// A list of security groups for the Lambda created to rotate the secret.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC where the Lambda rotation function will run.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The master secret for a multi user rotation scheme.
	MasterSecret ISecret `field:"required" json:"masterSecret" yaml:"masterSecret"`
}

