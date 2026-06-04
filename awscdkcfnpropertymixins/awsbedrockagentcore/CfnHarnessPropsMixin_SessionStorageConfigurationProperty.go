package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sessionStorageConfigurationProperty := &SessionStorageConfigurationProperty{
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-sessionstorageconfiguration.html
//
type CfnHarnessPropsMixin_SessionStorageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-sessionstorageconfiguration.html#cfn-bedrockagentcore-harness-sessionstorageconfiguration-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

