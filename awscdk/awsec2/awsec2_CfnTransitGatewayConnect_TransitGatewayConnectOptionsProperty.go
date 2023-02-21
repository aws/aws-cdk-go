package awsec2


// Describes the Connect attachment options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayConnectOptionsProperty := &transitGatewayConnectOptionsProperty{
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnTransitGatewayConnect_TransitGatewayConnectOptionsProperty struct {
	// The tunnel protocol.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

