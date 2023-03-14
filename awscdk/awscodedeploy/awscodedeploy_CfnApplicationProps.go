package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	applicationName: jsii.String("applicationName"),
//   	computePlatform: jsii.String("computePlatform"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// A name for the application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > Updates to `ApplicationName` are not supported.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The compute platform that CodeDeploy deploys the application to.
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// The metadata that you apply to CodeDeploy applications to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

