package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// FileSystem configurations for the Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
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
type FileSystemConfig struct {
	// ARN of the access point.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// mount path in the lambda runtime environment.
	LocalMountPath *string `field:"required" json:"localMountPath" yaml:"localMountPath"`
	// connections object used to allow ingress traffic from lambda function.
	// Default: - no connections required to add extra ingress rules for Lambda function.
	//
	Connections awsec2.Connections `field:"optional" json:"connections" yaml:"connections"`
	// array of IDependable that lambda function depends on.
	// Default: - no dependency.
	//
	Dependency *[]constructs.IDependable `field:"optional" json:"dependency" yaml:"dependency"`
	// additional IAM policies required for the lambda function.
	// Default: - no additional policies required.
	//
	Policies *[]awsiam.PolicyStatement `field:"optional" json:"policies" yaml:"policies"`
}

