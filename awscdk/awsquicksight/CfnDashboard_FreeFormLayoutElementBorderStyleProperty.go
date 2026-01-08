package awsquicksight


// The background style configuration of a free-form layout element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeFormLayoutElementBorderStyleProperty := &FreeFormLayoutElementBorderStyleProperty{
//   	Color: jsii.String("color"),
//   	Visibility: jsii.String("visibility"),
//   	Width: jsii.String("width"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementborderstyle.html
//
type CfnDashboard_FreeFormLayoutElementBorderStyleProperty struct {
	// The border color of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementborderstyle.html#cfn-quicksight-dashboard-freeformlayoutelementborderstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The border visibility of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementborderstyle.html#cfn-quicksight-dashboard-freeformlayoutelementborderstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// The border width of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementborderstyle.html#cfn-quicksight-dashboard-freeformlayoutelementborderstyle-width
	//
	Width *string `field:"optional" json:"width" yaml:"width"`
}

