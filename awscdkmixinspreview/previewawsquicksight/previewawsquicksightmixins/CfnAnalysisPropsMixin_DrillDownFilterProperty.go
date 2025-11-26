package previewawsquicksightmixins


// The drill down filter for the column hierarchies.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-drilldownfilter.html
//
type CfnAnalysisPropsMixin_DrillDownFilterProperty struct {
	// The category type drill down filter.
	//
	// This filter is used for string type columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-drilldownfilter.html#cfn-quicksight-analysis-drilldownfilter-categoryfilter
	//
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// The numeric equality type drill down filter.
	//
	// This filter is used for number type columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-drilldownfilter.html#cfn-quicksight-analysis-drilldownfilter-numericequalityfilter
	//
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// The time range drill down filter.
	//
	// This filter is used for date time columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-drilldownfilter.html#cfn-quicksight-analysis-drilldownfilter-timerangefilter
	//
	TimeRangeFilter interface{} `field:"optional" json:"timeRangeFilter" yaml:"timeRangeFilter"`
}

