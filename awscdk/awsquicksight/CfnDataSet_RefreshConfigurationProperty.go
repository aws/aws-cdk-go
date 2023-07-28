package awsquicksight


// The refresh configuration of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshConfigurationProperty := &RefreshConfigurationProperty{
//   	IncrementalRefresh: &IncrementalRefreshProperty{
//   		LookbackWindow: &LookbackWindowProperty{
//   			ColumnName: jsii.String("columnName"),
//   			Size: jsii.Number(123),
//   			SizeUnit: jsii.String("sizeUnit"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshconfiguration.html
//
type CfnDataSet_RefreshConfigurationProperty struct {
	// The incremental refresh for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshconfiguration.html#cfn-quicksight-dataset-refreshconfiguration-incrementalrefresh
	//
	IncrementalRefresh interface{} `field:"optional" json:"incrementalRefresh" yaml:"incrementalRefresh"`
}

