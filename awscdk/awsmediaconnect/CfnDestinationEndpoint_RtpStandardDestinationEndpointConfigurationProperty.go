package awsmediaconnect


// The configuration settings for a destination endpoint that uses the RTP protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtpStandardDestinationEndpointConfigurationProperty := &RtpStandardDestinationEndpointConfigurationProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationIngestPort: jsii.Number(123),
//   	ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration.html
//
type CfnDestinationEndpoint_RtpStandardDestinationEndpointConfigurationProperty struct {
	// The IP address on which the destination endpoint receives content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration-destinationaddress
	//
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// The port on which the destination endpoint receives content.
	//
	// The valid range is 1024 to 65531.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration-destinationingestport
	//
	DestinationIngestPort *float64 `field:"required" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-rtpstandarddestinationendpointconfiguration-forwarderrorcorrectionstate
	//
	ForwardErrorCorrectionState *string `field:"required" json:"forwardErrorCorrectionState" yaml:"forwardErrorCorrectionState"`
}

