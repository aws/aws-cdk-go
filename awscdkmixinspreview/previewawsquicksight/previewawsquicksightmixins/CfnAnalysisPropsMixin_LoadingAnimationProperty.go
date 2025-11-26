package previewawsquicksightmixins


// The configuration of loading animation in free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loadingAnimationProperty := &LoadingAnimationProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-loadinganimation.html
//
type CfnAnalysisPropsMixin_LoadingAnimationProperty struct {
	// The visibility configuration of `LoadingAnimation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-loadinganimation.html#cfn-quicksight-analysis-loadinganimation-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

