package mixinsawsquicksight


// The options that determine how a date axis is displayed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dateAxisOptionsProperty := &DateAxisOptionsProperty{
//   	MissingDateVisibility: jsii.String("missingDateVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-dateaxisoptions.html
//
type CfnTemplatePropsMixin_DateAxisOptionsProperty struct {
	// Determines whether or not missing dates are displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-dateaxisoptions.html#cfn-quicksight-template-dateaxisoptions-missingdatevisibility
	//
	MissingDateVisibility *string `field:"optional" json:"missingDateVisibility" yaml:"missingDateVisibility"`
}

