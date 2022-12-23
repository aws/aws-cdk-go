package awscloudtrail


// Options for adding an event selector.
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
type AddEventSelectorOptions struct {
	// An optional list of service event sources from which you do not want management events to be logged on your trail.
	ExcludeManagementEventSources *[]ManagementEventSources `field:"optional" json:"excludeManagementEventSources" yaml:"excludeManagementEventSources"`
	// Specifies whether the event selector includes management events for the trail.
	IncludeManagementEvents *bool `field:"optional" json:"includeManagementEvents" yaml:"includeManagementEvents"`
	// Specifies whether to log read-only events, write-only events, or all events.
	ReadWriteType ReadWriteType `field:"optional" json:"readWriteType" yaml:"readWriteType"`
}

