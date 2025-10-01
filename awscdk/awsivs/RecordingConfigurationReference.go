package awsivs


// A reference to a RecordingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordingConfigurationReference := &RecordingConfigurationReference{
//   	RecordingConfigurationArn: jsii.String("recordingConfigurationArn"),
//   }
//
type RecordingConfigurationReference struct {
	// The Arn of the RecordingConfiguration resource.
	RecordingConfigurationArn *string `field:"required" json:"recordingConfigurationArn" yaml:"recordingConfigurationArn"`
}

