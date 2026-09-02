package awsmedialivealpha


// Audio watermarking settings.
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
//   audioWatermarkSettings := &AudioWatermarkSettings{
//   	NielsenWatermarks: &NielsenWatermarksSettings{
//   		CbetSettings: &NielsenCbetSettings{
//   			CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   			Csid: jsii.String("csid"),
//
//   			// the properties below are optional
//   			CbetStepaside: nielsenCbetStepaside,
//   		},
//   		DistributionType: nielsenDistributionType,
//   		NaesIiNwSettings: &NielsenNaesIiNwSettings{
//   			CheckDigitString: jsii.String("checkDigitString"),
//   			Sid: jsii.Number(123),
//
//   			// the properties below are optional
//   			Timezone: nielsenWatermarkTimezone,
//   		},
//   	},
//   }
//
// Experimental.
type AudioWatermarkSettings struct {
	// Nielsen watermark settings.
	// Default: - no Nielsen watermarks.
	//
	// Experimental.
	NielsenWatermarks *NielsenWatermarksSettings `field:"optional" json:"nielsenWatermarks" yaml:"nielsenWatermarks"`
}

