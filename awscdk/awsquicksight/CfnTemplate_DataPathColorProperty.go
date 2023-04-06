package awsquicksight


// The color map that determines the color options for a particular element.
//
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
	// The color that needs to be applied to the element.
	Color *string `field:"required" json:"color" yaml:"color"`
	// The element that the color needs to be applied to.
	Element interface{} `field:"required" json:"element" yaml:"element"`
	// The time granularity of the field that the color needs to be applied to.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

