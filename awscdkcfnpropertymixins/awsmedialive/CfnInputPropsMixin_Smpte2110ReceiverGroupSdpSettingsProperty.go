package awsmedialive


// Information about the SDP files that describe the SMPTE 2110 streams that go into one SMPTE 2110 receiver group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   smpte2110ReceiverGroupSdpSettingsProperty := &Smpte2110ReceiverGroupSdpSettingsProperty{
//   	AncillarySdps: []interface{}{
//   		&InputSdpLocationProperty{
//   			MediaIndex: jsii.Number(123),
//   			SdpUrl: jsii.String("sdpUrl"),
//   		},
//   	},
//   	AudioSdps: []interface{}{
//   		&InputSdpLocationProperty{
//   			MediaIndex: jsii.Number(123),
//   			SdpUrl: jsii.String("sdpUrl"),
//   		},
//   	},
//   	VideoSdp: &InputSdpLocationProperty{
//   		MediaIndex: jsii.Number(123),
//   		SdpUrl: jsii.String("sdpUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html
//
type CfnInputPropsMixin_Smpte2110ReceiverGroupSdpSettingsProperty struct {
	// A list of InputSdpLocations.
	//
	// Each item in the list specifies the SDP file and index for one ancillary SMPTE 2110 stream.
	// Each stream encapsulates one captions stream (out of any number you can include) or the single SCTE 35 stream that you can include.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html#cfn-medialive-input-smpte2110receivergroupsdpsettings-ancillarysdps
	//
	AncillarySdps interface{} `field:"optional" json:"ancillarySdps" yaml:"ancillarySdps"`
	// A list of InputSdpLocations.
	//
	// Each item in the list specifies the SDP file and index for one audio SMPTE 2110 stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html#cfn-medialive-input-smpte2110receivergroupsdpsettings-audiosdps
	//
	AudioSdps interface{} `field:"optional" json:"audioSdps" yaml:"audioSdps"`
	// The location of the SDP file for one of the SMPTE 2110 streams in a receiver group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html#cfn-medialive-input-smpte2110receivergroupsdpsettings-videosdp
	//
	VideoSdp interface{} `field:"optional" json:"videoSdp" yaml:"videoSdp"`
}

