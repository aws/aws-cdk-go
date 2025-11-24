package mixinsawsecs


// The `DockerVolumeConfiguration` property specifies a Docker volume configuration and is used when you use Docker volumes.
//
// Docker volumes are only supported when you are using the EC2 launch type. Windows containers only support the use of the `local` driver. To use bind mounts, specify a `host` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dockerVolumeConfigurationProperty := &DockerVolumeConfigurationProperty{
//   	Autoprovision: jsii.Boolean(false),
//   	Driver: jsii.String("driver"),
//   	DriverOpts: map[string]*string{
//   		"driverOptsKey": jsii.String("driverOpts"),
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html
//
type CfnTaskDefinitionPropsMixin_DockerVolumeConfigurationProperty struct {
	// If this value is `true` , the Docker volume is created if it doesn't already exist.
	//
	// > This field is only used if the `scope` is `shared` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-autoprovision
	//
	Autoprovision interface{} `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// The Docker volume driver to use.
	//
	// The driver value must match the driver name provided by Docker because it is used for task placement. If the driver was installed using the Docker plugin CLI, use `docker plugin ls` to retrieve the driver name from your container instance. If the driver was installed using another method, use Docker plugin discovery to retrieve the driver name. This parameter maps to `Driver` in the docker container create command and the `xxdriver` option to docker volume create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-driver
	//
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// A map of Docker driver-specific options passed through.
	//
	// This parameter maps to `DriverOpts` in the docker create-volume command and the `xxopt` option to docker volume create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-driveropts
	//
	DriverOpts interface{} `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	//
	// This parameter maps to `Labels` in the docker container create command and the `xxlabel` option to docker volume create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The scope for the Docker volume that determines its lifecycle.
	//
	// Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as `shared` persist after the task stops.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html#cfn-ecs-taskdefinition-dockervolumeconfiguration-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

