package awslakeformation


// A structure for the database object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseResourceProperty := &databaseResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	name: jsii.String("name"),
//   }
//
type CfnPermissions_DatabaseResourceProperty struct {
	// `CfnPermissions.DatabaseResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database resource.
	//
	// Unique to the Data Catalog.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

