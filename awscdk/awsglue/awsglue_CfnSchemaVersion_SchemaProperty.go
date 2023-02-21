package awsglue


// A wrapper structure to contain schema identity fields.
//
// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaProperty := &schemaProperty{
//   	registryName: jsii.String("registryName"),
//   	schemaArn: jsii.String("schemaArn"),
//   	schemaName: jsii.String("schemaName"),
//   }
//
type CfnSchemaVersion_SchemaProperty struct {
	// The name of the registry where the schema is stored.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

