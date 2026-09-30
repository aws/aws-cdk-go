package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutElementBorderStyleProperty := &GridLayoutElementBorderStyleProperty{
//   	Color: jsii.String("color"),
//   	Visibility: jsii.String("visibility"),
//   	Width: jsii.String("width"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementborderstyle.html
//
type CfnTemplate_GridLayoutElementBorderStyleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementborderstyle.html#cfn-quicksight-template-gridlayoutelementborderstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementborderstyle.html#cfn-quicksight-template-gridlayoutelementborderstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// String to encapsulate the most generic way Width can be formatted with whatever units (px, em etc).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutelementborderstyle.html#cfn-quicksight-template-gridlayoutelementborderstyle-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
}

