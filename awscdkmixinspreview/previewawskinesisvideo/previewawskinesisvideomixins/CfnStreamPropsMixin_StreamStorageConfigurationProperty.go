package previewawskinesisvideomixins


// The configuration for stream storage, including the default storage tier for stream data.
//
// This configuration determines how stream data is stored and accessed, with different tiers offering varying levels of performance and cost optimization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamStorageConfigurationProperty := &StreamStorageConfigurationProperty{
//   	DefaultStorageTier: jsii.String("defaultStorageTier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisvideo-stream-streamstorageconfiguration.html
//
type CfnStreamPropsMixin_StreamStorageConfigurationProperty struct {
	// The default storage tier for the stream data.
	//
	// This setting determines the storage class used for stream data, affecting both performance characteristics and storage costs.
	//
	// Available storage tiers:
	//
	// - `HOT` - Optimized for frequent access with the lowest latency and highest performance. Ideal for real-time applications and frequently accessed data.
	// - `WARM` - Balanced performance and cost for moderately accessed data. Suitable for data that is accessed regularly but not continuously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisvideo-stream-streamstorageconfiguration.html#cfn-kinesisvideo-stream-streamstorageconfiguration-defaultstoragetier
	//
	// Default: - "HOT".
	//
	DefaultStorageTier *string `field:"optional" json:"defaultStorageTier" yaml:"defaultStorageTier"`
}

