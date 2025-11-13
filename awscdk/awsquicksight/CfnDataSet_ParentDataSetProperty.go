package awsquicksight


// References a parent dataset that serves as a data source, including its columns and metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parentDataSetProperty := &ParentDataSetProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Id: jsii.String("id"),
//   			SubType: jsii.String("subType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-parentdataset.html
//
type CfnDataSet_ParentDataSetProperty struct {
	// The Amazon Resource Name (ARN) of the parent dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-parentdataset.html#cfn-quicksight-dataset-parentdataset-datasetarn
	//
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// The list of input columns available from the parent dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-parentdataset.html#cfn-quicksight-dataset-parentdataset-inputcolumns
	//
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
}

