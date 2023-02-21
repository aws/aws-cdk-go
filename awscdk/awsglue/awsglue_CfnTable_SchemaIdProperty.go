package awsglue


// A structure that contains schema identity fields.
//
// Either this or the `SchemaVersionId` has to be
// provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaIdProperty := &schemaIdProperty{
//   	registryName: jsii.String("registryName"),
//   	schemaArn: jsii.String("schemaArn"),
//   	schemaName: jsii.String("schemaName"),
//   }
//
type CfnTable_SchemaIdProperty struct {
	// The name of the schema registry that contains the schema.
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be
	// provided.
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// One of `SchemaArn` or `SchemaName` has to be provided.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

