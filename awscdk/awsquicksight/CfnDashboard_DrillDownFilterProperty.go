package awsquicksight


// The drill down filter for the column hierarchies.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   drillDownFilterProperty := &DrillDownFilterProperty{
//   	CategoryFilter: &CategoryDrillDownFilterProperty{
//   		CategoryValues: []*string{
//   			jsii.String("categoryValues"),
//   		},
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   	},
//   	NumericEqualityFilter: &NumericEqualityDrillDownFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		Value: jsii.Number(123),
//   	},
//   	TimeRangeFilter: &TimeRangeDrillDownFilterProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		RangeMaximum: jsii.String("rangeMaximum"),
//   		RangeMinimum: jsii.String("rangeMinimum"),
//   		TimeGranularity: jsii.String("timeGranularity"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-drilldownfilter.html
//
type CfnDashboard_DrillDownFilterProperty struct {
	// The category type drill down filter.
	//
	// This filter is used for string type columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-drilldownfilter.html#cfn-quicksight-dashboard-drilldownfilter-categoryfilter
	//
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// The numeric equality type drill down filter.
	//
	// This filter is used for number type columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-drilldownfilter.html#cfn-quicksight-dashboard-drilldownfilter-numericequalityfilter
	//
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// The time range drill down filter.
	//
	// This filter is used for date time columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-drilldownfilter.html#cfn-quicksight-dashboard-drilldownfilter-timerangefilter
	//
	TimeRangeFilter interface{} `field:"optional" json:"timeRangeFilter" yaml:"timeRangeFilter"`
}

