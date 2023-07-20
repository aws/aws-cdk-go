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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-freeformlayoutelementborderstyle.html
//
type CfnAnalysis_FreeFormLayoutElementBorderStyleProperty struct {
	// The border color of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-freeformlayoutelementborderstyle.html#cfn-quicksight-analysis-freeformlayoutelementborderstyle-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The border visibility of a free-form layout element.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-freeformlayoutelementborderstyle.html#cfn-quicksight-analysis-freeformlayoutelementborderstyle-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

