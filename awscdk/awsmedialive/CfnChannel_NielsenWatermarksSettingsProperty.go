package awsmedialive


// Settings to configure Nielsen Watermarks in the audio encode.
//
// The parent of this entity is AudioWatermarkSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nielsenWatermarksSettingsProperty := &NielsenWatermarksSettingsProperty{
//   	NielsenCbetSettings: &NielsenCBETProperty{
//   		CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   		CbetStepaside: jsii.String("cbetStepaside"),
//   		Csid: jsii.String("csid"),
//   	},
//   	NielsenDistributionType: jsii.String("nielsenDistributionType"),
//   	NielsenNaesIiNwSettings: &NielsenNaesIiNwProperty{
//   		CheckDigitString: jsii.String("checkDigitString"),
//   		Sid: jsii.Number(123),
//   		Timezone: jsii.String("timezone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenwatermarkssettings.html
//
type CfnChannel_NielsenWatermarksSettingsProperty struct {
	// Complete these fields only if you want to insert watermarks of type Nielsen CBET.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenwatermarkssettings.html#cfn-medialive-channel-nielsenwatermarkssettings-nielsencbetsettings
	//
	NielsenCbetSettings interface{} `field:"optional" json:"nielsenCbetSettings" yaml:"nielsenCbetSettings"`
	// Choose the distribution types that you want to assign to the watermarks: - PROGRAM_CONTENT - FINAL_DISTRIBUTOR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenwatermarkssettings.html#cfn-medialive-channel-nielsenwatermarkssettings-nielsendistributiontype
	//
	NielsenDistributionType *string `field:"optional" json:"nielsenDistributionType" yaml:"nielsenDistributionType"`
	// Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-nielsenwatermarkssettings.html#cfn-medialive-channel-nielsenwatermarkssettings-nielsennaesiinwsettings
	//
	NielsenNaesIiNwSettings interface{} `field:"optional" json:"nielsenNaesIiNwSettings" yaml:"nielsenNaesIiNwSettings"`
}

