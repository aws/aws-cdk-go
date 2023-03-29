package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relativeDatesFilterProperty := &RelativeDatesFilterProperty{
//   	AnchorDateConfiguration: &AnchorDateConfigurationProperty{
//   		AnchorOption: jsii.String("anchorOption"),
//   		ParameterName: jsii.String("parameterName"),
//   	},
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//   	NullOption: jsii.String("nullOption"),
//   	RelativeDateType: jsii.String("relativeDateType"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//
//   	// the properties below are optional
//   	ExcludePeriodConfiguration: &ExcludePeriodConfigurationProperty{
//   		Amount: jsii.Number(123),
//   		Granularity: jsii.String("granularity"),
//
//   		// the properties below are optional
//   		Status: jsii.String("status"),
//   	},
//   	MinimumGranularity: jsii.String("minimumGranularity"),
//   	ParameterName: jsii.String("parameterName"),
//   	RelativeDateValue: jsii.Number(123),
//   }
//
type CfnAnalysis_RelativeDatesFilterProperty struct {
	// `CfnAnalysis.RelativeDatesFilterProperty.AnchorDateConfiguration`.
	AnchorDateConfiguration interface{} `field:"required" json:"anchorDateConfiguration" yaml:"anchorDateConfiguration"`
	// `CfnAnalysis.RelativeDatesFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnAnalysis.RelativeDatesFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnAnalysis.RelativeDatesFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnAnalysis.RelativeDatesFilterProperty.RelativeDateType`.
	RelativeDateType *string `field:"required" json:"relativeDateType" yaml:"relativeDateType"`
	// `CfnAnalysis.RelativeDatesFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"required" json:"timeGranularity" yaml:"timeGranularity"`
	// `CfnAnalysis.RelativeDatesFilterProperty.ExcludePeriodConfiguration`.
	ExcludePeriodConfiguration interface{} `field:"optional" json:"excludePeriodConfiguration" yaml:"excludePeriodConfiguration"`
	// `CfnAnalysis.RelativeDatesFilterProperty.MinimumGranularity`.
	MinimumGranularity *string `field:"optional" json:"minimumGranularity" yaml:"minimumGranularity"`
	// `CfnAnalysis.RelativeDatesFilterProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// `CfnAnalysis.RelativeDatesFilterProperty.RelativeDateValue`.
	RelativeDateValue *float64 `field:"optional" json:"relativeDateValue" yaml:"relativeDateValue"`
}

