package awsquicksight


// The configuration options that determine how missing data is treated during the rendering of a line chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   missingDataConfigurationProperty := &MissingDataConfigurationProperty{
//   	TreatmentOption: jsii.String("treatmentOption"),
//   }
//
type CfnTemplate_MissingDataConfigurationProperty struct {
	// The treatment option that determines how missing data should be rendered. Choose from the following options:.
	//
	// - `INTERPOLATE` : Interpolate missing values between the prior and the next known value.
	// - `SHOW_AS_ZERO` : Show missing values as the value `0` .
	// - `SHOW_AS_BLANK` : Display a blank space when rendering missing data.
	TreatmentOption *string `field:"optional" json:"treatmentOption" yaml:"treatmentOption"`
}

