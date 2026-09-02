package awsmedialivealpha


// Nielsen CBET watermark settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var nielsenCbetStepaside NielsenCbetStepaside
//
//   nielsenCbetSettings := &NielsenCbetSettings{
//   	CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   	Csid: jsii.String("csid"),
//
//   	// the properties below are optional
//   	CbetStepaside: nielsenCbetStepaside,
//   }
//
// Experimental.
type NielsenCbetSettings struct {
	// The CBET check digit string.
	// Experimental.
	CbetCheckDigitString *string `field:"required" json:"cbetCheckDigitString" yaml:"cbetCheckDigitString"`
	// The CBET Source ID (CSID).
	// Experimental.
	Csid *string `field:"required" json:"csid" yaml:"csid"`
	// The CBET stepaside behavior when prior encoding is detected.
	// Default: - service default.
	//
	// Experimental.
	CbetStepaside NielsenCbetStepaside `field:"optional" json:"cbetStepaside" yaml:"cbetStepaside"`
}

