package previewawsquicksightmixins


// The option that determines the text display weight, or boldness.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fontWeightProperty := &FontWeightProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fontweight.html
//
type CfnAnalysisPropsMixin_FontWeightProperty struct {
	// The lexical name for the level of boldness of the text display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-fontweight.html#cfn-quicksight-analysis-fontweight-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

