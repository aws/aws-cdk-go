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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-colorscale.html
//
type CfnTemplate_ColorScaleProperty struct {
	// Determines the color fill type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-colorscale.html#cfn-quicksight-template-colorscale-colorfilltype
	//
	ColorFillType *string `field:"required" json:"colorFillType" yaml:"colorFillType"`
	// Determines the list of colors that are applied to the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-colorscale.html#cfn-quicksight-template-colorscale-colors
	//
	Colors interface{} `field:"required" json:"colors" yaml:"colors"`
	// Determines the color that is applied to null values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-colorscale.html#cfn-quicksight-template-colorscale-nullvaluecolor
	//
	NullValueColor interface{} `field:"optional" json:"nullValueColor" yaml:"nullValueColor"`
}

