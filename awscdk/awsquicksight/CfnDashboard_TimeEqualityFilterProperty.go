package awsquicksight


// A `TimeEqualityFilter` filters values that are equal to a given value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeEqualityFilterProperty := &TimeEqualityFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//
//   	// the properties below are optional
//   	ParameterName: jsii.String("parameterName"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   	Value: jsii.String("value"),
//   }
//
type CfnDashboard_TimeEqualityFilterProperty struct {
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The parameter whose value should be used for the filter value.
	//
	// This field is mutually exclusive to `Value` .
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The level of time precision that is used to aggregate `DateTime` values.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
	// The value of a `TimeEquality` filter.
	//
	// This field is mutually exclusive to `ParameterName` .
	Value *string `field:"optional" json:"value" yaml:"value"`
}

