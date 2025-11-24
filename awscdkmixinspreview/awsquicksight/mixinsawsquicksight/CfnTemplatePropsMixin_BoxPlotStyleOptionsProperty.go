package mixinsawsquicksight


// The style options of the box plot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   boxPlotStyleOptionsProperty := &BoxPlotStyleOptionsProperty{
//   	FillStyle: jsii.String("fillStyle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotstyleoptions.html
//
type CfnTemplatePropsMixin_BoxPlotStyleOptionsProperty struct {
	// The fill styles (solid, transparent) of the box plot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-boxplotstyleoptions.html#cfn-quicksight-template-boxplotstyleoptions-fillstyle
	//
	FillStyle *string `field:"optional" json:"fillStyle" yaml:"fillStyle"`
}

