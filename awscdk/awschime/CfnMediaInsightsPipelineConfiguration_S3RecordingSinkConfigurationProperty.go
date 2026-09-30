package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3RecordingSinkConfigurationProperty := &S3RecordingSinkConfigurationProperty{
//   	Destination: jsii.String("destination"),
//   	RecordingFileFormat: jsii.String("recordingFileFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-s3recordingsinkconfiguration.html
//
type CfnMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty struct {
	// The default URI of the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-s3recordingsinkconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-s3recordingsinkconfiguration-destination
	//
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The recording file format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-s3recordingsinkconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-s3recordingsinkconfiguration-recordingfileformat
	//
	RecordingFileFormat *string `field:"optional" json:"recordingFileFormat" yaml:"recordingFileFormat"`
}

