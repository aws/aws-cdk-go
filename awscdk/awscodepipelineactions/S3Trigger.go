package awscodepipelineactions


// How should the S3 Action detect changes.
//
// This is the type of the `S3SourceAction.trigger` property.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceBucket bucket
//
//   sourceOutput := codepipeline.NewArtifact()
//   key := "some/key.zip"
//   trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
//   trail.AddS3EventSelector([]s3EventSelector{
//   	&s3EventSelector{
//   		Bucket: sourceBucket,
//   		ObjectPrefix: key,
//   	},
//   }, &AddEventSelectorOptions{
//   	ReadWriteType: cloudtrail.ReadWriteType_WRITE_ONLY,
//   })
//   sourceAction := codepipeline_actions.NewS3SourceAction(&S3SourceActionProps{
//   	ActionName: jsii.String("S3Source"),
//   	BucketKey: key,
//   	Bucket: sourceBucket,
//   	Output: sourceOutput,
//   	Trigger: codepipeline_actions.S3Trigger_EVENTS,
//   })
//
type S3Trigger string

const (
	// The Action will never detect changes - the Pipeline it's part of will only begin a run when explicitly started.
	S3Trigger_NONE S3Trigger = "NONE"
	// CodePipeline will poll S3 to detect changes.
	//
	// This is the default method of detecting changes.
	S3Trigger_POLL S3Trigger = "POLL"
	// CodePipeline will use CloudWatch Events to be notified of changes.
	//
	// Note that the Bucket that the Action uses needs to be part of a CloudTrail Trail
	// for the events to be delivered.
	S3Trigger_EVENTS S3Trigger = "EVENTS"
)

