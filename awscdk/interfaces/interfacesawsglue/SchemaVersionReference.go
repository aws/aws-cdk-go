package interfacesawsglue


// A reference to a SchemaVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaVersionReference := &SchemaVersionReference{
//   	VersionId: jsii.String("versionId"),
//   }
//
type SchemaVersionReference struct {
	// The VersionId of the SchemaVersion resource.
	VersionId *string `field:"required" json:"versionId" yaml:"versionId"`
}

