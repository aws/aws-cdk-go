package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	StorageCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-metadataconfiguration.html
//
type CfnFileCache_MetadataConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-metadataconfiguration.html#cfn-fsx-filecache-metadataconfiguration-storagecapacity
	//
	StorageCapacity *float64 `field:"required" json:"storageCapacity" yaml:"storageCapacity"`
}

