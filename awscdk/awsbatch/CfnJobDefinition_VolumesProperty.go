package awsbatch


// A list of volumes that are associated with the job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumesProperty := &VolumesProperty{
//   	EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//
//   		// the properties below are optional
//   		AuthorizationConfig: &AuthorizationConfigProperty{
//   			AccessPointId: jsii.String("accessPointId"),
//   			Iam: jsii.String("iam"),
//   		},
//   		RootDirectory: jsii.String("rootDirectory"),
//   		TransitEncryption: jsii.String("transitEncryption"),
//   		TransitEncryptionPort: jsii.Number(123),
//   	},
//   	Host: &VolumesHostProperty{
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html
//
type CfnJobDefinition_VolumesProperty struct {
	// This is used when you're using an Amazon Elastic File System file system for job storage.
	//
	// For more information, see [Amazon EFS Volumes](https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-efsvolumeconfiguration
	//
	EfsVolumeConfiguration interface{} `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// The contents of the `host` parameter determine whether your data volume persists on the host container instance and where it's stored.
	//
	// If the host parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-host
	//
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// It can be up to 255 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_). This name is referenced in the `sourceVolume` parameter of container definition `mountPoints` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

