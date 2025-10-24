package awss3


// An interface that represents the location of a specific object in an S3 Bucket.
//
// Example:
//   // Create an S3 bucket for recordings
//   recordingBucket := s3.NewBucket(this, jsii.String("RecordingBucket"), &BucketProps{
//   	BucketName: jsii.String("my-browser-recordings"),
//   	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
//   })
//
//   // Create browser with recording enabled
//   browser := agentcore.NewBrowserCustom(this, jsii.String("MyBrowser"), &BrowserCustomProps{
//   	BrowserCustomName: jsii.String("my_browser"),
//   	Description: jsii.String("Browser with recording enabled"),
//   	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingPublicNetwork(),
//   	RecordingConfig: &RecordingConfig{
//   		Enabled: jsii.Boolean(true),
//   		S3Location: &Location{
//   			BucketName: recordingBucket.BucketName,
//   			ObjectKey: jsii.String("browser-recordings/"),
//   		},
//   	},
//   })
//
type Location struct {
	// The name of the S3 Bucket the object is in.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The path inside the Bucket where the object is located at.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

