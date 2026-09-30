package awsiotsitewise


// Configuration for ephemeral storage attached to the container task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ephemeralStorageConfigurationProperty := &EphemeralStorageConfigurationProperty{
//   	StorageClass: jsii.String("storageClass"),
//   	StorageSizeInGiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-ephemeralstorageconfiguration.html
//
type CfnTaskPropsMixin_EphemeralStorageConfigurationProperty struct {
	// The storage type that determines I/O performance characteristics.
	//
	// Family name indicates workload pattern, level number indicates performance within that family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-ephemeralstorageconfiguration.html#cfn-iotsitewise-task-ephemeralstorageconfiguration-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Storage volume size in GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-task-ephemeralstorageconfiguration.html#cfn-iotsitewise-task-ephemeralstorageconfiguration-storagesizeingib
	//
	StorageSizeInGiB *float64 `field:"optional" json:"storageSizeInGiB" yaml:"storageSizeInGiB"`
}

