package awsglue


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
	// `CfnDatabase.FederatedDatabaseProperty.ConnectionName`.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// `CfnDatabase.FederatedDatabaseProperty.Identifier`.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

