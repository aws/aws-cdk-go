package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Construction properties of the `S3SourceAction S3 source Action`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cloudtrail "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.addS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		bucket: sourceBucket,
//   		objectPrefix: key,
//   	},
//   }, &addEventSelectorOptions{
//   	readWriteType: cloudtrail.readWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucketKey: key,
//   	bucket: sourceBucket,
//   	output: sourceOutput,
//   	trigger: codepipeline_actions.s3Trigger_EVENTS,
//   })
//
type S3SourceActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Amazon S3 bucket that stores the source code.
	//
	// If you import an encrypted bucket in your stack, please specify
	// the encryption key at import time by using `Bucket.fromBucketAttributes()` method.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The key within the S3 bucket that stores the source code.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "path/to/file.zip"
	//
	BucketKey *string `field:"required" json:"bucketKey" yaml:"bucketKey"`
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	Trigger S3Trigger `field:"optional" json:"trigger" yaml:"trigger"`
}

