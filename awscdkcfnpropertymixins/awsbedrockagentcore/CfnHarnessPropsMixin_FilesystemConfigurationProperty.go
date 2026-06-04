package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filesystemConfigurationProperty := &FilesystemConfigurationProperty{
//   	SessionStorage: &SessionStorageConfigurationProperty{
//   		MountPath: jsii.String("mountPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html
//
type CfnHarnessPropsMixin_FilesystemConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html#cfn-bedrockagentcore-harness-filesystemconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

