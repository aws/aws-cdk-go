package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for S3 sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   s3SourceOptions := &S3SourceOptions{
//   	ActionName: jsii.String("actionName"),
//   	Role: role,
//   	Trigger: awscdk.Aws_codepipeline_actions.S3Trigger_NONE,
//   }
//
type S3SourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The role that will be assumed by the pipeline prior to executing the `S3Source` action.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	Trigger awscodepipelineactions.S3Trigger `field:"optional" json:"trigger" yaml:"trigger"`
}

