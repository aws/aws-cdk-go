package interfacesawsmediastore


// A reference to a Container resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerReference := &ContainerReference{
//   	ContainerId: jsii.String("containerId"),
//   	ContainerName: jsii.String("containerName"),
//   }
//
type ContainerReference struct {
	// The Id of the Container resource.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// The ContainerName of the Container resource.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
}

