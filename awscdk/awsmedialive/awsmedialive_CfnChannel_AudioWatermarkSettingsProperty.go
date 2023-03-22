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
//   audioWatermarkSettingsProperty := &audioWatermarkSettingsProperty{
//   	nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   		nielsenCbetSettings: &nielsenCBETProperty{
//   			cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   			cbetStepaside: jsii.String("cbetStepaside"),
//   			csid: jsii.String("csid"),
//   		},
//   		nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   		nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   			checkDigitString: jsii.String("checkDigitString"),
//   			sid: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_AudioWatermarkSettingsProperty struct {
	// Settings to configure Nielsen Watermarks in the audio encode.
	NielsenWatermarksSettings interface{} `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}

