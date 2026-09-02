package awsbedrockagentcore


// Configuration for a CapacityProvider-managed volume to mount into the agent runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   capacityProviderVolumeConfigurationProperty := &CapacityProviderVolumeConfigurationProperty{
//   	MountPath: jsii.String("mountPath"),
//   	VolumeName: jsii.String("volumeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-capacityprovidervolumeconfiguration.html
//
type CfnRuntimePropsMixin_CapacityProviderVolumeConfigurationProperty struct {
	// Mount path for filesystem configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-capacityprovidervolumeconfiguration.html#cfn-bedrockagentcore-runtime-capacityprovidervolumeconfiguration-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// Name of the capacity provider volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-capacityprovidervolumeconfiguration.html#cfn-bedrockagentcore-runtime-capacityprovidervolumeconfiguration-volumename
	//
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

