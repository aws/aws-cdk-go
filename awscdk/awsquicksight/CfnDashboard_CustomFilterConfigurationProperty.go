package awsquicksight


// A custom filter that filters based on a single value.
//
// This filter can be partially matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customFilterConfigurationProperty := &CustomFilterConfigurationProperty{
//   	MatchOperator: jsii.String("matchOperator"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	CategoryValue: jsii.String("categoryValue"),
//   	ParameterName: jsii.String("parameterName"),
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html
//
type CfnDashboard_CustomFilterConfigurationProperty struct {
	// The match operator that is used to determine if a filter should be applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-matchoperator
	//
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-nulloption
	//
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// The category value for the filter.
	//
	// This field is mutually exclusive to `ParameterName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-categoryvalue
	//
	CategoryValue *string `field:"optional" json:"categoryValue" yaml:"categoryValue"`
	// The parameter whose value should be used for the filter value.
	//
	// This field is mutually exclusive to `CategoryValue` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Select all of the values. Null is not the assigned value of select all.
	//
	// - `FILTER_ALL_VALUES`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-selectalloptions
	//
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}

