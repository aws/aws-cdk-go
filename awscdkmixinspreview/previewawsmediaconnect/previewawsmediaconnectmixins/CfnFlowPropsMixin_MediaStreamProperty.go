package previewawsmediaconnectmixins


// A media stream represents one component of your content, such as video, audio, or ancillary data.
//
// After you add a media stream to your flow, you can associate it with sources and outputs that use the ST 2110 JPEG XS or CDI protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaStreamProperty := &MediaStreamProperty{
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
//   	MediaStreamId: jsii.Number(123),
//   	MediaStreamName: jsii.String("mediaStreamName"),
//   	MediaStreamType: jsii.String("mediaStreamType"),
//   	VideoFormat: jsii.String("videoFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html
//
type CfnFlowPropsMixin_MediaStreamProperty struct {
	// Attributes that are related to the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The sample rate for the stream.
	//
	// This value is measured in Hz.
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
	// A unique identifier for the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-mediastreamid
	//
	MediaStreamId *float64 `field:"optional" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-mediastreamname
	//
	MediaStreamName *string `field:"optional" json:"mediaStreamName" yaml:"mediaStreamName"`
	// The type of media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-mediastreamtype
	//
	MediaStreamType *string `field:"optional" json:"mediaStreamType" yaml:"mediaStreamType"`
	// The resolution of the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-mediastream.html#cfn-mediaconnect-flow-mediastream-videoformat
	//
	VideoFormat *string `field:"optional" json:"videoFormat" yaml:"videoFormat"`
}

