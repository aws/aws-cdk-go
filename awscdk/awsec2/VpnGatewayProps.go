package awsec2


// The VpnGateway Properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpnGatewayProps := &VpnGatewayProps{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AmazonSideAsn: jsii.Number(123),
//   }
//
type VpnGatewayProps struct {
	// Default type ipsec.1.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Explicitly specify an Asn or let aws pick an Asn for you.
	// Default: 65000.
	//
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
}

