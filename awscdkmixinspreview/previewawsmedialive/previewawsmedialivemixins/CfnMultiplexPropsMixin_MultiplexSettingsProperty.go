package previewawsmedialivemixins


// Contains configuration for a Multiplex event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexSettingsProperty := &MultiplexSettingsProperty{
//   	MaximumVideoBufferDelayMilliseconds: jsii.Number(123),
//   	TransportStreamBitrate: jsii.Number(123),
//   	TransportStreamId: jsii.Number(123),
//   	TransportStreamReservedBitrate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html
//
type CfnMultiplexPropsMixin_MultiplexSettingsProperty struct {
	// Maximum video buffer delay in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-maximumvideobufferdelaymilliseconds
	//
	MaximumVideoBufferDelayMilliseconds *float64 `field:"optional" json:"maximumVideoBufferDelayMilliseconds" yaml:"maximumVideoBufferDelayMilliseconds"`
	// Transport stream bit rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-transportstreambitrate
	//
	TransportStreamBitrate *float64 `field:"optional" json:"transportStreamBitrate" yaml:"transportStreamBitrate"`
	// Transport stream ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-transportstreamid
	//
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// Transport stream reserved bit rate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplex-multiplexsettings.html#cfn-medialive-multiplex-multiplexsettings-transportstreamreservedbitrate
	//
	TransportStreamReservedBitrate *float64 `field:"optional" json:"transportStreamReservedBitrate" yaml:"transportStreamReservedBitrate"`
}

