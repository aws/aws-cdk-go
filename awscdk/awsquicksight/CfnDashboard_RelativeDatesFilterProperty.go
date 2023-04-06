package awsquicksight


// A `RelativeDatesFilter` filters relative dates values.
//
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
type CfnDashboard_RelativeDatesFilterProperty struct {
	// The date configuration of the filter.
	AnchorDateConfiguration interface{} `field:"required" json:"anchorDateConfiguration" yaml:"anchorDateConfiguration"`
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// The range date type of the filter. Choose one of the options below:.
	//
	// - `PREVIOUS`
	// - `THIS`
	// - `LAST`
	// - `NOW`
	// - `NEXT`.
	RelativeDateType *string `field:"required" json:"relativeDateType" yaml:"relativeDateType"`
	// The level of time precision that is used to aggregate `DateTime` values.
	TimeGranularity *string `field:"required" json:"timeGranularity" yaml:"timeGranularity"`
	// The configuration for the exclude period of the filter.
	ExcludePeriodConfiguration interface{} `field:"optional" json:"excludePeriodConfiguration" yaml:"excludePeriodConfiguration"`
	// The minimum granularity (period granularity) of the relative dates filter.
	MinimumGranularity *string `field:"optional" json:"minimumGranularity" yaml:"minimumGranularity"`
	// The parameter whose value should be used for the filter value.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The date value of the filter.
	RelativeDateValue *float64 `field:"optional" json:"relativeDateValue" yaml:"relativeDateValue"`
}

