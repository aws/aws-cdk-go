package awsecs


// The configuration for an Elastic FileSystem volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsVolumeConfiguration := &efsVolumeConfiguration{
//   	fileSystemId: jsii.String("fileSystemId"),
//
//   	// the properties below are optional
//   	authorizationConfig: &authorizationConfig{
//   		accessPointId: jsii.String("accessPointId"),
//   		iam: jsii.String("iam"),
//   	},
//   	rootDirectory: jsii.String("rootDirectory"),
//   	transitEncryption: jsii.String("transitEncryption"),
//   	transitEncryptionPort: jsii.Number(123),
//   }
//
type EfsVolumeConfiguration struct {
	// The Amazon EFS file system ID to use.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The directory within the Amazon EFS file system to mount as the root directory inside the host.
	//
	// Specifying / will have the same effect as omitting this parameter.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server.
	//
	// Transit encryption must be enabled if Amazon EFS IAM authorization is used.
	//
	// Valid values: ENABLED | DISABLED.
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server.
	//
	// EFS mount helper uses.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

