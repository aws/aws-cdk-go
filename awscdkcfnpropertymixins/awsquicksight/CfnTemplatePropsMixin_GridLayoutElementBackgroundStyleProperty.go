package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gridLayoutElementBackgroundStyleProperty := &GridLayoutElementBackgroundStyleProperty{
//   	Color: jsii.String("color"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementbackgroundstyle.html
//
type CfnTemplatePropsMixin_GridLayoutElementBackgroundStyleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementbackgroundstyle.html#cfn-quicksight-template-gridlayoutelementbackgroundstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementbackgroundstyle.html#cfn-quicksight-template-gridlayoutelementbackgroundstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

