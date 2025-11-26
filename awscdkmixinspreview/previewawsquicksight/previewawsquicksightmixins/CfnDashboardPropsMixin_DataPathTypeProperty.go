package previewawsquicksightmixins


// The type of the data path value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPathTypeProperty := &DataPathTypeProperty{
//   	PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapathtype.html
//
type CfnDashboardPropsMixin_DataPathTypeProperty struct {
	// The type of data path value utilized in a pivot table. Choose one of the following options:.
	//
	// - `HIERARCHY_ROWS_LAYOUT_COLUMN` - The type of data path for the rows layout column, when `RowsLayout` is set to `HIERARCHY` .
	// - `MULTIPLE_ROW_METRICS_COLUMN` - The type of data path for the metric column when the row is set to Metric Placement.
	// - `EMPTY_COLUMN_HEADER` - The type of data path for the column with empty column header, when there is no field in `ColumnsFieldWell` and the row is set to Metric Placement.
	// - `COUNT_METRIC_COLUMN` - The type of data path for the column with `COUNT` as the metric, when there is no field in the `ValuesFieldWell` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datapathtype.html#cfn-quicksight-dashboard-datapathtype-pivottabledatapathtype
	//
	PivotTableDataPathType *string `field:"optional" json:"pivotTableDataPathType" yaml:"pivotTableDataPathType"`
}

