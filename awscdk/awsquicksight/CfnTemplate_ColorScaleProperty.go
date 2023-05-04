package awsquicksight


// Determines the color scale that is applied to the visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   colorScaleProperty := &ColorScaleProperty{
//   	ColorFillType: jsii.String("colorFillType"),
//   	Colors: []interface{}{
//   		&DataColorProperty{
//   			Color: jsii.String("color"),
//   			DataValue: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	NullValueColor: &DataColorProperty{
//   		Color: jsii.String("color"),
//   		DataValue: jsii.Number(123),
//   	},
//   }
//
type CfnTemplate_ColorScaleProperty struct {
	// Determines the color fill type.
	ColorFillType *string `field:"required" json:"colorFillType" yaml:"colorFillType"`
	// Determines the list of colors that are applied to the visual.
	Colors interface{} `field:"required" json:"colors" yaml:"colors"`
	// Determines the color that is applied to null values.
	NullValueColor interface{} `field:"optional" json:"nullValueColor" yaml:"nullValueColor"`
}

