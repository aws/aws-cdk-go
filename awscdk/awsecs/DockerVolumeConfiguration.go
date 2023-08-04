package awsecs


// The configuration for a Docker volume.
//
// Docker volumes are only supported when you are using the EC2 launch type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerVolumeConfiguration := &DockerVolumeConfiguration{
//   	Driver: jsii.String("driver"),
//   	Scope: awscdk.Aws_ecs.Scope_TASK,
//
//   	// the properties below are optional
//   	Autoprovision: jsii.Boolean(false),
//   	DriverOpts: map[string]*string{
//   		"driverOptsKey": jsii.String("driverOpts"),
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   }
//
type DockerVolumeConfiguration struct {
	// The Docker volume driver to use.
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// The scope for the Docker volume that determines its lifecycle.
	Scope Scope `field:"required" json:"scope" yaml:"scope"`
	// Specifies whether the Docker volume should be created if it does not already exist.
	//
	// If true is specified, the Docker volume will be created for you.
	// Default: false.
	//
	Autoprovision *bool `field:"optional" json:"autoprovision" yaml:"autoprovision"`
	// A map of Docker driver-specific options passed through.
	// Default: No options.
	//
	DriverOpts *map[string]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Custom metadata to add to your Docker volume.
	// Default: No labels.
	//
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

