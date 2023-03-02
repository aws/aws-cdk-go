package awsecs


// The `HostVolumeProperties` property specifies details on a container instance bind mount host volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostVolumePropertiesProperty := &hostVolumePropertiesProperty{
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type CfnTaskDefinition_HostVolumePropertiesProperty struct {
	// When the `host` parameter is used, specify a `sourcePath` to declare the path on the host container instance that's presented to the container.
	//
	// If this parameter is empty, then the Docker daemon has assigned a host path for you. If the `host` parameter contains a `sourcePath` file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the `sourcePath` value doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.
	//
	// If you're using the Fargate launch type, the `sourcePath` parameter is not supported.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

