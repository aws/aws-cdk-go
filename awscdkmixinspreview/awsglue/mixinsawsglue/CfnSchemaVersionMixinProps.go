package mixinsawsglue


// Properties for CfnSchemaVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchemaVersionMixinProps := &CfnSchemaVersionMixinProps{
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
type CfnSchemaVersionMixinProps struct {
	// The schema that includes the schema version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversion.html#cfn-glue-schemaversion-schema
	//
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// The schema definition for the schema version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversion.html#cfn-glue-schemaversion-schemadefinition
	//
	SchemaDefinition *string `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

