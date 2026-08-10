package interfacesawsec2


// A reference to a TransitGatewayPolicyTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayPolicyTableReference := &TransitGatewayPolicyTableReference{
//   	TransitGatewayPolicyTableId: jsii.String("transitGatewayPolicyTableId"),
//   }
//
type TransitGatewayPolicyTableReference struct {
	// The TransitGatewayPolicyTableId of the TransitGatewayPolicyTable resource.
	TransitGatewayPolicyTableId *string `field:"required" json:"transitGatewayPolicyTableId" yaml:"transitGatewayPolicyTableId"`
}

