package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexContainerSettingsProperty := &MultiplexContainerSettingsProperty{
//   	MultiplexM2TsSettings: &MultiplexM2tsSettingsProperty{
//   		AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   		Arib: jsii.String("arib"),
//   		AudioBufferModel: jsii.String("audioBufferModel"),
//   		AudioFramesPerPes: jsii.Number(123),
//   		AudioStreamType: jsii.String("audioStreamType"),
//   		CcDescriptor: jsii.String("ccDescriptor"),
//   		Ebif: jsii.String("ebif"),
//   		EsRateInPes: jsii.String("esRateInPes"),
//   		Klv: jsii.String("klv"),
//   		NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		PcrControl: jsii.String("pcrControl"),
//   		PcrPeriod: jsii.Number(123),
//   		Scte35Control: jsii.String("scte35Control"),
//   		Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexcontainersettings.html
//
type CfnChannelPropsMixin_MultiplexContainerSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexcontainersettings.html#cfn-medialive-channel-multiplexcontainersettings-multiplexm2tssettings
	//
	MultiplexM2TsSettings interface{} `field:"optional" json:"multiplexM2TsSettings" yaml:"multiplexM2TsSettings"`
}

