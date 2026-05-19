package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Recording configuration for browser.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   recordingConfig := &RecordingConfig{
//   	Enabled: jsii.Boolean(false),
//   	S3Location: &Location{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
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

