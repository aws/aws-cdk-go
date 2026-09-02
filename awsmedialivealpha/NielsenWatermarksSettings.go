package awsmedialivealpha


// Nielsen watermark settings for audio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var nielsenCbetStepaside NielsenCbetStepaside
//   var nielsenDistributionType NielsenDistributionType
//   var nielsenWatermarkTimezone NielsenWatermarkTimezone
//
//   nielsenWatermarksSettings := &NielsenWatermarksSettings{
//   	CbetSettings: &NielsenCbetSettings{
//   		CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   		Csid: jsii.String("csid"),
//
//   		// the properties below are optional
//   		CbetStepaside: nielsenCbetStepaside,
//   	},
//   	DistributionType: nielsenDistributionType,
//   	NaesIiNwSettings: &NielsenNaesIiNwSettings{
//   		CheckDigitString: jsii.String("checkDigitString"),
//   		Sid: jsii.Number(123),
//
//   		// the properties below are optional
//   		Timezone: nielsenWatermarkTimezone,
//   	},
//   }
//
// Experimental.
type NielsenWatermarksSettings struct {
	// Nielsen CBET watermark settings.
	// Default: - no CBET watermarks.
	//
	// Experimental.
	CbetSettings *NielsenCbetSettings `field:"optional" json:"cbetSettings" yaml:"cbetSettings"`
	// The distribution type for the watermark.
	// Default: - service default.
	//
	// Experimental.
	DistributionType NielsenDistributionType `field:"optional" json:"distributionType" yaml:"distributionType"`
	// Nielsen NAES II/NW watermark settings.
	// Default: - no NAES II/NW watermarks.
	//
	// Experimental.
	NaesIiNwSettings *NielsenNaesIiNwSettings `field:"optional" json:"naesIiNwSettings" yaml:"naesIiNwSettings"`
}

