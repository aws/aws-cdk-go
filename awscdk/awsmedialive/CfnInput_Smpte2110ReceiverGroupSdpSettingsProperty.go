package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnInput_Smpte2110ReceiverGroupSdpSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html#cfn-medialive-input-smpte2110receivergroupsdpsettings-ancillarysdps
	//
	AncillarySdps interface{} `field:"optional" json:"ancillarySdps" yaml:"ancillarySdps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html#cfn-medialive-input-smpte2110receivergroupsdpsettings-audiosdps
	//
	AudioSdps interface{} `field:"optional" json:"audioSdps" yaml:"audioSdps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsdpsettings.html#cfn-medialive-input-smpte2110receivergroupsdpsettings-videosdp
	//
	VideoSdp interface{} `field:"optional" json:"videoSdp" yaml:"videoSdp"`
}

