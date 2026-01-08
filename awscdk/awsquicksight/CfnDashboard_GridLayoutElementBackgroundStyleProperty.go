package awsquicksight


// The background style configuration of a grid layout element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutElementBackgroundStyleProperty := &GridLayoutElementBackgroundStyleProperty{
//   	Color: jsii.String("color"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementbackgroundstyle.html
//
type CfnDashboard_GridLayoutElementBackgroundStyleProperty struct {
	// The background color of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementbackgroundstyle.html#cfn-quicksight-dashboard-gridlayoutelementbackgroundstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The background visibility of a grid layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutelementbackgroundstyle.html#cfn-quicksight-dashboard-gridlayoutelementbackgroundstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

