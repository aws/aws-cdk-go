package previewawsquicksightmixins


// The time range drill down filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeRangeDrillDownFilterProperty := &TimeRangeDrillDownFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	RangeMaximum: jsii.String("rangeMaximum"),
//   	RangeMinimum: jsii.String("rangeMinimum"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangedrilldownfilter.html
//
type CfnTemplatePropsMixin_TimeRangeDrillDownFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangedrilldownfilter.html#cfn-quicksight-template-timerangedrilldownfilter-column
	//
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// The maximum value for the filter value range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangedrilldownfilter.html#cfn-quicksight-template-timerangedrilldownfilter-rangemaximum
	//
	RangeMaximum *string `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// The minimum value for the filter value range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangedrilldownfilter.html#cfn-quicksight-template-timerangedrilldownfilter-rangeminimum
	//
	RangeMinimum *string `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timerangedrilldownfilter.html#cfn-quicksight-template-timerangedrilldownfilter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

