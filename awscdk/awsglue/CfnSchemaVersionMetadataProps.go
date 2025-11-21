package awsglue


// Properties for defining a `CfnSchemaVersionMetadata`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaVersionMetadataProps := &CfnSchemaVersionMetadataProps{
//   	Key: jsii.String("key"),
//   	SchemaVersionId: jsii.String("schemaVersionId"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversionmetadata.html
//
type CfnSchemaVersionMetadataProps struct {
	// A metadata key in a key-value pair for metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversionmetadata.html#cfn-glue-schemaversionmetadata-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The version number of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversionmetadata.html#cfn-glue-schemaversionmetadata-schemaversionid
	//
	SchemaVersionId interface{} `field:"required" json:"schemaVersionId" yaml:"schemaVersionId"`
	// A metadata key's corresponding value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-schemaversionmetadata.html#cfn-glue-schemaversionmetadata-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

