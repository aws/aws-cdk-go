package interfacesawsroute53recoverycontrol


// A reference to a RoutingControl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingControlReference := &RoutingControlReference{
//   	RoutingControlArn: jsii.String("routingControlArn"),
//   }
//
type RoutingControlReference struct {
	// The RoutingControlArn of the RoutingControl resource.
	RoutingControlArn *string `field:"required" json:"routingControlArn" yaml:"routingControlArn"`
}

