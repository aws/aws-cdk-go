package awsefs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemprotection.html#cfn-efs-filesystem-filesystemprotection-replicationoverwriteprotection
	//
	ReplicationOverwriteProtection *string `field:"optional" json:"replicationOverwriteProtection" yaml:"replicationOverwriteProtection"`
}

