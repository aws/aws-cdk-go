package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedInstancesStorageConfigurationProperty := &ManagedInstancesStorageConfigurationProperty{
//   	StorageSizeGiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesstorageconfiguration.html
//
type CfnCapacityProvider_ManagedInstancesStorageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesstorageconfiguration.html#cfn-ecs-capacityprovider-managedinstancesstorageconfiguration-storagesizegib
	//
	StorageSizeGiB *float64 `field:"required" json:"storageSizeGiB" yaml:"storageSizeGiB"`
}

