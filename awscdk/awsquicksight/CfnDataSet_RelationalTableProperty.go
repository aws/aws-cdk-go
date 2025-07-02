package awsquicksight


// A physical table type for relational data sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationalTableProperty := &RelationalTableProperty{
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Catalog: jsii.String("catalog"),
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			SubType: jsii.String("subType"),
//   		},
//   	},
//   	Schema: jsii.String("schema"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html
//
type CfnDataSet_RelationalTableProperty struct {
	// The Amazon Resource Name (ARN) for the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-datasourcearn
	//
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The name of the relational table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The catalog associated with a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-catalog
	//
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// The column schema of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-inputcolumns
	//
	InputColumns interface{} `field:"optional" json:"inputColumns" yaml:"inputColumns"`
	// The schema name.
	//
	// This name applies to certain relational database engines.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-relationaltable.html#cfn-quicksight-dataset-relationaltable-schema
	//
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

