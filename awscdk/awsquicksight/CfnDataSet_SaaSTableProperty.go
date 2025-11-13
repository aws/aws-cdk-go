package awsquicksight


// A table from a Software-as-a-Service (SaaS) data source, including connection details and column definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   saaSTableProperty := &SaaSTableProperty{
//   	DataSourceArn: jsii.String("dataSourceArn"),
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
//   	TablePath: []interface{}{
//   		&TablePathElementProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html
//
type CfnDataSet_SaaSTableProperty struct {
	// The Amazon Resource Name (ARN) of the SaaS data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html#cfn-quicksight-dataset-saastable-datasourcearn
	//
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The list of input columns available from the SaaS table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html#cfn-quicksight-dataset-saastable-inputcolumns
	//
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// The hierarchical path to the table within the SaaS data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html#cfn-quicksight-dataset-saastable-tablepath
	//
	TablePath interface{} `field:"required" json:"tablePath" yaml:"tablePath"`
}

