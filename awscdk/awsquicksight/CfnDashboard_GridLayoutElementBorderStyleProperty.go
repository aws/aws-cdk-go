package awsquicksight


// The border style configuration of a grid layout element.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementborderstyle.html
//
type CfnDashboard_GridLayoutElementBorderStyleProperty struct {
	// The border color of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementborderstyle.html#cfn-quicksight-dashboard-gridlayoutelementborderstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The border visibility of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementborderstyle.html#cfn-quicksight-dashboard-gridlayoutelementborderstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// The border width of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementborderstyle.html#cfn-quicksight-dashboard-gridlayoutelementborderstyle-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
}

