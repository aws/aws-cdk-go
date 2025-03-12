package awsmediaconnect


// A single track or stream of media that contains video, audio, or ancillary data.
//
// After you add a media stream to a flow, you can associate it with sources and outputs on that flow, as long as they use the CDI protocol or the ST 2110 JPEG XS protocol. Each source or output can consist of one or many media streams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mediaStreamProperty := &MediaStreamProperty{
//   	MediaStreamId: jsii.Number(123),
//   	MediaStreamName: jsii.String("mediaStreamName"),
//   	MediaStreamType: jsii.String("mediaStreamType"),
//
//   	// the properties below are optional
//   	Attributes: &MediaStreamAttributesProperty{
//   		Fmtp: &FmtpProperty{
//   			ChannelOrder: jsii.String("channelOrder"),
//   			Colorimetry: jsii.String("colorimetry"),
//   			ExactFramerate: jsii.String("exactFramerate"),
//   			Par: jsii.String("par"),
//   			Range: jsii.String("range"),
//   			ScanMode: jsii.String("scanMode"),
//   			Tcs: jsii.String("tcs"),
//   		},
//   		Lang: jsii.String("lang"),
//   	},
//   	ClockRate: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	Fmt: jsii.Number(123),
//   	VideoFormat: jsii.String("videoFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html
//
type CfnFlow_MediaStreamProperty struct {
	// A unique identifier for the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-mediastreamid
	//
	MediaStreamId *float64 `field:"required" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-mediastreamname
	//
	MediaStreamName *string `field:"required" json:"mediaStreamName" yaml:"mediaStreamName"`
	// The type of media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-mediastreamtype
	//
	MediaStreamType *string `field:"required" json:"mediaStreamType" yaml:"mediaStreamType"`
	// Attributes that are related to the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The sample rate for the stream.
	//
	// This value in measured in kHz.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-clockrate
	//
	ClockRate *float64 `field:"optional" json:"clockRate" yaml:"clockRate"`
	// A description that can help you quickly identify what your media stream is used for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The format type number (sometimes referred to as RTP payload type) of the media stream.
	//
	// MediaConnect assigns this value to the media stream. For ST 2110 JPEG XS outputs, you need to provide this value to the receiver.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-fmt
	//
	Fmt *float64 `field:"optional" json:"fmt" yaml:"fmt"`
	// The resolution of the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-videoformat
	//
	VideoFormat *string `field:"optional" json:"videoFormat" yaml:"videoFormat"`
}

