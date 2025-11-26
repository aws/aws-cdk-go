package previewawsefsmixins


// Describes the destination file system in the replication configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationDestinationProperty := &ReplicationDestinationProperty{
//   	AvailabilityZoneName: jsii.String("availabilityZoneName"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Region: jsii.String("region"),
//   	RoleArn: jsii.String("roleArn"),
//   	Status: jsii.String("status"),
//   	StatusMessage: jsii.String("statusMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html
//
type CfnFileSystemPropsMixin_ReplicationDestinationProperty struct {
	// For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located.
	//
	// Use the format `us-east-1a` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide* .
	//
	// > One Zone file system type is not available in all Availability Zones in AWS Regions where Amazon EFS is available.
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
	// > For One Zone file systems, the replication configuration must specify the AWS Region in which the destination file system is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The Amazon Resource Name (ARN) of the current source file system in the replication configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Describes the status of the replication configuration.
	//
	// For more information about replication status, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Message that provides details about the `PAUSED` or `ERRROR` state of the replication destination configuration.
	//
	// For more information about replication status messages, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-replicationdestination.html#cfn-efs-filesystem-replicationdestination-statusmessage
	//
	StatusMessage *string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}

