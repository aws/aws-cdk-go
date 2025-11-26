package previewawsgluemixins


// A `FederatedDatabase` structure that references an entity outside the AWS Glue Data Catalog .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   federatedDatabaseProperty := &FederatedDatabaseProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-federateddatabase.html
//
type CfnDatabasePropsMixin_FederatedDatabaseProperty struct {
	// The name of the connection to the external metastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-federateddatabase.html#cfn-glue-database-federateddatabase-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A unique identifier for the federated database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-federateddatabase.html#cfn-glue-database-federateddatabase-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

