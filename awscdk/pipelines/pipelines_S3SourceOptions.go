package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions"
)

// Options for S3 sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceOptions := &s3SourceOptions{
//   	actionName: jsii.String("actionName"),
//   	trigger: awscdk.Aws_codepipeline_actions.s3Trigger_NONE,
//   }
//
// Experimental.
type S3SourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Experimental.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	// Experimental.
	Trigger awscodepipelineactions.S3Trigger `field:"optional" json:"trigger" yaml:"trigger"`
}

