package interfacesawsrtbfabric


// A reference to a ResponderGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   responderGatewayReference := &ResponderGatewayReference{
//   	ResponderGatewayArn: jsii.String("responderGatewayArn"),
//   }
//
type ResponderGatewayReference struct {
	// The Arn of the ResponderGateway resource.
	ResponderGatewayArn *string `field:"required" json:"responderGatewayArn" yaml:"responderGatewayArn"`
}

