package awsglue


// Specifies the version of a schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaVersionProperty := &schemaVersionProperty{
//   	isLatest: jsii.Boolean(false),
//   	versionNumber: jsii.Number(123),
//   }
//
type CfnSchema_SchemaVersionProperty struct {
	// Indicates if this version is the latest version of the schema.
	IsLatest interface{} `field:"optional" json:"isLatest" yaml:"isLatest"`
	// The version number of the schema.
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

