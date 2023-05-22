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
type CfnTemplate_CustomColorProperty struct {
	// `CfnTemplate.CustomColorProperty.Color`.
	Color *string `field:"required" json:"color" yaml:"color"`
	// `CfnTemplate.CustomColorProperty.FieldValue`.
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// `CfnTemplate.CustomColorProperty.SpecialValue`.
	SpecialValue *string `field:"optional" json:"specialValue" yaml:"specialValue"`
}

