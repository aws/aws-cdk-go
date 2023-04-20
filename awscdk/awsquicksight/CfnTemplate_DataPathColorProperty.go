package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathColorProperty := &DataPathColorProperty{
//   	Color: jsii.String("color"),
//   	Element: &DataPathValueProperty{
//   		FieldId: jsii.String("fieldId"),
//   		FieldValue: jsii.String("fieldValue"),
//   	},
//
//   	// the properties below are optional
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
type CfnTemplate_DataPathColorProperty struct {
	// `CfnTemplate.DataPathColorProperty.Color`.
	Color *string `field:"required" json:"color" yaml:"color"`
	// `CfnTemplate.DataPathColorProperty.Element`.
	Element interface{} `field:"required" json:"element" yaml:"element"`
	// `CfnTemplate.DataPathColorProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

