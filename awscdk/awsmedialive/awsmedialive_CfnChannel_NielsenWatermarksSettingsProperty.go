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
//   nielsenWatermarksSettingsProperty := &nielsenWatermarksSettingsProperty{
//   	nielsenCbetSettings: &nielsenCBETProperty{
//   		cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   		cbetStepaside: jsii.String("cbetStepaside"),
//   		csid: jsii.String("csid"),
//   	},
//   	nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   	nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   		checkDigitString: jsii.String("checkDigitString"),
//   		sid: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_NielsenWatermarksSettingsProperty struct {
	// Complete these fields only if you want to insert watermarks of type Nielsen CBET.
	NielsenCbetSettings interface{} `field:"optional" json:"nielsenCbetSettings" yaml:"nielsenCbetSettings"`
	// Choose the distribution types that you want to assign to the watermarks: - PROGRAM_CONTENT - FINAL_DISTRIBUTOR.
	NielsenDistributionType *string `field:"optional" json:"nielsenDistributionType" yaml:"nielsenDistributionType"`
	// Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).
	NielsenNaesIiNwSettings interface{} `field:"optional" json:"nielsenNaesIiNwSettings" yaml:"nielsenNaesIiNwSettings"`
}

