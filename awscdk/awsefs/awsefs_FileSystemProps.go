package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties of EFS FileSystem.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &fileSystemProps{
//   	vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	lifecyclePolicy: efs.lifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	performanceMode: efs.performanceMode_GENERAL_PURPOSE,
//   	 // default
//   	outOfInfrequentAccessPolicy: efs.outOfInfrequentAccessPolicy_AFTER_1_ACCESS,
//   })
//
type FileSystemProps struct {
	// VPC to launch the file system in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to enable automatic backups for the file system.
	EnableAutomaticBackups *bool `field:"optional" json:"enableAutomaticBackups" yaml:"enableAutomaticBackups"`
	// Defines if the data at rest in the file system is encrypted or not.
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The file system's name.
	FileSystemName *string `field:"optional" json:"fileSystemName" yaml:"fileSystemName"`
	// The KMS key used for encryption.
	//
	// This is required to encrypt the data at rest if @encrypted is set to true.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// A policy used by EFS lifecycle management to transition files to the Infrequent Access (IA) storage class.
	LifecyclePolicy LifecyclePolicy `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// A policy used by EFS lifecycle management to transition files from Infrequent Access (IA) storage class to primary storage class.
	OutOfInfrequentAccessPolicy OutOfInfrequentAccessPolicy `field:"optional" json:"outOfInfrequentAccessPolicy" yaml:"outOfInfrequentAccessPolicy"`
	// The performance mode that the file system will operate under.
	//
	// An Amazon EFS file system's performance mode can't be changed after the file system has been created.
	// Updating this property will replace the file system.
	PerformanceMode PerformanceMode `field:"optional" json:"performanceMode" yaml:"performanceMode"`
	// Provisioned throughput for the file system.
	//
	// This is a required property if the throughput mode is set to PROVISIONED.
	// Must be at least 1MiB/s.
	ProvisionedThroughputPerSecond awscdk.Size `field:"optional" json:"provisionedThroughputPerSecond" yaml:"provisionedThroughputPerSecond"`
	// The removal policy to apply to the file system.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Enum to mention the throughput mode of the file system.
	ThroughputMode ThroughputMode `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// Which subnets to place the mount target in the VPC.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

