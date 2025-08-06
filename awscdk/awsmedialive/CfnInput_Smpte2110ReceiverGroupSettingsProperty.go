package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   smpte2110ReceiverGroupSettingsProperty := &Smpte2110ReceiverGroupSettingsProperty{
//   	Smpte2110ReceiverGroups: []interface{}{
//   		&Smpte2110ReceiverGroupProperty{
//   			SdpSettings: &Smpte2110ReceiverGroupSdpSettingsProperty{
//   				AncillarySdps: []interface{}{
//   					&InputSdpLocationProperty{
//   						MediaIndex: jsii.Number(123),
//   						SdpUrl: jsii.String("sdpUrl"),
//   					},
//   				},
//   				AudioSdps: []interface{}{
//   					&InputSdpLocationProperty{
//   						MediaIndex: jsii.Number(123),
//   						SdpUrl: jsii.String("sdpUrl"),
//   					},
//   				},
//   				VideoSdp: &InputSdpLocationProperty{
//   					MediaIndex: jsii.Number(123),
//   					SdpUrl: jsii.String("sdpUrl"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsettings.html
//
type CfnInput_Smpte2110ReceiverGroupSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-smpte2110receivergroupsettings.html#cfn-medialive-input-smpte2110receivergroupsettings-smpte2110receivergroups
	//
	Smpte2110ReceiverGroups interface{} `field:"optional" json:"smpte2110ReceiverGroups" yaml:"smpte2110ReceiverGroups"`
}

