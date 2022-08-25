package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   resourceProperty := &resourceProperty{
//   	catalog: catalog,
//   	database: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	table: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   		tableWildcard: tableWildcard,
//   	},
//   	tableWithColumns: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnTagAssociation_ResourceProperty struct {
	// `CfnTagAssociation.ResourceProperty.Catalog`.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// `CfnTagAssociation.ResourceProperty.Database`.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// `CfnTagAssociation.ResourceProperty.Table`.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// `CfnTagAssociation.ResourceProperty.TableWithColumns`.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}

