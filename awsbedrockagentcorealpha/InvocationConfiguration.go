package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Invocation configuration for self managed memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic Topic
//
//   invocationConfiguration := &InvocationConfiguration{
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   	Topic: topic,
//   }
//
// Experimental.
type InvocationConfiguration struct {
	// S3 Location Configuration.
	// Experimental.
	S3Location *awss3.Location `field:"required" json:"s3Location" yaml:"s3Location"`
	// SNS Topic Configuration.
	// Experimental.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
}

