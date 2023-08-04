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
//   myFileSystemPolicy := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("elasticfilesystem:ClientWrite"),
//   				jsii.String("elasticfilesystem:ClientMount"),
//   			},
//   			Principals: []iPrincipal{
//   				iam.NewAccountRootPrincipal(),
//   			},
//   			Resources: []*string{
//   				jsii.String("*"),
//   			},
//   			Conditions: map[string]interface{}{
//   				"Bool": map[string]*string{
//   					"elasticfilesystem:AccessedViaMountTarget": jsii.String("true"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	FileSystemPolicy: myFileSystemPolicy,
//   })
//
type FileSystemProps struct {
	// VPC to launch the file system in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
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
	// Security Group to assign to this file system.
	// Default: - creates new security group which allows all outbound traffic.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Enum to mention the throughput mode of the file system.
	// Default: ThroughputMode.BURSTING
	//
	ThroughputMode ThroughputMode `field:"optional" json:"throughputMode" yaml:"throughputMode"`
	// Which subnets to place the mount target in the VPC.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

