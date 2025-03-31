package awsquicksight


// The refresh properties of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetRefreshPropertiesProperty := &DataSetRefreshPropertiesProperty{
//   	FailureConfiguration: &RefreshFailureConfigurationProperty{
//   		EmailAlert: &RefreshFailureEmailAlertProperty{
//   			AlertStatus: jsii.String("alertStatus"),
//   		},
//   	},
//   	RefreshConfiguration: &RefreshConfigurationProperty{
//   		IncrementalRefresh: &IncrementalRefreshProperty{
//   			LookbackWindow: &LookbackWindowProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Size: jsii.Number(123),
//   				SizeUnit: jsii.String("sizeUnit"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetrefreshproperties.html
//
type CfnDataSet_DataSetRefreshPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetrefreshproperties.html#cfn-quicksight-dataset-datasetrefreshproperties-failureconfiguration
	//
	FailureConfiguration interface{} `field:"optional" json:"failureConfiguration" yaml:"failureConfiguration"`
	// The refresh configuration for a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetrefreshproperties.html#cfn-quicksight-dataset-datasetrefreshproperties-refreshconfiguration
	//
	RefreshConfiguration interface{} `field:"optional" json:"refreshConfiguration" yaml:"refreshConfiguration"`
}

