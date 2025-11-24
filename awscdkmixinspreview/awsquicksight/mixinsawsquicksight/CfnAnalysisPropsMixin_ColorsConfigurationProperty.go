package mixinsawsquicksight


// The color configurations for a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   colorsConfigurationProperty := &ColorsConfigurationProperty{
//   	CustomColors: []interface{}{
//   		&CustomColorProperty{
//   			Color: jsii.String("color"),
//   			FieldValue: jsii.String("fieldValue"),
//   			SpecialValue: jsii.String("specialValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-colorsconfiguration.html
//
type CfnAnalysisPropsMixin_ColorsConfigurationProperty struct {
	// A list of up to 50 custom colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-colorsconfiguration.html#cfn-quicksight-analysis-colorsconfiguration-customcolors
	//
	CustomColors interface{} `field:"optional" json:"customColors" yaml:"customColors"`
}

