package awsmediaconnect


// Attributes that are related to the media stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mediaStreamAttributesProperty := &MediaStreamAttributesProperty{
//   	Fmtp: &FmtpProperty{
//   		ChannelOrder: jsii.String("channelOrder"),
//   		Colorimetry: jsii.String("colorimetry"),
//   		ExactFramerate: jsii.String("exactFramerate"),
//   		Par: jsii.String("par"),
//   		Range: jsii.String("range"),
//   		ScanMode: jsii.String("scanMode"),
//   		Tcs: jsii.String("tcs"),
//   	},
//   	Lang: jsii.String("lang"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowmediastream-mediastreamattributes.html
//
type CfnFlowMediaStreamPropsMixin_MediaStreamAttributesProperty struct {
	// A set of parameters that define the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowmediastream-mediastreamattributes.html#cfn-mediaconnect-flowmediastream-mediastreamattributes-fmtp
	//
	Fmtp interface{} `field:"optional" json:"fmtp" yaml:"fmtp"`
	// The audio language, in a format that is recognized by the receiver.
	//
	// Can only be specified for an audio media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowmediastream-mediastreamattributes.html#cfn-mediaconnect-flowmediastream-mediastreamattributes-lang
	//
	Lang *string `field:"optional" json:"lang" yaml:"lang"`
}

