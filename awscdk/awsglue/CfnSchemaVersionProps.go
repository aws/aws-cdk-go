package awsglue


// Properties for defining a `CfnSchemaVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaVersionProps := &CfnSchemaVersionProps{
//   	Schema: &SchemaProperty{
//   		RegistryName: jsii.String("registryName"),
//   		SchemaArn: jsii.String("schemaArn"),
//   		SchemaName: jsii.String("schemaName"),
//   	},
//   	SchemaDefinition: jsii.String("schemaDefinition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversion.html
//
type CfnSchemaVersionProps struct {
	// The schema that includes the schema version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversion.html#cfn-glue-schemaversion-schema
	//
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
	// The schema definition for the schema version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversion.html#cfn-glue-schemaversion-schemadefinition
	//
	SchemaDefinition *string `field:"required" json:"schemaDefinition" yaml:"schemaDefinition"`
}

