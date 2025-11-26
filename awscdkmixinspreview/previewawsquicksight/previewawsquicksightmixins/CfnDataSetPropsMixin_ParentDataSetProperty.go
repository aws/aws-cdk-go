package previewawsquicksightmixins


// References a parent dataset that serves as a data source, including its columns and metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parentDataSetProperty := &ParentDataSetProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			SubType: jsii.String("subType"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-parentdataset.html
//
type CfnDataSetPropsMixin_ParentDataSetProperty struct {
	// The Amazon Resource Name (ARN) of the parent dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-parentdataset.html#cfn-quicksight-dataset-parentdataset-datasetarn
	//
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// The list of input columns available from the parent dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-parentdataset.html#cfn-quicksight-dataset-parentdataset-inputcolumns
	//
	InputColumns interface{} `field:"optional" json:"inputColumns" yaml:"inputColumns"`
}

