package awsquicksight


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
type CfnTemplate_ThousandSeparatorOptionsProperty struct {
	// `CfnTemplate.ThousandSeparatorOptionsProperty.Symbol`.
	Symbol *string `field:"optional" json:"symbol" yaml:"symbol"`
	// `CfnTemplate.ThousandSeparatorOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

