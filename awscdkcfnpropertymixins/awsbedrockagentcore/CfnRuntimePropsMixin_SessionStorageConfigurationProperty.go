package awsbedrockagentcore


// Configuration for session storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sessionStorageConfigurationProperty := &SessionStorageConfigurationProperty{
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-sessionstorageconfiguration.html
//
type CfnRuntimePropsMixin_SessionStorageConfigurationProperty struct {
	// Mount path for filesystem configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-sessionstorageconfiguration.html#cfn-bedrockagentcore-runtime-sessionstorageconfiguration-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

