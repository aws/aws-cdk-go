package awsquicksight


// The incremental refresh configuration for a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incrementalRefreshProperty := &IncrementalRefreshProperty{
//   	LookbackWindow: &LookbackWindowProperty{
//   		ColumnName: jsii.String("columnName"),
//   		Size: jsii.Number(123),
//   		SizeUnit: jsii.String("sizeUnit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-incrementalrefresh.html
//
type CfnDataSet_IncrementalRefreshProperty struct {
	// The lookback window setup for an incremental refresh configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-incrementalrefresh.html#cfn-quicksight-dataset-incrementalrefresh-lookbackwindow
	//
	LookbackWindow interface{} `field:"optional" json:"lookbackWindow" yaml:"lookbackWindow"`
}

