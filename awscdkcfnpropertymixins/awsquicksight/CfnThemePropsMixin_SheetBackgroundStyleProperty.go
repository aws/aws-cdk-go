package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sheetBackgroundStyleProperty := &SheetBackgroundStyleProperty{
//   	Color: jsii.String("color"),
//   	Gradient: jsii.String("gradient"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetbackgroundstyle.html
//
type CfnThemePropsMixin_SheetBackgroundStyleProperty struct {
	// String to encapsulate the most generic way Color can be formatted (words, hexStrings etc).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetbackgroundstyle.html#cfn-quicksight-theme-sheetbackgroundstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetbackgroundstyle.html#cfn-quicksight-theme-sheetbackgroundstyle-gradient
	//
	Gradient *string `field:"optional" json:"gradient" yaml:"gradient"`
}

