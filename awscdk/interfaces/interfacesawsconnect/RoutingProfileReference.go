package interfacesawsconnect


// A reference to a RoutingProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingProfileReference := &RoutingProfileReference{
//   	RoutingProfileArn: jsii.String("routingProfileArn"),
//   }
//
type RoutingProfileReference struct {
	// The RoutingProfileArn of the RoutingProfile resource.
	RoutingProfileArn *string `field:"required" json:"routingProfileArn" yaml:"routingProfileArn"`
}

