package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties of EFS FileSystem.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	LifecyclePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	PerformanceMode: efs.PerformanceMode_GENERAL_PURPOSE,
//   	 // default
//   	OutOfInfrequentAccessPolicy: efs.OutOfInfrequentAccessPolicy_AFTER_1_ACCESS,
//   })
//
// Experimental.
type FileSystemProps struct {
	// VPC to launch the file system in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to enable automatic backups for the file system.
	// Experimental.
	EnableAutomaticBackups *bool `field:"optional" json:"enableAutomaticBackups" yaml:"enableAutomaticBackups"`
	// Defines if the data at rest in the file system is encrypted or not.
	// Experimental.
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The file system's name.
	// Experimental.
	FileSystemName *string `field:"optional" json:"fileSystemName" yaml:"fileSystemName"`
	// The KMS key used for encryption.
	//
	// This is required to encrypt the data at rest if @encrypted is set to true.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// A policy used by EFS lifecycle management to transition files to the Infrequent Access (IA) storage class.
	// Experimental.
	LifecyclePolicy LifecyclePolicy `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// A policy used by EFS lifecycle management to transition files from Infrequent Access (IA) storage class to primary storage class.
	// Experimental.
	OutOfInfrequentAccessPolicy OutOfInfrequentAccessPolicy `field:"optional" json:"outOfInfrequentAccessPolicy" yaml:"outOfInfrequentAccessPolicy"`
	// The performance mode that the file system will operate under.
	//
	// An Amazon EFS file system's performance mode can't be changed after the file system has been created.
	// Updating this property will replace the file system.
	// Experimental.
	PerformanceMode PerformanceMode `field:"optional" json:"performanceMode" yaml:"performanceMode"`
	// Provisioned throughput for the file system.
	//
	// This is a required property if the throughput mode is set to PROVISIONED.
	// Must be at least 1MiB/s.
	// Experimental.
	ProvisionedThroughputPerSecond awscdk.Size `field:"optional" json:"provisionedThroughputPerSecond" yaml:"provisionedThroughputPerSecond"`
	// The removal policy to apply to the file system.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Enum to mention the throughput mode of the file system.
	// Experimental.
	ThroughputMode ThroughputMode `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// Which subnets to place the mount target in the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

