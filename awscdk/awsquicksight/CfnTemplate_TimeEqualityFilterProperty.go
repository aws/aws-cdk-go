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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html
//
type CfnTemplate_TimeEqualityFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-filterid
	//
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// The parameter whose value should be used for the filter value.
	//
	// This field is mutually exclusive to `Value` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
	// The value of a `TimeEquality` filter.
	//
	// This field is mutually exclusive to `ParameterName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timeequalityfilter.html#cfn-quicksight-template-timeequalityfilter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

