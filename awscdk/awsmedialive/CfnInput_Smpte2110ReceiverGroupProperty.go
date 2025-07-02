package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smpte2110ReceiverGroupProperty := &Smpte2110ReceiverGroupProperty{
//   	SdpSettings: &Smpte2110ReceiverGroupSdpSettingsProperty{
//   		AncillarySdps: []interface{}{
//   			&InputSdpLocationProperty{
//   				MediaIndex: jsii.Number(123),
//   				SdpUrl: jsii.String("sdpUrl"),
//   			},
//   		},
//   		AudioSdps: []interface{}{
//   			&InputSdpLocationProperty{
//   				MediaIndex: jsii.Number(123),
//   				SdpUrl: jsii.String("sdpUrl"),
//   			},
//   		},
//   		VideoSdp: &InputSdpLocationProperty{
//   			MediaIndex: jsii.Number(123),
//   			SdpUrl: jsii.String("sdpUrl"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroup.html
//
type CfnInput_Smpte2110ReceiverGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroup.html#cfn-medialive-input-smpte2110receivergroup-sdpsettings
	//
	SdpSettings interface{} `field:"optional" json:"sdpSettings" yaml:"sdpSettings"`
}

