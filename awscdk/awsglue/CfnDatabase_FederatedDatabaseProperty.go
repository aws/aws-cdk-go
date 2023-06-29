package awsglue


// A `FederatedDatabase` structure that references an entity outside the AWS Glue Data Catalog .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   federatedDatabaseProperty := &FederatedDatabaseProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	Identifier: jsii.String("identifier"),
//   }
//
type CfnDatabase_FederatedDatabaseProperty struct {
	// The name of the connection to the external metastore.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A unique identifier for the federated database.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

