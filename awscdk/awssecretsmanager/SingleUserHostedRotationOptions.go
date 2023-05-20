package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Single user hosted rotation options.
//
// Example:
//   var myVpc iVpc
//   var dbConnections connections
//   var secret secret
//
//
//   myHostedRotation := secretsmanager.HostedRotation_MysqlSingleUser(&SingleUserHostedRotationOptions{
//   	Vpc: myVpc,
//   })
//   secret.addRotationSchedule(jsii.String("RotationSchedule"), &RotationScheduleOptions{
//   	HostedRotation: myHostedRotation,
//   })
//   dbConnections.AllowDefaultPortFrom(myHostedRotation)
//
type SingleUserHostedRotationOptions struct {
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
}

