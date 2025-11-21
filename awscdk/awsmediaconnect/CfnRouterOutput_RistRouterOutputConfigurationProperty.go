package awsmediaconnect


// The configuration settings for a router output using the RIST (Reliable Internet Stream Transport) protocol, including the destination address and port.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ristRouterOutputConfigurationProperty := &RistRouterOutputConfigurationProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-ristrouteroutputconfiguration.html
//
type CfnRouterOutput_RistRouterOutputConfigurationProperty struct {
	// The destination IP address for the RIST protocol in the router output configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-ristrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-ristrouteroutputconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// The destination port number for the RIST protocol in the router output configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-ristrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-ristrouteroutputconfiguration-destinationport
	//
	DestinationPort *float64 `field:"required" json:"destinationPort" yaml:"destinationPort"`
}

