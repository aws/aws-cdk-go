package awsbatch


// A data volume that's used in a job's container properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeProperty := &VolumeProperty{
//   	EfsVolumeConfiguration: &EFSVolumeConfigurationProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//
//   		// the properties below are optional
//   		AuthorizationConfig: &EFSAuthorizationConfigProperty{
//   			AccessPointId: jsii.String("accessPointId"),
//   			Iam: jsii.String("iam"),
//   		},
//   		RootDirectory: jsii.String("rootDirectory"),
//   		TransitEncryption: jsii.String("transitEncryption"),
//   		TransitEncryptionPort: jsii.Number(123),
//   	},
//   	Host: &HostProperty{
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volume.html
//
type CfnJobDefinition_VolumeProperty struct {
	// This parameter is specified when you're using an Amazon Elastic File System file system for job storage.
	//
	// Jobs that are running on Fargate resources must specify a `platformVersion` of at least `1.4.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volume.html#cfn-batch-jobdefinition-volume-efsvolumeconfiguration
	//
	EfsVolumeConfiguration interface{} `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// The contents of the `host` parameter determine whether your data volume persists on the host container instance and where it's stored.
	//
	// If the host parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volume.html#cfn-batch-jobdefinition-volume-host
	//
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// It can be up to 255 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_). This name is referenced in the `sourceVolume` parameter of container definition `mountPoints` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volume.html#cfn-batch-jobdefinition-volume-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

