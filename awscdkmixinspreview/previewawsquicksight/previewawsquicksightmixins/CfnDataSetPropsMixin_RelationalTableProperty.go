package previewawsquicksightmixins


// A physical table type for relational data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relationalTableProperty := &RelationalTableProperty{
//   	Catalog: jsii.String("catalog"),
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			SubType: jsii.String("subType"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Schema: jsii.String("schema"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html
//
type CfnDataSetPropsMixin_RelationalTableProperty struct {
	// The catalog associated with a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-catalog
	//
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// The Amazon Resource Name (ARN) for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-datasourcearn
	//
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The column schema of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-inputcolumns
	//
	InputColumns interface{} `field:"optional" json:"inputColumns" yaml:"inputColumns"`
	// The name of the relational table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema name.
	//
	// This name applies to certain relational database engines.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-schema
	//
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

