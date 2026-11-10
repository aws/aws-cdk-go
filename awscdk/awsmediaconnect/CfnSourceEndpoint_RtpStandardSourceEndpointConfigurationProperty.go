package awsmediaconnect


// The configuration settings for a source endpoint that uses the RTP protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtpStandardSourceEndpointConfigurationProperty := &RtpStandardSourceEndpointConfigurationProperty{
//   	DestinationIngestPort: jsii.Number(123),
//   	ForwardErrorCorrectionState: jsii.String("forwardErrorCorrectionState"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-rtpstandardsourceendpointconfiguration.html
//
type CfnSourceEndpoint_RtpStandardSourceEndpointConfigurationProperty struct {
	// The port on the destination that receives the content.
	//
	// The valid range is 3000 to 30000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-rtpstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-rtpstandardsourceendpointconfiguration-destinationingestport
	//
	DestinationIngestPort *float64 `field:"required" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-rtpstandardsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-rtpstandardsourceendpointconfiguration-forwarderrorcorrectionstate
	//
	ForwardErrorCorrectionState *string `field:"required" json:"forwardErrorCorrectionState" yaml:"forwardErrorCorrectionState"`
}

