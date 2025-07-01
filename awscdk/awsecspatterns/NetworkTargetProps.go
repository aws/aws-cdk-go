package awsecspatterns


// Properties to define a network load balancer target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkTargetProps := &NetworkTargetProps{
//   	ContainerPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	Listener: jsii.String("listener"),
//   }
//
type NetworkTargetProps struct {
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// Name of the listener the target group attached to.
	// Default: - default listener (first added listener).
	//
	Listener *string `field:"optional" json:"listener" yaml:"listener"`
}

