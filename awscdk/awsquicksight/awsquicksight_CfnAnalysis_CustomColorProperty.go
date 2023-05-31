package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customColorProperty := &CustomColorProperty{
//   	Color: jsii.String("color"),
//
//   	// the properties below are optional
//   	FieldValue: jsii.String("fieldValue"),
//   	SpecialValue: jsii.String("specialValue"),
//   }
//
type CfnAnalysis_CustomColorProperty struct {
	// `CfnAnalysis.CustomColorProperty.Color`.
	Color *string `field:"required" json:"color" yaml:"color"`
	// `CfnAnalysis.CustomColorProperty.FieldValue`.
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// `CfnAnalysis.CustomColorProperty.SpecialValue`.
	SpecialValue *string `field:"optional" json:"specialValue" yaml:"specialValue"`
}

