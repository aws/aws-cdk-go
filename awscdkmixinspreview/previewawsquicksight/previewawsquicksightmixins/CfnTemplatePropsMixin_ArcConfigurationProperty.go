package previewawsquicksightmixins


// The arc configuration of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   arcConfigurationProperty := &ArcConfigurationProperty{
//   	ArcAngle: jsii.Number(123),
//   	ArcThickness: jsii.String("arcThickness"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-arcconfiguration.html
//
type CfnTemplatePropsMixin_ArcConfigurationProperty struct {
	// The option that determines the arc angle of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-arcconfiguration.html#cfn-quicksight-template-arcconfiguration-arcangle
	//
	ArcAngle *float64 `field:"optional" json:"arcAngle" yaml:"arcAngle"`
	// The options that determine the arc thickness of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-arcconfiguration.html#cfn-quicksight-template-arcconfiguration-arcthickness
	//
	ArcThickness *string `field:"optional" json:"arcThickness" yaml:"arcThickness"`
}

