package awsquicksight


// The options for data bars.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataBarsOptionsProperty := &DataBarsOptionsProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	NegativeColor: jsii.String("negativeColor"),
//   	PositiveColor: jsii.String("positiveColor"),
//   }
//
type CfnTemplate_DataBarsOptionsProperty struct {
	// The field ID for the data bars options.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The color of the negative data bar.
	NegativeColor *string `field:"optional" json:"negativeColor" yaml:"negativeColor"`
	// The color of the positive data bar.
	PositiveColor *string `field:"optional" json:"positiveColor" yaml:"positiveColor"`
}

