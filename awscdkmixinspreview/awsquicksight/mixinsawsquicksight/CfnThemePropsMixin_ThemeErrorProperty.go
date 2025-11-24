package mixinsawsquicksight


// Theme error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   themeErrorProperty := &ThemeErrorProperty{
//   	Message: jsii.String("message"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeerror.html
//
type CfnThemePropsMixin_ThemeErrorProperty struct {
	// The error message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeerror.html#cfn-quicksight-theme-themeerror-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The type of error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-themeerror.html#cfn-quicksight-theme-themeerror-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

