package awsglue


// A structure that describes a target database for resource linking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseIdentifierProperty := &databaseIdentifierProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   }
//
type CfnDatabase_DatabaseIdentifierProperty struct {
	// The ID of the Data Catalog in which the database resides.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the catalog database.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

