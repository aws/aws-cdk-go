package awsquicksight


// The lookback window setup of an incremental refresh configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lookbackWindowProperty := &LookbackWindowProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Size: jsii.Number(123),
//   	SizeUnit: jsii.String("sizeUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html
//
type CfnDataSet_LookbackWindowProperty struct {
	// The name of the lookback window column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The lookback window column size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-size
	//
	// Default: - 0.
	//
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The size unit that is used for the lookback window column.
	//
	// Valid values for this structure are `HOUR` , `DAY` , and `WEEK` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-lookbackwindow.html#cfn-quicksight-dataset-lookbackwindow-sizeunit
	//
	SizeUnit *string `field:"required" json:"sizeUnit" yaml:"sizeUnit"`
}

