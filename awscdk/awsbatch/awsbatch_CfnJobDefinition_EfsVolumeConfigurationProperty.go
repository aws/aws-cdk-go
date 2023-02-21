package awsbatch


// This is used when you're using an Amazon Elastic File System file system for job storage.
//
// For more information, see [Amazon EFS Volumes](https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsVolumeConfigurationProperty := &efsVolumeConfigurationProperty{
//   	fileSystemId: jsii.String("fileSystemId"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfigProperty{
//   		accessPointId: jsii.String("accessPointId"),
//   		iam: jsii.String("iam"),
//   	},
//   	rootDirectory: jsii.String("rootDirectory"),
//   	transitEncryption: jsii.String("transitEncryption"),
//   	transitEncryptionPort: jsii.Number(123),
//   }
//
type CfnJobDefinition_EfsVolumeConfigurationProperty struct {
	// The Amazon EFS file system ID to use.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// If this parameter is omitted, the root of the Amazon EFS volume is used instead. Specifying `/` has the same effect as omitting this parameter. The maximum length is 4,096 characters.
	//
	// > If an EFS access point is specified in the `authorizationConfig` , the root directory parameter must either be omitted or set to `/` , which enforces the path set on the Amazon EFS access point.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Determines whether to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be enabled if Amazon EFS IAM authorization is used. If this parameter is omitted, the default value of `DISABLED` is used. For more information, see [Encrypting data in transit](https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html) in the *Amazon Elastic File System User Guide* .
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// If you don't specify a transit encryption port, it uses the port selection strategy that the Amazon EFS mount helper uses. The value must be between 0 and 65,535. For more information, see [EFS mount helper](https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html) in the *Amazon Elastic File System User Guide* .
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

