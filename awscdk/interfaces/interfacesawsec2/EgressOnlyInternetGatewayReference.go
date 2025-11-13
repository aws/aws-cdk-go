package interfacesawsec2


// A reference to a EgressOnlyInternetGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressOnlyInternetGatewayReference := &EgressOnlyInternetGatewayReference{
//   	EgressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   }
//
type EgressOnlyInternetGatewayReference struct {
	// The Id of the EgressOnlyInternetGateway resource.
	EgressOnlyInternetGatewayId *string `field:"required" json:"egressOnlyInternetGatewayId" yaml:"egressOnlyInternetGatewayId"`
}

