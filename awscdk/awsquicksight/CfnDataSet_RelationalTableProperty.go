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
//   	InputColumns: []interface{}{
//   		&InputColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Catalog: jsii.String("catalog"),
//   	Schema: jsii.String("schema"),
//   }
//
type CfnDataSet_RelationalTableProperty struct {
	// The Amazon Resource Name (ARN) for the data source.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// The column schema of the table.
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// The name of the relational table.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDataSet.RelationalTableProperty.Catalog`.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// The schema name.
	//
	// This name applies to certain relational database engines.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

