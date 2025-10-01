package awslightsail


// A reference to a Container resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerReference := &ContainerReference{
//   	ContainerArn: jsii.String("containerArn"),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
type ContainerReference struct {
	// The ARN of the Container resource.
	ContainerArn *string `field:"required" json:"containerArn" yaml:"containerArn"`
	// The ServiceName of the Container resource.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

