package awsquicksight


// The exclude period of `TimeRangeFilter` or `RelativeDatesFilter` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   excludePeriodConfigurationProperty := &ExcludePeriodConfigurationProperty{
//   	Amount: jsii.Number(123),
//   	Granularity: jsii.String("granularity"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-excludeperiodconfiguration.html
//
type CfnAnalysis_ExcludePeriodConfigurationProperty struct {
	// The amount or number of the exclude period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-excludeperiodconfiguration.html#cfn-quicksight-analysis-excludeperiodconfiguration-amount
	//
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// The granularity or unit (day, month, year) of the exclude period.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-excludeperiodconfiguration.html#cfn-quicksight-analysis-excludeperiodconfiguration-granularity
	//
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// The status of the exclude period. Choose from the following options:.
	//
	// - `ENABLED`
	// - `DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-excludeperiodconfiguration.html#cfn-quicksight-analysis-excludeperiodconfiguration-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

