package awsmedialivealpha


// Properties for HDR10 color space settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hdr10SettingsProps := &Hdr10SettingsProps{
//   	MaxCll: jsii.Number(123),
//   	MaxFall: jsii.Number(123),
//   }
//
// Experimental.
type Hdr10SettingsProps struct {
	// Maximum Content Light Level — the maximum light level of any single pixel in nits.
	// Default: - service default.
	//
	// Experimental.
	MaxCll *float64 `field:"optional" json:"maxCll" yaml:"maxCll"`
	// Maximum Frame Average Light Level — the maximum average light level of any single frame in nits.
	// Default: - service default.
	//
	// Experimental.
	MaxFall *float64 `field:"optional" json:"maxFall" yaml:"maxFall"`
}

