package awsentityresolution


// A reference to a SchemaMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaMappingReference := &SchemaMappingReference{
//   	SchemaName: jsii.String("schemaName"),
//   }
//
type SchemaMappingReference struct {
	// The SchemaName of the SchemaMapping resource.
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
}

