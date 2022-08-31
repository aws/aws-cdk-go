package awsappsync


// The options for configuring a schema.
//
// If no options are specified, then the schema will
// be generated code-first.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaOptions := &schemaOptions{
//   	filePath: jsii.String("filePath"),
//   }
//
// Experimental.
type SchemaOptions struct {
	// The file path for the schema.
	//
	// When this option is
	// configured, then the schema will be generated from an
	// existing file from disk.
	// Experimental.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
}

