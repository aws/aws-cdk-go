package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filesystemConfigurationProperty := &FilesystemConfigurationProperty{
//   	EfsAccessPoint: &EfsAccessPointConfigurationProperty{
//   		AccessPointArn: jsii.String("accessPointArn"),
//   		MountPath: jsii.String("mountPath"),
//   	},
//   	S3FilesAccessPoint: &S3FilesAccessPointConfigurationProperty{
//   		AccessPointArn: jsii.String("accessPointArn"),
//   		MountPath: jsii.String("mountPath"),
//   	},
//   	SessionStorage: &SessionStorageConfigurationProperty{
//   		MountPath: jsii.String("mountPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html
//
type CfnHarness_FilesystemConfigurationProperty struct {
	// Configuration for an Amazon EFS access point to mount into the AgentCore Runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html#cfn-bedrockagentcore-harness-filesystemconfiguration-efsaccesspoint
	//
	EfsAccessPoint interface{} `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// Configuration for an Amazon S3 Files access point to mount into the AgentCore Runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html#cfn-bedrockagentcore-harness-filesystemconfiguration-s3filesaccesspoint
	//
	S3FilesAccessPoint interface{} `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-filesystemconfiguration.html#cfn-bedrockagentcore-harness-filesystemconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

