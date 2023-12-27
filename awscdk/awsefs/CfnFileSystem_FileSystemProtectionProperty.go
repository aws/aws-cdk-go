package awsefs


// Describes the protection on the file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemProtectionProperty := &FileSystemProtectionProperty{
//   	ReplicationOverwriteProtection: jsii.String("replicationOverwriteProtection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemprotection.html
//
type CfnFileSystem_FileSystemProtectionProperty struct {
	// The status of the file system's replication overwrite protection.
	//
	// - `ENABLED` – The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is `ENABLED` by default.
	// - `DISABLED` – The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.
	// - `REPLICATING` – The file system is being used as the destination file system in a replication configuration. The file system is read-only and is only modified only by EFS replication.
	//
	// If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemprotection.html#cfn-efs-filesystem-filesystemprotection-replicationoverwriteprotection
	//
	ReplicationOverwriteProtection *string `field:"optional" json:"replicationOverwriteProtection" yaml:"replicationOverwriteProtection"`
}

