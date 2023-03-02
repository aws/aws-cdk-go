package awsglue


// Properties for defining a `CfnSchemaVersionMetadata`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaVersionMetadataProps := &cfnSchemaVersionMetadataProps{
//   	key: jsii.String("key"),
//   	schemaVersionId: jsii.String("schemaVersionId"),
//   	value: jsii.String("value"),
//   }
//
type CfnSchemaVersionMetadataProps struct {
	// A metadata key in a key-value pair for metadata.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The version number of the schema.
	SchemaVersionId *string `field:"required" json:"schemaVersionId" yaml:"schemaVersionId"`
	// A metadata key's corresponding value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

