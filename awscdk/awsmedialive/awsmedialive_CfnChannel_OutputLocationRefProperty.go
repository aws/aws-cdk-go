package awsmedialive


// A reference to an OutputDestination ID that is defined in the channel.
//
// This entity is used by ArchiveGroupSettings, FrameCaptureGroupSettings, HlsGroupSettings, MediaPackageGroupSettings, MSSmoothGroupSettings, RtmpOutputSettings, and UdpOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputLocationRefProperty := &outputLocationRefProperty{
//   	destinationRefId: jsii.String("destinationRefId"),
//   }
//
type CfnChannel_OutputLocationRefProperty struct {
	// A reference ID for this destination.
	DestinationRefId *string `field:"optional" json:"destinationRefId" yaml:"destinationRefId"`
}

