package mixinsawsglue


// A wrapper structure to contain schema identity fields.
//
// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaProperty := &SchemaProperty{
//   	RegistryName: jsii.String("registryName"),
//   	SchemaArn: jsii.String("schemaArn"),
//   	SchemaName: jsii.String("schemaName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schemaversion-schema.html
//
type CfnSchemaVersionPropsMixin_SchemaProperty struct {
	// The name of the registry where the schema is stored.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schemaversion-schema.html#cfn-glue-schemaversion-schema-registryname
	//
	RegistryName *string `field:"optional" json:"registryName" yaml:"registryName"`
	// The Amazon Resource Name (ARN) of the schema.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schemaversion-schema.html#cfn-glue-schemaversion-schema-schemaarn
	//
	SchemaArn *string `field:"optional" json:"schemaArn" yaml:"schemaArn"`
	// The name of the schema.
	//
	// Either `SchemaArn` , or `SchemaName` and `RegistryName` has to be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schemaversion-schema.html#cfn-glue-schemaversion-schema-schemaname
	//
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

