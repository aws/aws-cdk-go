package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFlowMediaStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowMediaStreamProps := &CfnFlowMediaStreamProps{
//   	FlowArn: jsii.String("flowArn"),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VideoFormat: jsii.String("videoFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html
//
type CfnFlowMediaStreamProps struct {
	// The Amazon Resource Name (ARN) of the flow that the media stream belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-flowarn
	//
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// A unique identifier for the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-mediastreamid
	//
	MediaStreamId *float64 `field:"required" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-mediastreamname
	//
	MediaStreamName *string `field:"required" json:"mediaStreamName" yaml:"mediaStreamName"`
	// The type of media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-mediastreamtype
	//
	MediaStreamType *string `field:"required" json:"mediaStreamType" yaml:"mediaStreamType"`
	// Attributes that are related to the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The sample rate (in Hz) for the stream.
	//
	// If the media stream type is video or ancillary data, set this value to 90000. If the media stream type is audio, set this value to either 48000 or 96000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-clockrate
	//
	ClockRate *float64 `field:"optional" json:"clockRate" yaml:"clockRate"`
	// A description that can help you quickly identify what your media stream is used for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The key-value pairs that can be used to tag and organize the media stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The resolution of the video.
	//
	// Required for a video media stream and rejected for other media stream types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowmediastream.html#cfn-mediaconnect-flowmediastream-videoformat
	//
	VideoFormat *string `field:"optional" json:"videoFormat" yaml:"videoFormat"`
}

