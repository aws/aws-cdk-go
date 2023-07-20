package awsmedialive


// The settings for the ad avail setup in the output.
//
// The parent of this entity is AvailConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availSettingsProperty := &AvailSettingsProperty{
//   	Esam: &EsamProperty{
//   		AcquisitionPointId: jsii.String("acquisitionPointId"),
//   		AdAvailOffset: jsii.Number(123),
//   		PasswordParam: jsii.String("passwordParam"),
//   		PoisEndpoint: jsii.String("poisEndpoint"),
//   		Username: jsii.String("username"),
//   		ZoneIdentity: jsii.String("zoneIdentity"),
//   	},
//   	Scte35SpliceInsert: &Scte35SpliceInsertProperty{
//   		AdAvailOffset: jsii.Number(123),
//   		NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   		WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   	},
//   	Scte35TimeSignalApos: &Scte35TimeSignalAposProperty{
//   		AdAvailOffset: jsii.Number(123),
//   		NoRegionalBlackoutFlag: jsii.String("noRegionalBlackoutFlag"),
//   		WebDeliveryAllowedFlag: jsii.String("webDeliveryAllowedFlag"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html
//
type CfnChannel_AvailSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html#cfn-medialive-channel-availsettings-esam
	//
	Esam interface{} `field:"optional" json:"esam" yaml:"esam"`
	// The setup for SCTE-35 splice insert handling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html#cfn-medialive-channel-availsettings-scte35spliceinsert
	//
	Scte35SpliceInsert interface{} `field:"optional" json:"scte35SpliceInsert" yaml:"scte35SpliceInsert"`
	// The setup for SCTE-35 time signal APOS handling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availsettings.html#cfn-medialive-channel-availsettings-scte35timesignalapos
	//
	Scte35TimeSignalApos interface{} `field:"optional" json:"scte35TimeSignalApos" yaml:"scte35TimeSignalApos"`
}

