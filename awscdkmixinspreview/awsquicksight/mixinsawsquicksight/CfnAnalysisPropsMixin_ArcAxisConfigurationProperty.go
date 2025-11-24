package mixinsawsquicksight


// The arc axis configuration of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   arcAxisConfigurationProperty := &ArcAxisConfigurationProperty{
//   	Range: &ArcAxisDisplayRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	ReserveRange: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-arcaxisconfiguration.html
//
type CfnAnalysisPropsMixin_ArcAxisConfigurationProperty struct {
	// The arc axis range of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-arcaxisconfiguration.html#cfn-quicksight-analysis-arcaxisconfiguration-range
	//
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The reserved range of the arc axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-arcaxisconfiguration.html#cfn-quicksight-analysis-arcaxisconfiguration-reserverange
	//
	// Default: - 0.
	//
	ReserveRange *float64 `field:"optional" json:"reserveRange" yaml:"reserveRange"`
}

