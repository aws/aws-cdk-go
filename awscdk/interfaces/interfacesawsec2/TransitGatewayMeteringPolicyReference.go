package interfacesawsec2


// A reference to a TransitGatewayMeteringPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayMeteringPolicyReference := &TransitGatewayMeteringPolicyReference{
//   	TransitGatewayMeteringPolicyId: jsii.String("transitGatewayMeteringPolicyId"),
//   }
//
type TransitGatewayMeteringPolicyReference struct {
	// The TransitGatewayMeteringPolicyId of the TransitGatewayMeteringPolicy resource.
	TransitGatewayMeteringPolicyId *string `field:"required" json:"transitGatewayMeteringPolicyId" yaml:"transitGatewayMeteringPolicyId"`
}

