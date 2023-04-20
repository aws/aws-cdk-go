package awsquicksight


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
type CfnAnalysis_DataBarsOptionsProperty struct {
	// `CfnAnalysis.DataBarsOptionsProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnAnalysis.DataBarsOptionsProperty.NegativeColor`.
	NegativeColor *string `field:"optional" json:"negativeColor" yaml:"negativeColor"`
	// `CfnAnalysis.DataBarsOptionsProperty.PositiveColor`.
	PositiveColor *string `field:"optional" json:"positiveColor" yaml:"positiveColor"`
}

