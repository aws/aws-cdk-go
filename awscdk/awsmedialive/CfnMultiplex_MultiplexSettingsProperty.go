package awsmedialive


// Contains configuration for a Multiplex event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexSettingsProperty := &MultiplexSettingsProperty{
//   	TransportStreamBitrate: jsii.Number(123),
//   	TransportStreamId: jsii.Number(123),
//
//   	// the properties below are optional
//   	MaximumVideoBufferDelayMilliseconds: jsii.Number(123),
//   	TransportStreamReservedBitrate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html
//
type CfnMultiplex_MultiplexSettingsProperty struct {
	// Transport stream bit rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-transportstreambitrate
	//
	TransportStreamBitrate *float64 `field:"required" json:"transportStreamBitrate" yaml:"transportStreamBitrate"`
	// Transport stream ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-transportstreamid
	//
	TransportStreamId *float64 `field:"required" json:"transportStreamId" yaml:"transportStreamId"`
	// Maximum video buffer delay in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-maximumvideobufferdelaymilliseconds
	//
	MaximumVideoBufferDelayMilliseconds *float64 `field:"optional" json:"maximumVideoBufferDelayMilliseconds" yaml:"maximumVideoBufferDelayMilliseconds"`
	// Transport stream reserved bit rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-transportstreamreservedbitrate
	//
	TransportStreamReservedBitrate *float64 `field:"optional" json:"transportStreamReservedBitrate" yaml:"transportStreamReservedBitrate"`
}

