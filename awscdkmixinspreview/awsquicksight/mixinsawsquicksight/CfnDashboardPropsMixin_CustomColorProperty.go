package mixinsawsquicksight


// Determines the color that's applied to a particular data value in a column.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customColorProperty := &CustomColorProperty{
//   	Color: jsii.String("color"),
//   	FieldValue: jsii.String("fieldValue"),
//   	SpecialValue: jsii.String("specialValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customcolor.html
//
type CfnDashboardPropsMixin_CustomColorProperty struct {
	// The color that is applied to the data value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customcolor.html#cfn-quicksight-dashboard-customcolor-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The data value that the color is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customcolor.html#cfn-quicksight-dashboard-customcolor-fieldvalue
	//
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// The value of a special data value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customcolor.html#cfn-quicksight-dashboard-customcolor-specialvalue
	//
	SpecialValue *string `field:"optional" json:"specialValue" yaml:"specialValue"`
}

