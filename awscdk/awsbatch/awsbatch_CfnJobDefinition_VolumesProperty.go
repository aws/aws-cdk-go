package awsbatch


// A list of volumes associated with the job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumesProperty := &volumesProperty{
//   	efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   		fileSystemId: jsii.String("fileSystemId"),
//
//   		// the properties below are optional
//   		authorizationConfig: &authorizationConfigProperty{
//   			accessPointId: jsii.String("accessPointId"),
//   			iam: jsii.String("iam"),
//   		},
//   		rootDirectory: jsii.String("rootDirectory"),
//   		transitEncryption: jsii.String("transitEncryption"),
//   		transitEncryptionPort: jsii.Number(123),
//   	},
//   	host: &volumesHostProperty{
//   		sourcePath: jsii.String("sourcePath"),
//   	},
//   	name: jsii.String("name"),
//   }
//
type CfnJobDefinition_VolumesProperty struct {
	// This is used when you're using an Amazon Elastic File System file system for job storage.
	//
	// For more information, see [Amazon EFS Volumes](https://docs.aws.amazon.com/batch/latest/userguide/efs-volumes.html) in the *AWS Batch User Guide* .
	EfsVolumeConfiguration interface{} `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// The contents of the `host` parameter determine whether your data volume persists on the host container instance and where it is stored.
	//
	// If the host parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers associated with it stop running.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// It can be up to 255 letters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_). This name is referenced in the `sourceVolume` parameter of container definition `mountPoints` .
	Name *string `field:"optional" json:"name" yaml:"name"`
}

