package awsbedrockagentcore


// Filesystem configuration for the runtime.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html
//
type CfnRuntime_FilesystemConfigurationProperty struct {
	// Configuration for session storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html#cfn-bedrockagentcore-runtime-filesystemconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

