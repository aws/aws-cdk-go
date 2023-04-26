package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Single user hosted rotation options.
//
// Example:
//   var myVpc vpc
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
// Experimental.
type SingleUserHostedRotationOptions struct {
	// A name for the Lambda created to rotate the secret.
	// Experimental.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// A list of security groups for the Lambda created to rotate the secret.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The VPC where the Lambda rotation function will run.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The type of subnets in the VPC where the Lambda rotation function will run.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

