package previewawsquicksightmixins


// The options that determine the presentation of the progress bar of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   progressBarOptionsProperty := &ProgressBarOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-progressbaroptions.html
//
type CfnTemplatePropsMixin_ProgressBarOptionsProperty struct {
	// The visibility of the progress bar.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-progressbaroptions.html#cfn-quicksight-template-progressbaroptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

