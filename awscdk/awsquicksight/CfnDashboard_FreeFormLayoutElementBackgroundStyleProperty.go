package awsquicksight


// The background style configuration of a free-form layout element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeFormLayoutElementBackgroundStyleProperty := &FreeFormLayoutElementBackgroundStyleProperty{
//   	Color: jsii.String("color"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.html
//
type CfnDashboard_FreeFormLayoutElementBackgroundStyleProperty struct {
	// The background color of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.html#cfn-quicksight-dashboard-freeformlayoutelementbackgroundstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The background visibility of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-freeformlayoutelementbackgroundstyle.html#cfn-quicksight-dashboard-freeformlayoutelementbackgroundstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

