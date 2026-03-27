package awsglue


// A FederatedCatalog structure that references an entity outside the Glue Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   federatedCatalogProperty := &FederatedCatalogProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html
//
type CfnCatalogPropsMixin_FederatedCatalogProperty struct {
	// The name of the connection to an external data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html#cfn-glue-catalog-federatedcatalog-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A unique identifier for the federated catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-catalog-federatedcatalog.html#cfn-glue-catalog-federatedcatalog-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

