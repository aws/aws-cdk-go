package awsfsx


// The configuration for the optional provisioned SSD read cache on Amazon FSx for Lustre file systems that use the Intelligent-Tiering storage class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataReadCacheConfigurationProperty := &DataReadCacheConfigurationProperty{
//   	SizeGiB: jsii.Number(123),
//   	SizingMode: jsii.String("sizingMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-datareadcacheconfiguration.html
//
type CfnFileSystem_DataReadCacheConfigurationProperty struct {
	// Required if `SizingMode` is set to `USER_PROVISIONED` .
	//
	// Specifies the size of the file system's SSD read cache, in gibibytes (GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-datareadcacheconfiguration.html#cfn-fsx-filesystem-datareadcacheconfiguration-sizegib
	//
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// Specifies how the provisioned SSD read cache is sized, as follows:.
	//
	// - Set to `NO_CACHE` if you do not want to use an SSD read cache with your Intelligent-Tiering file system.
	// - Set to `USER_PROVISIONED` to specify the exact size of your SSD read cache.
	// - Set to `PROPORTIONAL_TO_THROUGHPUT_CAPACITY` to have your SSD read cache automatically sized based on your throughput capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-datareadcacheconfiguration.html#cfn-fsx-filesystem-datareadcacheconfiguration-sizingmode
	//
	SizingMode *string `field:"optional" json:"sizingMode" yaml:"sizingMode"`
}

