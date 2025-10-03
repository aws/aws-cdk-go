package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties specific to the Lustre version of the FSx file system.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//   var bucket bucket
//
//
//   lustreConfiguration := map[string]interface{}{
//   	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
//   	"exportPath": bucket.s3UrlForObject(),
//   	"importPath": bucket.s3UrlForObject(),
//   	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
//   }
//
//   fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	Vpc: vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	StorageCapacityGiB: jsii.Number(1200),
//   	LustreConfiguration: LustreConfiguration,
//   })
//
type LustreFileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2, PERSISTENT_2 and PERSISTENT_1 deployment types using SSD storage type, the valid values are 1200 GiB, 2400 GiB, and increments of 2400 GiB.
	// For PERSISTENT_1 HDD file systems, valid values are increments of 6000 GiB for 12 MB/s/TiB file systems and increments of 1800 GiB for 40 MB/s/TiB file systems.
	StorageCapacityGiB *float64 `field:"required" json:"storageCapacityGiB" yaml:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	// Default: - no backup will be used.
	//
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	// Default: - the aws/fsx default KMS key for the AWS account being deployed into.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	// Default: - creates new security group which allows all outbound traffic.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The storage type for the file system that you're creating.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-storagetype
	//
	// Default: StorageType.SSD
	//
	StorageType StorageType `field:"optional" json:"storageType" yaml:"storageType"`
	// Additional configuration for FSx specific to Lustre.
	LustreConfiguration *LustreConfiguration `field:"required" json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// The subnet that the file system will be accessible from.
	VpcSubnet awsec2.ISubnet `field:"required" json:"vpcSubnet" yaml:"vpcSubnet"`
	// The Lustre version for the file system.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-filesystemtypeversion
	//
	// Default: - V_2_10, except for PERSISTENT_2 deployment type, where it is V_2_12 without metadata configuration mode and V_2_15 with metadata configuration mode.
	//
	FileSystemTypeVersion FileSystemTypeVersion `field:"optional" json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
}

