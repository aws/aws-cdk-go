package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filesystemConfigurationProperty := &FilesystemConfigurationProperty{
//   	SessionStorage: &SessionStorageConfigurationProperty{
//   		MountPath: jsii.String("mountPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html
//
type CfnHarness_FilesystemConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html#cfn-bedrockagentcore-harness-filesystemconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"required" json:"sessionStorage" yaml:"sessionStorage"`
}

