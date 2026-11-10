package awsmediaconnect


// The configuration settings for a destination endpoint that uses the RIST protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ristStandardDestinationEndpointConfigurationProperty := &RistStandardDestinationEndpointConfigurationProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationIngestPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-riststandarddestinationendpointconfiguration.html
//
type CfnDestinationEndpoint_RistStandardDestinationEndpointConfigurationProperty struct {
	// The IP address on which the destination endpoint receives content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-riststandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-riststandarddestinationendpointconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// The port on which the destination endpoint receives content.
	//
	// The valid range is 1024 to 65531.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-riststandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-riststandarddestinationendpointconfiguration-destinationingestport
	//
	DestinationIngestPort *float64 `field:"required" json:"destinationIngestPort" yaml:"destinationIngestPort"`
}

