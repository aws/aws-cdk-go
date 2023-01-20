package awslakeformation


// A structure for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceProperty := &resourceProperty{
//   	databaseResource: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	dataLocationResource: &dataLocationResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		s3Resource: jsii.String("s3Resource"),
//   	},
//   	tableResource: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   		tableWildcard: &tableWildcardProperty{
//   		},
//   	},
//   	tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		columnWildcard: &columnWildcardProperty{
//   			excludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnPermissions_ResourceProperty struct {
	// A structure for the database object.
	DatabaseResource interface{} `field:"optional" json:"databaseResource" yaml:"databaseResource"`
	// A structure for a data location object where permissions are granted or revoked.
	DataLocationResource interface{} `field:"optional" json:"dataLocationResource" yaml:"dataLocationResource"`
	// A structure for the table object.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	TableResource interface{} `field:"optional" json:"tableResource" yaml:"tableResource"`
	// Currently not supported by AWS CloudFormation .
	TableWithColumnsResource interface{} `field:"optional" json:"tableWithColumnsResource" yaml:"tableWithColumnsResource"`
}

