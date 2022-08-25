package awsecs


// A data volume used in a task definition.
//
// For tasks that use a Docker volume, specify a DockerVolumeConfiguration.
// For tasks that use a bind mount host volume, specify a host and optional sourcePath.
//
// For more information, see [Using Data Volumes in Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html).
//
// Example:
//   fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	memoryLimitMiB: jsii.Number(512),
//   	cpu: jsii.Number(256),
//   })
//   volume := map[string]interface{}{
//   	// Use an Elastic FileSystem
//   	"name": jsii.String("mydatavolume"),
//   	"efsVolumeConfiguration": map[string]*string{
//   		"fileSystemId": jsii.String("EFS"),
//   	},
//   }
//
//   container := fargateTaskDefinition.addVolume(volume)
//
type Volume struct {
	// The name of the volume.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.
	// This name is referenced in the sourceVolume parameter of container definition mountPoints.
	Name *string `field:"required" json:"name" yaml:"name"`
	// This property is specified when you are using Docker volumes.
	//
	// Docker volumes are only supported when you are using the EC2 launch type.
	// Windows containers only support the use of the local driver.
	// To use bind mounts, specify a host instead.
	DockerVolumeConfiguration *DockerVolumeConfiguration `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// This property is specified when you are using Amazon EFS.
	//
	// When specifying Amazon EFS volumes in tasks using the Fargate launch type,
	// Fargate creates a supervisor container that is responsible for managing the Amazon EFS volume.
	// The supervisor container uses a small amount of the task's memory.
	// The supervisor container is visible when querying the task metadata version 4 endpoint,
	// but is not visible in CloudWatch Container Insights.
	EfsVolumeConfiguration *EfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// This property is specified when you are using bind mount host volumes.
	//
	// Bind mount host volumes are supported when you are using either the EC2 or Fargate launch types.
	// The contents of the host parameter determine whether your bind mount host volume persists on the
	// host container instance and where it is stored. If the host parameter is empty, then the Docker
	// daemon assigns a host path for your data volume. However, the data is not guaranteed to persist
	// after the containers associated with it stop running.
	Host *Host `field:"optional" json:"host" yaml:"host"`
}

