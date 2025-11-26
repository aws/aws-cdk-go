package previewawsquicksightmixins


// The options that determine the arc thickness of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   arcOptionsProperty := &ArcOptionsProperty{
//   	ArcThickness: jsii.String("arcThickness"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-arcoptions.html
//
type CfnTemplatePropsMixin_ArcOptionsProperty struct {
	// The arc thickness of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-arcoptions.html#cfn-quicksight-template-arcoptions-arcthickness
	//
	ArcThickness *string `field:"optional" json:"arcThickness" yaml:"arcThickness"`
}

