package mixinsawsmedialive


// The setup of ad avail handling in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   availConfigurationProperty := &AvailConfigurationProperty{
//   	AvailSettings: &AvailSettingsProperty{
//   		Esam: &EsamProperty{
//   			AcquisitionPointId: jsii.String("acquisitionPointId"),
//   			AdAvailOffset: jsii.Number(123),
//   			PasswordParam: jsii.String("passwordParam"),
//   			PoisEndpoint: jsii.String("poisEndpoint"),
//   			Username: jsii.String("username"),
//   			ZoneIdentity: jsii.String("zoneIdentity"),
//   		},
//   		Scte35SpliceInsert: &Scte35SpliceInsertProperty{
//   			AdAvailOffset: jsii.Number(123),
//   			NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   		Scte35TimeSignalApos: &Scte35TimeSignalAposProperty{
//   			AdAvailOffset: jsii.Number(123),
//   			NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   			WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   		},
//   	},
//   	Scte35SegmentationScope: jsii.String("scte35SegmentationScope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availconfiguration.html
//
type CfnChannelPropsMixin_AvailConfigurationProperty struct {
	// The setup of ad avail handling in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availconfiguration.html#cfn-medialive-channel-availconfiguration-availsettings
	//
	AvailSettings interface{} `field:"optional" json:"availSettings" yaml:"availSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availconfiguration.html#cfn-medialive-channel-availconfiguration-scte35segmentationscope
	//
	Scte35SegmentationScope *string `field:"optional" json:"scte35SegmentationScope" yaml:"scte35SegmentationScope"`
}

