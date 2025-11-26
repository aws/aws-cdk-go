package previewawsecsmixins


// The storage configuration for Amazon ECS Managed Instances.
//
// This defines the root volume configuration for the instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedInstancesStorageConfigurationProperty := &ManagedInstancesStorageConfigurationProperty{
//   	StorageSizeGiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesstorageconfiguration.html
//
type CfnCapacityProviderPropsMixin_ManagedInstancesStorageConfigurationProperty struct {
	// The size of the tasks volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstancesstorageconfiguration.html#cfn-ecs-capacityprovider-managedinstancesstorageconfiguration-storagesizegib
	//
	StorageSizeGiB *float64 `field:"optional" json:"storageSizeGiB" yaml:"storageSizeGiB"`
}

