package awsec2


// Describes the Connect attachment options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayConnectOptionsProperty := &TransitGatewayConnectOptionsProperty{
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnect-transitgatewayconnectoptions.html
//
type CfnTransitGatewayConnect_TransitGatewayConnectOptionsProperty struct {
	// The tunnel protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayconnect-transitgatewayconnectoptions.html#cfn-ec2-transitgatewayconnect-transitgatewayconnectoptions-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

