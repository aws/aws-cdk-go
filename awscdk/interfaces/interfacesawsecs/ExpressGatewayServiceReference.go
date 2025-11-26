package interfacesawsecs


// A reference to a ExpressGatewayService resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressGatewayServiceReference := &ExpressGatewayServiceReference{
//   	ServiceArn: jsii.String("serviceArn"),
//   }
//
type ExpressGatewayServiceReference struct {
	// The ServiceArn of the ExpressGatewayService resource.
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
}

