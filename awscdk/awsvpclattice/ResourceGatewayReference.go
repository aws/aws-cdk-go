package awsvpclattice


// A reference to a ResourceGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceGatewayReference := &ResourceGatewayReference{
//   	ResourceGatewayArn: jsii.String("resourceGatewayArn"),
//   }
//
type ResourceGatewayReference struct {
	// The Arn of the ResourceGateway resource.
	ResourceGatewayArn *string `field:"required" json:"resourceGatewayArn" yaml:"resourceGatewayArn"`
}

