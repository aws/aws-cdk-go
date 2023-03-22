package awsivs


// The DestinationConfiguration property type describes the location where recorded videos will be stored.
//
// Each member represents a type of destination configuration. For recording, you define one and only one type of destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationConfigurationProperty := &destinationConfigurationProperty{
//   	s3: &s3DestinationConfigurationProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   }
//
type CfnRecordingConfiguration_DestinationConfigurationProperty struct {
	// An S3 destination configuration where recorded videos will be stored.
	//
	// See the [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.html) property type for more information.
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

