package awsquicksight


// The side border options for a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableSideBorderOptionsProperty := &TableSideBorderOptionsProperty{
//   	Bottom: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	InnerHorizontal: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	InnerVertical: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	Left: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	Right: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   	Top: &TableBorderOptionsProperty{
//   		Color: jsii.String("color"),
//   		Style: jsii.String("style"),
//   		Thickness: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html
//
type CfnTemplate_TableSideBorderOptionsProperty struct {
	// The table border options of the bottom border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-bottom
	//
	Bottom interface{} `field:"optional" json:"bottom" yaml:"bottom"`
	// The table border options of the inner horizontal border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-innerhorizontal
	//
	InnerHorizontal interface{} `field:"optional" json:"innerHorizontal" yaml:"innerHorizontal"`
	// The table border options of the inner vertical border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-innervertical
	//
	InnerVertical interface{} `field:"optional" json:"innerVertical" yaml:"innerVertical"`
	// The table border options of the left border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-left
	//
	Left interface{} `field:"optional" json:"left" yaml:"left"`
	// The table border options of the right border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-right
	//
	Right interface{} `field:"optional" json:"right" yaml:"right"`
	// The table border options of the top border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-top
	//
	Top interface{} `field:"optional" json:"top" yaml:"top"`
}

