package awsglue


// A reference to a SchemaVersionMetadata resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaVersionMetadataReference := &SchemaVersionMetadataReference{
//   	Key: jsii.String("key"),
//   	SchemaVersionId: jsii.String("schemaVersionId"),
//   	Value: jsii.String("value"),
//   }
//
type SchemaVersionMetadataReference struct {
	// The Key of the SchemaVersionMetadata resource.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The SchemaVersionId of the SchemaVersionMetadata resource.
	SchemaVersionId *string `field:"required" json:"schemaVersionId" yaml:"schemaVersionId"`
	// The Value of the SchemaVersionMetadata resource.
	Value *string `field:"required" json:"value" yaml:"value"`
}

