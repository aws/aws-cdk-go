package awsmedialive


// Configuration of a Multiplex output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexOutputSettingsProperty := &MultiplexOutputSettingsProperty{
//   	ContainerSettings: &MultiplexContainerSettingsProperty{
//   		MultiplexM2TsSettings: &MultiplexM2tsSettingsProperty{
//   			AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   			Arib: jsii.String("arib"),
//   			AudioBufferModel: jsii.String("audioBufferModel"),
//   			AudioFramesPerPes: jsii.Number(123),
//   			AudioStreamType: jsii.String("audioStreamType"),
//   			CcDescriptor: jsii.String("ccDescriptor"),
//   			Ebif: jsii.String("ebif"),
//   			EsRateInPes: jsii.String("esRateInPes"),
//   			Klv: jsii.String("klv"),
//   			NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			PcrControl: jsii.String("pcrControl"),
//   			PcrPeriod: jsii.Number(123),
//   			Scte35Control: jsii.String("scte35Control"),
//   			Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   		},
//   	},
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexoutputsettings.html
//
type CfnChannel_MultiplexOutputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexoutputsettings.html#cfn-medialive-channel-multiplexoutputsettings-containersettings
	//
	ContainerSettings interface{} `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// Destination is a Multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexoutputsettings.html#cfn-medialive-channel-multiplexoutputsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
}

