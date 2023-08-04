package awsquicksight


// <p>The dataset refresh properties for the dataset.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetRefreshPropertiesProperty := &DataSetRefreshPropertiesProperty{
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
	// <p> Refresh Configuration.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetrefreshproperties.html#cfn-quicksight-dataset-datasetrefreshproperties-refreshconfiguration
	//
	RefreshConfiguration interface{} `field:"optional" json:"refreshConfiguration" yaml:"refreshConfiguration"`
}

