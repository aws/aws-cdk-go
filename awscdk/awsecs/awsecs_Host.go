package awsecs


// The details on a container instance bind mount host volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   host := &host{
//   	sourcePath: jsii.String("sourcePath"),
//   }
//
type Host struct {
	// Specifies the path on the host container instance that is presented to the container.
	//
	// If the sourcePath value does not exist on the host container instance, the Docker daemon creates it.
	// If the location does exist, the contents of the source path folder are exported.
	//
	// This property is not supported for tasks that use the Fargate launch type.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

