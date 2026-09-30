package awsquicksight


// The display options for tile borders for visuals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   borderStyleProperty := &BorderStyleProperty{
//   	Color: jsii.String("color"),
//   	Show: jsii.Boolean(false),
//   	Width: jsii.String("width"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-borderstyle.html
//
type CfnTheme_BorderStyleProperty struct {
	// String to encapsulate the most generic way Color can be formatted (words, hexStrings etc).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-borderstyle.html#cfn-quicksight-theme-borderstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The option to enable display of borders for visuals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-borderstyle.html#cfn-quicksight-theme-borderstyle-show
	//
	Show interface{} `field:"optional" json:"show" yaml:"show"`
	// String to encapsulate the most generic way Width can be formatted with whatever units (px, em etc).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-borderstyle.html#cfn-quicksight-theme-borderstyle-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
}

