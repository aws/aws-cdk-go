package interfacesawsec2


// A reference to a RouteServerAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeServerAssociationReference := &RouteServerAssociationReference{
//   	RouteServerId: jsii.String("routeServerId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type RouteServerAssociationReference struct {
	// The RouteServerId of the RouteServerAssociation resource.
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// The VpcId of the RouteServerAssociation resource.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

