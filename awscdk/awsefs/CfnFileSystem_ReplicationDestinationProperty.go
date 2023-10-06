package awsefs


// Describes the destination file system in the replication configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationDestinationProperty := &ReplicationDestinationProperty{
//   	AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html
//
type CfnFileSystem_ReplicationDestinationProperty struct {
	// The AWS Availability Zone in which to create the file system.
	//
	// > For file systems using One Zone storage classes, the replication configuration must specify the Availability Zone in which the destination file system is located.
	//
	// Use the format `us-east-1a` to specify the Availability Zone. For more information about One Zone storage classes, see [Using EFS storage classes](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide* .
	//
	// > One Zone storage classes are not available in all Availability Zones in AWS Regions where Amazon EFS is available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-availabilityzonename
	//
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// The ID of the destination Amazon EFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The ID of an AWS KMS key used to protect the encrypted file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The AWS Region in which the destination file system is located.
	//
	// > For file systems using Standard storage classes, the replication configuration must specify the AWS Region in which the destination file system is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

