package awsecs


// This parameter is specified when you're using an Amazon S3 Files file system for task storage.
//
// For more information, see [Amazon S3 Files volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/s3files-volumes.html) in the *Amazon Elastic Container Service Developer Guide*.
//   Your task definition must include a Task IAM Role. See [IAM role for attaching your file system to compute resources](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-files-prereq-policies.html#s3-files-prereq-iam-compute-role) for required permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FilesVolumeConfigurationProperty := &S3FilesVolumeConfigurationProperty{
//   	FileSystemArn: jsii.String("fileSystemArn"),
//
//   	// the properties below are optional
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	RootDirectory: jsii.String("rootDirectory"),
//   	TransitEncryptionPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html
//
type CfnTaskDefinition_S3FilesVolumeConfigurationProperty struct {
	// The full ARN of the S3 Files file system to mount.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-filesystemarn
	//
	FileSystemArn *string `field:"required" json:"fileSystemArn" yaml:"fileSystemArn"`
	// The full ARN of the S3 Files access point to use.
	//
	// If an access point is specified, the root directory value specified in the ``S3FilesVolumeConfiguration`` must either be omitted or set to ``/`` which will enforce the path set on the S3 Files access point. For more information, see [Creating S3 Files access points](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-files-access-points-creating.html).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The directory within the Amazon S3 Files file system to mount as the root directory.
	//
	// If this parameter is omitted, the root of the Amazon S3 Files file system will be used. Specifying ``/`` will have the same effect as omitting this parameter.
	//   If a S3 Files access point is specified in the ``accessPointArn``, the root directory parameter must either be omitted or set to ``/`` which will enforce the path set on the S3 Files access point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-rootdirectory
	//
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// The port to use for sending encrypted data between the ECS host and the S3 Files file system.
	//
	// If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon S3 Files mount helper uses. For more information, see [S3 Files mount helper](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-files-mounting.html).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-transitencryptionport
	//
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

