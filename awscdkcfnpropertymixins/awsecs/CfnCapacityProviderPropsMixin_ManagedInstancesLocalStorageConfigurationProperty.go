package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   managedInstancesLocalStorageConfigurationProperty := &ManagedInstancesLocalStorageConfigurationProperty{
//   	UseLocalStorage: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstanceslocalstorageconfiguration.html
//
type CfnCapacityProviderPropsMixin_ManagedInstancesLocalStorageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-managedinstanceslocalstorageconfiguration.html#cfn-ecs-capacityprovider-managedinstanceslocalstorageconfiguration-uselocalstorage
	//
	UseLocalStorage interface{} `field:"optional" json:"useLocalStorage" yaml:"useLocalStorage"`
}

