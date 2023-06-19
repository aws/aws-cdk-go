package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// FileSystem configurations for the Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connections connections
//   var dependable iDependable
//   var policyStatement policyStatement
//
//   fileSystemConfig := &FileSystemConfig{
//   	Arn: jsii.String("arn"),
//   	LocalMountPath: jsii.String("localMountPath"),
//
//   	// the properties below are optional
//   	Connections: connections,
//   	Dependency: []*iDependable{
//   		dependable,
//   	},
//   	Policies: []*policyStatement{
//   		policyStatement,
//   	},
//   }
//
// Experimental.
type FileSystemConfig struct {
	// ARN of the access point.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// mount path in the lambda runtime environment.
	// Experimental.
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
	// connections object used to allow ingress traffic from lambda function.
	// Experimental.
	Connections awsec2.Connections `field:"optional" json:"connections" yaml:"connections"`
	// array of IDependable that lambda function depends on.
	// Experimental.
	Dependency *[]awscdk.IDependable `field:"optional" json:"dependency" yaml:"dependency"`
	// additional IAM policies required for the lambda function.
	// Experimental.
	Policies *[]awsiam.PolicyStatement `field:"optional" json:"policies" yaml:"policies"`
}

