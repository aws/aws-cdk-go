package awsquicksight


// The options that determine the thousands separator configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thousandSeparatorOptionsProperty := &ThousandSeparatorOptionsProperty{
//   	Symbol: jsii.String("symbol"),
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_ThousandSeparatorOptionsProperty struct {
	// Determines the thousands separator symbol.
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
	// Determines the visibility of the thousands separator.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

