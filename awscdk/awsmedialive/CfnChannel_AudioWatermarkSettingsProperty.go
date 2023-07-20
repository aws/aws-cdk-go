package awsmedialive


// Audio Watermark Settings.
//
// The parent of this entity is AudioDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioWatermarkSettingsProperty := &AudioWatermarkSettingsProperty{
//   	NielsenWatermarksSettings: &NielsenWatermarksSettingsProperty{
//   		NielsenCbetSettings: &NielsenCBETProperty{
//   			CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   			CbetStepaside: jsii.String("cbetStepaside"),
//   			Csid: jsii.String("csid"),
//   		},
//   		NielsenDistributionType: jsii.String("nielsenDistributionType"),
//   		NielsenNaesIiNwSettings: &NielsenNaesIiNwProperty{
//   			CheckDigitString: jsii.String("checkDigitString"),
//   			Sid: jsii.Number(123),
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiowatermarksettings.html
//
type CfnChannel_AudioWatermarkSettingsProperty struct {
	// Settings to configure Nielsen Watermarks in the audio encode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiowatermarksettings.html#cfn-medialive-channel-audiowatermarksettings-nielsenwatermarkssettings
	//
	NielsenWatermarksSettings interface{} `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}

