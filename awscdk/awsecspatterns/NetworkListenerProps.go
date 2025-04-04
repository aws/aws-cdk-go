package awsecspatterns


// Properties to define an network listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkListenerProps := &NetworkListenerProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Port: jsii.Number(123),
//   }
//
type NetworkListenerProps struct {
	// Name of the listener.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port on which the listener listens for requests.
	// Default: 80.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

