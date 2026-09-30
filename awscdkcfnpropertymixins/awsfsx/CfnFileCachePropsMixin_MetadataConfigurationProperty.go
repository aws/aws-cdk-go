package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metadataConfigurationProperty := &MetadataConfigurationProperty{
//   	StorageCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-metadataconfiguration.html
//
type CfnFileCachePropsMixin_MetadataConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-metadataconfiguration.html#cfn-fsx-filecache-metadataconfiguration-storagecapacity
	//
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
}

