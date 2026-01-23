package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Recording configuration for browser.
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
//   			BucketName: recordingBucket.bucketName,
//   			ObjectKey: jsii.String("browser-recordings/"),
//   		},
//   	},
//   })
//
// Experimental.
type RecordingConfig struct {
	// Whether recording is enabled.
	// Default: false.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// S3 Location Configuration.
	// Default: - undefined.
	//
	// Experimental.
	S3Location *awss3.Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

