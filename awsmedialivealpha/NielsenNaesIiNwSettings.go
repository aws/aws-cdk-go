package awsmedialivealpha


// Nielsen NAES II/NW watermark settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var nielsenWatermarkTimezone NielsenWatermarkTimezone
//
//   nielsenNaesIiNwSettings := &NielsenNaesIiNwSettings{
//   	CheckDigitString: jsii.String("checkDigitString"),
//   	Sid: jsii.Number(123),
//
//   	// the properties below are optional
//   	Timezone: nielsenWatermarkTimezone,
//   }
//
// Experimental.
type NielsenNaesIiNwSettings struct {
	// The check digit string for the watermark.
	// Experimental.
	CheckDigitString *string `field:"required" json:"checkDigitString" yaml:"checkDigitString"`
	// The Nielsen Source ID (SID).
	// Experimental.
	Sid *float64 `field:"required" json:"sid" yaml:"sid"`
	// The timezone for the timestamps in the watermark.
	// Default: - Coordinated Universal Time (UTC).
	//
	// Experimental.
	Timezone NielsenWatermarkTimezone `field:"optional" json:"timezone" yaml:"timezone"`
}

