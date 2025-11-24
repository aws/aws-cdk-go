package mixinsawsquicksight


// A table from a Software-as-a-Service (SaaS) data source, including connection details and column definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   saaSTableProperty := &SaaSTableProperty{
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			SubType: jsii.String("subType"),
//   			Type: jsii.String("type"),
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
type CfnDataSetPropsMixin_SaaSTableProperty struct {
	// The Amazon Resource Name (ARN) of the SaaS data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html#cfn-quicksight-dataset-saastable-datasourcearn
	//
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The list of input columns available from the SaaS table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html#cfn-quicksight-dataset-saastable-inputcolumns
	//
	InputColumns interface{} `field:"optional" json:"inputColumns" yaml:"inputColumns"`
	// The hierarchical path to the table within the SaaS data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-saastable.html#cfn-quicksight-dataset-saastable-tablepath
	//
	TablePath interface{} `field:"optional" json:"tablePath" yaml:"tablePath"`
}

