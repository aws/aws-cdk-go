package previewawsquicksightmixins


// The exclude period of `TimeRangeFilter` or `RelativeDatesFilter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   excludePeriodConfigurationProperty := &ExcludePeriodConfigurationProperty{
//   	Amount: jsii.Number(123),
//   	Granularity: jsii.String("granularity"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-excludeperiodconfiguration.html
//
type CfnDashboardPropsMixin_ExcludePeriodConfigurationProperty struct {
	// The amount or number of the exclude period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-excludeperiodconfiguration.html#cfn-quicksight-dashboard-excludeperiodconfiguration-amount
	//
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// The granularity or unit (day, month, year) of the exclude period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-excludeperiodconfiguration.html#cfn-quicksight-dashboard-excludeperiodconfiguration-granularity
	//
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// The status of the exclude period. Choose from the following options:.
	//
	// - `ENABLED`
	// - `DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-excludeperiodconfiguration.html#cfn-quicksight-dashboard-excludeperiodconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

