package awsquicksight


// The display options for tile borders for visuals.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   borderStyleProperty := &BorderStyleProperty{
//   	Show: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-borderstyle.html
//
type CfnTheme_BorderStyleProperty struct {
	// The option to enable display of borders for visuals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-borderstyle.html#cfn-quicksight-theme-borderstyle-show
	//
	Show interface{} `field:"optional" json:"show" yaml:"show"`
}

