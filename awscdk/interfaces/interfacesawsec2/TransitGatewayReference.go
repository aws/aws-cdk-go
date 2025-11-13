package interfacesawsec2


// A reference to a TransitGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayReference := &TransitGatewayReference{
//   	TransitGatewayArn: jsii.String("transitGatewayArn"),
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   }
//
type TransitGatewayReference struct {
	// The ARN of the TransitGateway resource.
	TransitGatewayArn *string `field:"required" json:"transitGatewayArn" yaml:"transitGatewayArn"`
	// The Id of the TransitGateway resource.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
}

