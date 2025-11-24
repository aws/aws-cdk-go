package mixinsawsquicksight


// A source table that provides initial data from either a physical table or parent dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceTableProperty := &SourceTableProperty{
//   	DataSet: &ParentDataSetProperty{
//   		DataSetArn: jsii.String("dataSetArn"),
//   		InputColumns: []interface{}{
//   			&InputColumnProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				SubType: jsii.String("subType"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	PhysicalTableId: jsii.String("physicalTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-sourcetable.html
//
type CfnDataSetPropsMixin_SourceTableProperty struct {
	// A parent dataset that serves as the data source instead of a physical table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-sourcetable.html#cfn-quicksight-dataset-sourcetable-dataset
	//
	DataSet interface{} `field:"optional" json:"dataSet" yaml:"dataSet"`
	// The identifier of the physical table that serves as the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-sourcetable.html#cfn-quicksight-dataset-sourcetable-physicaltableid
	//
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

