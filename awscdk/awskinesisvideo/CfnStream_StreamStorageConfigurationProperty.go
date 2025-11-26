package awskinesisvideo


// Configuration for the storage tier of the Kinesis Video Stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamStorageConfigurationProperty := &StreamStorageConfigurationProperty{
//   	DefaultStorageTier: jsii.String("defaultStorageTier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisvideo-stream-streamstorageconfiguration.html
//
type CfnStream_StreamStorageConfigurationProperty struct {
	// The storage tier for the Kinesis Video Stream.
	//
	// Determines the storage class used for stream data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisvideo-stream-streamstorageconfiguration.html#cfn-kinesisvideo-stream-streamstorageconfiguration-defaultstoragetier
	//
	// Default: - "HOT".
	//
	DefaultStorageTier *string `field:"optional" json:"defaultStorageTier" yaml:"defaultStorageTier"`
}

