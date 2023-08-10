package awsecs


// The `Volume` property specifies a data volume used in a task definition.
//
// For tasks that use a Docker volume, specify a `DockerVolumeConfiguration` . For tasks that use a bind mount host volume, specify a `host` and optional `sourcePath` . For more information about `host` and optional `sourcePath` , see [Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#volumes) and [Using Data Volumes in Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeProperty := &VolumeProperty{
//   	DockerVolumeConfiguration: &DockerVolumeConfigurationProperty{
//   		Autoprovision: jsii.Boolean(false),
//   		Driver: jsii.String("driver"),
//   		DriverOpts: map[string]*string{
//   			"driverOptsKey": jsii.String("driverOpts"),
//   		},
//   		Labels: map[string]*string{
//   			"labelsKey": jsii.String("labels"),
//   		},
//   		Scope: jsii.String("scope"),
//   	},
//   	EfsVolumeConfiguration: &EFSVolumeConfigurationProperty{
//   		FilesystemId: jsii.String("filesystemId"),
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
//   	Host: &HostVolumePropertiesProperty{
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volume.html
//
type CfnTaskDefinition_VolumeProperty struct {
	// This parameter is specified when you use Docker volumes.
	//
	// Windows containers only support the use of the `local` driver. To use bind mounts, specify the `host` parameter instead.
	//
	// > Docker volumes aren't supported by tasks run on AWS Fargate .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volume.html#cfn-ecs-taskdefinition-volume-dockervolumeconfiguration
	//
	DockerVolumeConfiguration interface{} `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// This parameter is specified when you use an Amazon Elastic File System file system for task storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volume.html#cfn-ecs-taskdefinition-volume-efsvolumeconfiguration
	//
	EfsVolumeConfiguration interface{} `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// This parameter is specified when you use bind mount host volumes.
	//
	// The contents of the `host` parameter determine whether your bind mount host volume persists on the host container instance and where it's stored. If the `host` parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
	//
	// Windows containers can mount whole directories on the same drive as `$env:ProgramData` . Windows containers can't mount directories on a different drive, and mount point can't be across drives. For example, you can mount `C:\my\path:C:\my\path` and `D:\:D:\` , but not `D:\my\path:C:\my\path` or `D:\:C:\my\path` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volume.html#cfn-ecs-taskdefinition-volume-host
	//
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. This name is referenced in the `sourceVolume` parameter of container definition `mountPoints` .
	//
	// This is required wwhen you use an Amazon EFS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-volume.html#cfn-ecs-taskdefinition-volume-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

