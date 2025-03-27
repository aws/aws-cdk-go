package awsecs


// This parameter is specified when you're using an Amazon Elastic File System file system for task storage.
//
// For more information, see [Amazon EFS volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eFSVolumeConfigurationProperty := &EFSVolumeConfigurationProperty{
//   	FilesystemId: jsii.String("filesystemId"),
//
//   	// the properties below are optional
//   	AuthorizationConfig: &AuthorizationConfigProperty{
//   		AccessPointId: jsii.String("accessPointId"),
//   		Iam: jsii.String("iam"),
//   	},
//   	RootDirectory: jsii.String("rootDirectory"),
//   	TransitEncryption: jsii.String("transitEncryption"),
//   	TransitEncryptionPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html
//
type CfnTaskDefinition_EFSVolumeConfigurationProperty struct {
	// The Amazon EFS file system ID to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html#cfn-ecs-taskdefinition-efsvolumeconfiguration-filesystemid
	//
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// The authorization configuration details for the Amazon EFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html#cfn-ecs-taskdefinition-efsvolumeconfiguration-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying `/` will have the same effect as omitting this parameter.
	//
	// > If an EFS access point is specified in the `authorizationConfig` , the root directory parameter must either be omitted or set to `/` which will enforce the path set on the EFS access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html#cfn-ecs-taskdefinition-efsvolumeconfiguration-rootdirectory
	//
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Determines whether to use encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be turned on if Amazon EFS IAM authorization is used. If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Encrypting data in transit](https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html) in the *Amazon Elastic File System User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html#cfn-ecs-taskdefinition-efsvolumeconfiguration-transitencryption
	//
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses. For more information, see [EFS mount helper](https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html) in the *Amazon Elastic File System User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html#cfn-ecs-taskdefinition-efsvolumeconfiguration-transitencryptionport
	//
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

