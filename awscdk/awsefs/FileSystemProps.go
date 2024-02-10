package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties of EFS FileSystem.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   role := iam.NewRole(this, jsii.String("ClientRole"), &RoleProps{
//   	AssumedBy: iam.NewAnyPrincipal(),
//   })
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	AllowAnonymousAccess: jsii.Boolean(true),
//   })
//
//   fileSystem.grantRead(role)
//
type FileSystemProps struct {
	// VPC to launch the file system in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Allow access from anonymous client that doesn't use IAM authentication.
	// Default: false when using `grantRead`, `grantWrite`, `grantRootAccess`
	// or set `@aws-cdk/aws-efs:denyAnonymousAccess` feature flag, otherwise true.
	//
	AllowAnonymousAccess *bool `field:"optional" json:"allowAnonymousAccess" yaml:"allowAnonymousAccess"`
	// Whether to enable automatic backups for the file system.
	// Default: false.
	//
	EnableAutomaticBackups *bool `field:"optional" json:"enableAutomaticBackups" yaml:"enableAutomaticBackups"`
	// Defines if the data at rest in the file system is encrypted or not.
	// Default: - If your application has the '@aws-cdk/aws-efs:defaultEncryptionAtRest' feature flag set, the default is true, otherwise, the default is false.
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The file system's name.
	// Default: - CDK generated name.
	//
	FileSystemName *string `field:"optional" json:"fileSystemName" yaml:"fileSystemName"`
	// File system policy is an IAM resource policy used to control NFS access to an EFS file system.
	// Default: none.
	//
	FileSystemPolicy awsiam.PolicyDocument `field:"optional" json:"fileSystemPolicy" yaml:"fileSystemPolicy"`
	// The KMS key used for encryption.
	//
	// This is required to encrypt the data at rest if.
	// Default: - if 'encrypted' is true, the default key for EFS (/aws/elasticfilesystem) is used.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// A policy used by EFS lifecycle management to transition files to the Infrequent Access (IA) storage class.
	// Default: - None. EFS will not transition files to the IA storage class.
	//
	LifecyclePolicy LifecyclePolicy `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// Whether this is a One Zone file system.
	//
	// If enabled, `performanceMode` must be set to `GENERAL_PURPOSE` and `vpcSubnets` cannot be set.
	// Default: false.
	//
	OneZone *bool `field:"optional" json:"oneZone" yaml:"oneZone"`
	// A policy used by EFS lifecycle management to transition files from Infrequent Access (IA) storage class to primary storage class.
	// Default: - None. EFS will not transition files from IA storage to primary storage.
	//
	OutOfInfrequentAccessPolicy OutOfInfrequentAccessPolicy `field:"optional" json:"outOfInfrequentAccessPolicy" yaml:"outOfInfrequentAccessPolicy"`
	// The performance mode that the file system will operate under.
	//
	// An Amazon EFS file system's performance mode can't be changed after the file system has been created.
	// Updating this property will replace the file system.
	// Default: PerformanceMode.GENERAL_PURPOSE
	//
	PerformanceMode PerformanceMode `field:"optional" json:"performanceMode" yaml:"performanceMode"`
	// Provisioned throughput for the file system.
	//
	// This is a required property if the throughput mode is set to PROVISIONED.
	// Must be at least 1MiB/s.
	// Default: - none, errors out.
	//
	ProvisionedThroughputPerSecond awscdk.Size `field:"optional" json:"provisionedThroughputPerSecond" yaml:"provisionedThroughputPerSecond"`
	// The removal policy to apply to the file system.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Whether to enable the filesystem's replication overwrite protection or not.
	//
	// Set false if you want to create a read-only filesystem for use as a replication destination.
	// See: https://docs.aws.amazon.com/efs/latest/ug/replication-use-cases.html#replicate-existing-destination
	//
	// Default: ReplicationOverwriteProtection.ENABLED
	//
	ReplicationOverwriteProtection ReplicationOverwriteProtection `field:"optional" json:"replicationOverwriteProtection" yaml:"replicationOverwriteProtection"`
	// Security Group to assign to this file system.
	// Default: - creates new security group which allows all outbound traffic.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Enum to mention the throughput mode of the file system.
	// Default: ThroughputMode.BURSTING
	//
	ThroughputMode ThroughputMode `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage.
	//
	// Metadata operations such as listing the contents of a directory don't count as file access events.
	// Default: - None. EFS will not transition files to Archive storage class.
	//
	TransitionToArchivePolicy LifecyclePolicy `field:"optional" json:"transitionToArchivePolicy" yaml:"transitionToArchivePolicy"`
	// Which subnets to place the mount target in the VPC.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

