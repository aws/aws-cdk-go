package awsquicksight


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
type CfnAnalysis_ColorScaleProperty struct {
	// `CfnAnalysis.ColorScaleProperty.ColorFillType`.
	ColorFillType *string `field:"required" json:"colorFillType" yaml:"colorFillType"`
	// `CfnAnalysis.ColorScaleProperty.Colors`.
	Colors interface{} `field:"required" json:"colors" yaml:"colors"`
	// `CfnAnalysis.ColorScaleProperty.NullValueColor`.
	NullValueColor interface{} `field:"optional" json:"nullValueColor" yaml:"nullValueColor"`
}

