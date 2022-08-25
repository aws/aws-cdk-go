package awsecs


// The `DockerVolumeConfiguration` property specifies a Docker volume configuration and is used when you use Docker volumes.
//
// Docker volumes are only supported when you are using the EC2 launch type. Windows containers only support the use of the `local` driver. To use bind mounts, specify a `host` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerVolumeConfigurationProperty := &dockerVolumeConfigurationProperty{
//   	autoprovision: jsii.Boolean(false),
//   	driver: jsii.String("driver"),
//   	driverOpts: map[string]*string{
//   		"driverOptsKey": jsii.String("driverOpts"),
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	scope: jsii.String("scope"),
//   }
//
type CfnTaskDefinition_DockerVolumeConfigurationProperty struct {
	// If this value is `true` , the Docker volume is created if it doesn't already exist.
	//
	// > This field is only used if the `scope` is `shared` .
	Autoprovision interface{} `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// The Docker volume driver to use.
	//
	// The driver value must match the driver name provided by Docker because it is used for task placement. If the driver was installed using the Docker plugin CLI, use `docker plugin ls` to retrieve the driver name from your container instance. If the driver was installed using another method, use Docker plugin discovery to retrieve the driver name. For more information, see [Docker plugin discovery](https://docs.aws.amazon.com/https://docs.docker.com/engine/extend/plugin_api/#plugin-discovery) . This parameter maps to `Driver` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxdriver` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// A map of Docker driver-specific options passed through.
	//
	// This parameter maps to `DriverOpts` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxopt` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	DriverOpts interface{} `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	//
	// This parameter maps to `Labels` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the `xxlabel` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/) .
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The scope for the Docker volume that determines its lifecycle.
	//
	// Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as `shared` persist after the task stops.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

