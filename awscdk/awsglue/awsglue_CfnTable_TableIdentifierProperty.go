package awsglue


// A structure that describes a target table for resource linking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableIdentifierProperty := &tableIdentifierProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   }
//
type CfnTable_TableIdentifierProperty struct {
	// The ID of the Data Catalog in which the table resides.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database that contains the target table.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the target table.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

