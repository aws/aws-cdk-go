package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstancesStorageConfigurationProperty := &ManagedInstancesStorageConfigurationProperty{
//   	StorageSizeGiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesstorageconfiguration.html
//
type CfnComputeEnvironment_ManagedInstancesStorageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-managedinstancesstorageconfiguration.html#cfn-batch-computeenvironment-managedinstancesstorageconfiguration-storagesizegib
	//
	StorageSizeGiB *float64 `field:"optional" json:"storageSizeGiB" yaml:"storageSizeGiB"`
}

